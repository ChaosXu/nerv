package manager

import (
	"fmt"
	"github.com/ChaosXu/nerv/lib/log"
	"github.com/ChaosXu/nerv/lib/lock"
	"github.com/ChaosXu/nerv/lib/automation/model/topology"
	"time"
	"github.com/ChaosXu/nerv/lib/db"
	templaterep "github.com/ChaosXu/nerv/lib/automation/repository"
	classrep "github.com/ChaosXu/nerv/lib/resource/repository"
	"github.com/ChaosXu/nerv/lib/operation"
	"encoding/json"
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
	Executor    operation.Executor `inject:""`
}

// Create a topology in db
func (p *Deployer) Create(topoName string, templatePath string, inputs map[string]interface{}) (uint, error) {
	return p.create(topoName, templatePath, inputs, 1)
}

func (p *Deployer) Install(topoId uint) error {
	topo := &topology.Topology{}
	if err := db.DB.First(topo, topoId).Error; err != nil {
		return err
	}

	lock := lock.GetLock("Topology", topo.ID)
	if !lock.TryLock() {
		return fmt.Errorf("topology is doing. ID=%d", topo.ID)
	}
	defer lock.Unlock()

	return p.postTraverse(topo, "contained", "Create")
}

// Uninstall the topology
func (p *Deployer) Uninstall(topoId uint) error {
	topo := &topology.Topology{}
	if err := db.DB.First(topo, topoId).Error; err != nil {
		return err
	}

	lock := lock.GetLock("Topology", topo.ID)
	if !lock.TryLock() {
		return fmt.Errorf("topology is doing. ID=%d", topo.ID)
	}
	defer lock.Unlock()

	return p.preTraverse(topo, "contained", "Delete")
}

// Start the topology
func (p *Deployer) Start(topoId uint) error {
	topo := &topology.Topology{}
	if err := db.DB.First(topo, topoId).Error; err != nil {
		return err
	}

	lock := lock.GetLock("Topology", topo.ID)
	if !lock.TryLock() {
		return fmt.Errorf("topology is doing. ID=%d", topo.ID)
	}
	defer lock.Unlock()

	return p.postTraverse(topo, "contained", "Start")
}

// Stop the topology
func (p *Deployer) Stop(topoId uint) error {
	topo := &topology.Topology{}
	if err := db.DB.First(topo, topoId).Error; err != nil {
		return err
	}

	lock := lock.GetLock("Topology", topo.ID)
	if !lock.TryLock() {
		return fmt.Errorf("topology is doing. ID=%d", topo.ID)
	}
	defer lock.Unlock()

	return p.preTraverse(topo, "contained", "Stop")
}

// Restart the topology
func (p *Deployer) Restart(topoId uint) error {
	err := p.Stop(topoId)
	if err != nil {
		return err
	}

	lock := lock.GetLock("Topology", topoId)
	if !lock.TryLock() {
		return fmt.Errorf("topology is doing. ID=%d", topoId)
	}
	defer lock.Unlock()

	return p.Start(topoId)
}

// Setup the topology
func (p *Deployer) Setup(topoId uint) error {
	topo := &topology.Topology{}
	if err := db.DB.First(topo, topoId).Error; err != nil {
		return err
	}

	lock := lock.GetLock("Topology", topo.ID)
	if !lock.TryLock() {
		return fmt.Errorf("topology is doing. ID=%d", topo.ID)
	}
	defer lock.Unlock()

	return p.postTraverse(topo, "contained", "Setup")
}

// Reload configuration after setup when the topology has been started.
func (p *Deployer) Reload(topoId uint) error {
	topo := &topology.Topology{}
	if err := db.DB.First(topo, topoId).Error; err != nil {
		return err
	}

	lock := lock.GetLock("Topology", topo.ID)
	if !lock.TryLock() {
		return fmt.Errorf("topology is doing. ID=%d", topo.ID)
	}
	defer lock.Unlock()

	return p.postTraverse(topo, "contained", "Reload")
}

// Migrate a topology with topoId to a new topology.
// Update all inputs but continue to have the template of source topology.
func (p *Deployer) Migrate(topoId uint, inputs map[string]interface{}) (uint, error) {
	topo := &topology.Topology{}
	if err := db.DB.First(topo, topoId).Error; err != nil {
		return 0, err
	}

	lock := lock.GetLock("Topology", topo.ID)
	if !lock.TryLock() {
		return 0, fmt.Errorf("topology is doing. ID=%d", topo.ID)
	}
	defer lock.Unlock()

	return p.create(topo.Name, topo.Template, inputs, topo.Version + 1)
}

func (p *Deployer) create(topoName string, templatePath string, inputs map[string]interface{}, version int) (uint, error) {
	template, err := p.TemplateRep.GetTemplate(templatePath)
	if err != nil {
		return 0, fmt.Errorf("could not found template %s. %s", templatePath, err)
	}
	ctx := topology.NewContext(template.Inputs, inputs)
	topo := template.NewTopology(topoName, version, ctx)

	p.DBService.GetDB().Save(topo)
	return topo.ID, nil
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
	tnodes := []*topology.Node{}
	p.DBService.GetDB().Where("topology_id =?", topo.ID).Preload("Links").Find(&tnodes)
	topo.Nodes = tnodes

	template, err := p.TemplateRep.GetTemplate(topo.Template)
	if err != nil {
		return err
	}

	status := []PerformStatus{}
	for _, node := range topo.Nodes {
		done, timeout := p.preTraverseNode(topo, depType, node, template, operation)
		status = append(status, PerformStatus{Node:node, Done:done, Timeout:timeout})
	}

	for _, item := range status {
		select {
		case <-item.Done:
			if item.Node.Error != "" {
				//fmt.Println("postTraNode " + item.Node.Error)
				err = fmt.Errorf(item.Node.Error)
			}
		case <-item.Timeout:
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
	tnodes := []*topology.Node{}
	p.DBService.GetDB().Where("topology_id =?", topo.ID).Preload("Links").Find(&tnodes)
	topo.Nodes = tnodes

	template, err := p.TemplateRep.GetTemplate(topo.Template)
	if err != nil {
		return err
	}

	status := []PerformStatus{}
	for _, node := range topo.Nodes {
		done, timeout := p.postTraverseNode(topo, depType, node, template, operation)
		status = append(status, PerformStatus{Node:node, Done:done, Timeout:timeout})
	}

	for _, item := range status {
		select {
		case <-item.Done:
			if item.Node.Error != "" {
				//fmt.Println("postTraNode " + item.Node.Error)
				err = fmt.Errorf(item.Node.Error)
			}
		case <-item.Timeout:
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
		status := []PerformStatus{}
		for _, link := range links {
			node := topo.GetNode(link.Target)
			done, timeout := p.preTraverseNode(topo, depType, node, template, operation)
			status = append(status, PerformStatus{Node:node, Done:done, Timeout:timeout})
		}

		for _, item := range status {
			select {
			case <-item.Done:
				if item.Node.Error != "" {
					//fmt.Println("postTraNode " + item.Node.Error)
					childErr = fmt.Errorf(item.Node.Error)
				}
			case t := <-item.Timeout:
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
		status := []PerformStatus{}
		for _, link := range links {
			node := topo.GetNode(link.Target)
			done, timeout := p.postTraverseNode(topo, depType, node, template, operation)
			status = append(status, PerformStatus{Node:node, Done:done, Timeout:timeout})
		}

		var err error = nil
		for _, item := range status {
			select {
			case <-item.Done:
				if item.Node.Error != "" {
					//fmt.Println("postTraNode " + item.Node.Error)
					err = fmt.Errorf(item.Node.Error)
				}
			case <-item.Timeout:
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
	//TBD:don't hard code
	args["address"] = node.Address
	args["credential"] = node.Credential

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

