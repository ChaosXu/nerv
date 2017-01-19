package manager

import (
	"fmt"
	"github.com/ChaosXu/nerv/lib/log"
	"github.com/ChaosXu/nerv/lib/lock"
	"github.com/ChaosXu/nerv/lib/deploy/model/topology"
	"time"
	"github.com/ChaosXu/nerv/lib/db"
	templaterep "github.com/ChaosXu/nerv/lib/deploy/repository"
	classrep "github.com/ChaosXu/nerv/lib/resource/repository"
	"k8s.io/kubernetes/pkg/util/json"
	"github.com/ChaosXu/nerv/lib/resource/environment"
)

// PerformStatus trace the status of node executing
type PerformStatus struct {
	Node    *topology.Node
	Done    <-chan error
	Timeout <-chan bool
}

// Deployer execute the deployment task.
type Deployer struct {
	DBService   *db.DBService `inject:""`
	TemplateRep templaterep.TemplateRepository `inject:""`
	ClassRep    classrep.ClassRepository `inject:""`
	Executor    environment.Executor `inject:""`
}

// Create a topology in db
func (p *Deployer) Create(topoName string, templatePath string) (uint, error) {
	log.LogCodeLine()
	template, err := p.TemplateRep.GetTemplate(templatePath)
	if err != nil {
		return 0, err
	}
	topo := template.NewTopology(topoName)

	p.DBService.GetDB().Save(topo)
	return topo.ID, nil
}

func (p *Deployer) Install(topoId uint) error {
	topo := &topology.Topology{}
	if err := db.DB.First(topo, topoId).Error; err != nil {
		return err
	}

	return p.postTraverse(topo, "contained", "Create")
}

// Uninstall the topology
func (p *Deployer) Uninstall(topoId uint) error {
	topo := &topology.Topology{}
	if err := db.DB.First(topo, topoId).Error; err != nil {
		return err
	}

	return p.preTraverse(topo, "contained", "Delete")
}

// Configure the topology for start
func (p *Deployer) Configure(topology *topology.Topology) error {
	return fmt.Errorf("TBD")
}

// Start the topology
func (p *Deployer) Start(topoId uint) error {
	topo := &topology.Topology{}
	if err := db.DB.First(topo, topoId).Error; err != nil {
		return err
	}
	return p.postTraverse(topo, "contained", "Start")
}

// Stop the topology
func (p *Deployer) Stop(topoId uint) error {
	topo := &topology.Topology{}
	if err := db.DB.First(topo, topoId).Error; err != nil {
		return err
	}
	return p.preTraverse(topo, "contained", "Stop")
}

// Restart the topology
func (p *Deployer) Restart(topoId uint) error {
	err := p.Stop(topoId)
	if err != nil {
		return err
	}

	return p.Start(topoId)
}

// Setup the topology
func (p *Deployer) Setup(topoId uint) error {
	topo := &topology.Topology{}
	if err := db.DB.First(topo, topoId).Error; err != nil {
		return err
	}
	return p.postTraverse(topo, "contained", "Setup")
}

func (p *Deployer) dump(topo *topology.Topology) error {
	data, err := json.Marshal(topo)
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}

func (p *Deployer) preTraverse(topo *topology.Topology, depType string, operation string) error {
	lock := lock.GetLock("Topology", topo.ID)
	if !lock.TryLock() {
		return fmt.Errorf("topology is doing. ID=%d", topo.ID)
	}
	defer lock.Unlock()

	tnodes := []*topology.Node{}
	p.DBService.GetDB().Where("topology_id =?", topo.ID).Preload("Links").Find(&tnodes)
	topo.Nodes = tnodes

	template, err := p.TemplateRep.GetTemplate(topo.Template)
	if err != nil {
		return err
	}

	dones := []<-chan error{}
	timeouts := []<-chan bool{}

	for _, node := range topo.Nodes {
		done, timeout := p.preTraverseNode(topo, depType, node, template, operation)
		dones = append(dones, done)
		timeouts = append(timeouts, timeout)
	}

	for i, done := range dones {
		select {
		case e := <-done:
			if e != nil {
				err = e
			}
		case <-timeouts[i]:
			fmt.Println("timeout")
		}
	}

	if err == nil {
		topo.RunStatus = topology.RunStatusGreen
	} else {
		topo.Error = err.Error()
		topo.RunStatus = topology.RunStatusRed
	}
	p.DBService.GetDB().Save(topo)
	return err
}

func (p *Deployer) postTraverse(topo *topology.Topology, depType string, operation string) error {
	lock := lock.GetLock("Topology", topo.ID)
	if !lock.TryLock() {
		return fmt.Errorf("topology is doing. ID=%d", topo.ID)
	}
	defer lock.Unlock()

	tnodes := []*topology.Node{}
	p.DBService.GetDB().Where("topology_id =?", topo.ID).Preload("Links").Find(&tnodes)
	topo.Nodes = tnodes

	template, err := p.TemplateRep.GetTemplate(topo.Template)
	if err != nil {
		return err
	}

	dones := []<-chan error{}
	timeouts := []<-chan bool{}

	for _, node := range topo.Nodes {
		done, timeout := p.postTraverseNode(topo, depType, node, template, operation)
		dones = append(dones, done)
		timeouts = append(timeouts, timeout)
	}

	for i, done := range dones {
		select {
		case e := <-done:
			if e != nil {
				fmt.Println("postTra " + e.Error())
				err = e
			}
		case <-timeouts[i]:
			fmt.Println("timeout")
		}
	}

	if err == nil {
		topo.RunStatus = topology.RunStatusGreen
	} else {
		topo.Error = err.Error()
		topo.RunStatus = topology.RunStatusRed
	}
	p.DBService.GetDB().Save(topo)
	return err
}

func (p *Deployer) preTraverseNode(topo *topology.Topology, depType string, parent *topology.Node, template *topology.ServiceTemplate, operation string) (<-chan error, <-chan bool) {
	err, timeout := p.executeNode(operation, parent, template)
	select {
	case e := <-err:
		ec := make(chan error, 1)
		ec <- e
		return ec, timeout
	case <-timeout:
		fmt.Println("timeout")
	}

	var childErr error = nil
	var childTimeout bool

	links := parent.FindLinksByType(depType)
	if links != nil && len(links) > 0 {
		dones := []<-chan error{}
		timeouts := []<-chan bool{}
		for _, link := range links {
			node := topo.GetNode(link.Target)
			done, timeout := p.preTraverseNode(topo, depType, node, template, operation)
			dones = append(dones, done)
			timeouts = append(timeouts, timeout)
		}

		for i, done := range dones {
			select {
			case e := <-done:
				if e != nil {
					childErr = e
				}
			case t := <-timeouts[i]:
				if t {
					childTimeout = t
					fmt.Println("timeout")
				}
			}
		}
	}

	dc := make(chan error, 1)
	tc := make(chan bool, 1)

	if err == nil {
		dc <- childErr
	}
	if timeout == nil {
		tc <- childTimeout
	}
	return dc, tc
}

func (p *Deployer) postTraverseNode(topo *topology.Topology, depType string, parent *topology.Node, template *topology.ServiceTemplate, operation string) (<-chan error, <-chan bool) {
	links := parent.FindLinksByType(depType)
	if links != nil && len(links) > 0 {
		dones := []PerformStatus{}
		for _, link := range links {
			node := topo.GetNode(link.Target)
			done, timeout := p.postTraverseNode(topo, depType, node, template, operation)
			dones = append(dones, PerformStatus{Node:node, Done:done, Timeout:timeout})
		}

		var err error = nil
		for _, status := range dones {
			select {
			case <-status.Done:
				if status.Node.Error != "" {
					fmt.Println("postTraNode " + status.Node.Error)
					err = fmt.Errorf(status.Node.Error)
				}
			case <-status.Timeout:
				fmt.Println("timeout")
			}
		}

		if err == nil {
			return p.executeNode(operation, parent, template)
		} else {
			done := make(chan error, 1)
			timeout := make(chan bool, 1)
			done <- err
			return done, timeout
		}
	} else {
		return p.executeNode(operation, parent, template)
	}
}

func (p *Deployer) executeNode(operation string, node *topology.Node, template *topology.ServiceTemplate) (<-chan error, <-chan bool) {

	if node.Done == nil {
		node.Done = make(chan error, 1)
		node.Timeout = make(chan bool, 1)

		go func() {
			time.Sleep(30 * time.Second)
			node.Timeout <- true
			close(node.Timeout)
		}()

		go func() {
			if err := p.invoke(node, operation, template); err != nil {
				node.Done <- err

			} else {
				node.Done <- nil
			}
			close(node.Done)
		}()
	} else {
		fmt.Println("doing " + node.Name)
	}

	return node.Done, node.Timeout
}

// invoke the operation
func (p *Deployer) invoke(node *topology.Node, operation string, template *topology.ServiceTemplate) error {
	log.LogCodeLine()

	nodeTemplate := template.FindTemplate(node.Template)

	if nodeTemplate == nil {
		node.RunStatus = topology.RunStatusRed
		err := fmt.Errorf("template %s of node %s isn't exist", node.Template, node.Name)
		node.Error = err.Error()
		return err
	}
	node.RunStatus = topology.RunStatusGreen

	args := map[string]string{}
	for _, param := range nodeTemplate.Parameters {
		args[param.Name] = param.Value
	}

	class, err := p.ClassRep.Get(nodeTemplate.Type)
	if err != nil {
		node.RunStatus = topology.RunStatusRed
		err = fmt.Errorf("load class %s by template %s failed. error: %s", node.Name, node.Template, err.Error())
		node.Error = err.Error()
		return err
	}

	err = p.Executor.Perform(template.Environment, class, operation, args)
	if err != nil {
		node.RunStatus = topology.RunStatusRed
		node.Error = fmt.Errorf("%s execute %s error:%s", node.Name, operation, err.Error()).Error()

	} else {
		node.RunStatus = topology.RunStatusGreen
		node.Error = ""
	}

	err = p.DBService.GetDB().Save(node).Error
	if err != nil {
		node.RunStatus = topology.RunStatusRed
		err = fmt.Errorf("save node status %s failed. error: %s", node.Name, err.Error())
		node.Error = err.Error()
	}
	return err
}

