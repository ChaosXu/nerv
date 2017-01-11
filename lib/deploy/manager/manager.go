package manager

import (
	"github.com/ChaosXu/nerv/lib/deploy/repository"
	"fmt"
	"github.com/ChaosXu/nerv/lib/log"
	"github.com/ChaosXu/nerv/lib/lock"
	"github.com/ChaosXu/nerv/lib/deploy/model/topology"
	"time"
	"github.com/ChaosXu/nerv/lib/db"
)

// Manager execute the deployment task.
type Manager struct {
	TemplateRep repository.TemplateRepository `inject:""`
	DBService   *db.DBService `inject:""`
}

//Install the topology and start to serve
func (p *Manager) Install(topoName string, templatePath string) error {
	log.LogCodeLine()
	template,err := p.TemplateRep.GetTemplate(templatePath)
	if err!=nil{
		return err
	}
	topo := template.NewTopology(topoName)
	return p.postTraverse(topo, "contained", "Create")
}

//Uninstall the topology
func (p *Manager) Uninstall(topology *topology.Topology) error {
	log.LogCodeLine()
	return p.preTraverse(topology, "contained", "Delete")
}

//Configure the topology for start
func (p *Manager) Configure(topology *topology.Topology) error {
	return fmt.Errorf("TBD")
}

//Start the Topology
func (p *Manager) Start(topology *topology.Topology) error {
	log.LogCodeLine()
	return p.postTraverse(topology, "contained", "Start")
}

//Stop the Topology
func (p *Manager) Stop(topology *topology.Topology) error {
	log.LogCodeLine()
	return p.preTraverse(topology, "contained", "Stop")
}

func (p *Manager) preTraverse(topo *topology.Topology, depType string, operation string) error {
	lock := lock.GetLock("Topology", topo.ID)
	if !lock.TryLock() {
		return fmt.Errorf("topology is doing. ID=%d", topo.ID)
	}
	defer lock.Unlock()

	tnodes := []*topology.Node{}
	p.DBService.GetDB().Where("topology_id =?", topo.ID).Preload("Links").Find(&tnodes)
	topo.Nodes = tnodes

	template := topology.ServiceTemplate{}
	if err := p.DBService.GetDB().Where("name=? and version=?", topo.Template, topo.Version).Preload("Nodes").Preload("Nodes.Parameters").First(&template).Error; err != nil {
		return err
	}

	dones := []<-chan error{}
	timeouts := []<-chan bool{}

	for _, node := range topo.Nodes {
		done, timeout := p.preTraverseNode(topo, depType, node, &template, operation)
		dones = append(dones, done)
		timeouts = append(timeouts, timeout)
	}

	var err error = nil
	for i, done := range dones {
		select {
		case e := <-done:
			if e != nil {
				err = e
			}
		case <-timeouts[i]:
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

func (p *Manager) postTraverse(topo *topology.Topology, depType string, operation string) error {
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
				err = e
			}
		case <-timeouts[i]:
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

func (p *Manager) executeNode(operation string, node *topology.Node, template *topology.ServiceTemplate) (<-chan error, <-chan bool) {
	done := make(chan error, 1)
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(30 * time.Second)
		timeout <- true
	}()

	go func() {
		if err := node.Execute(operation, template); err != nil {
			done <- err
		} else {
			done <- nil
		}
	}()

	return done, timeout
}

func (p *Manager) preTraverseNode(topo *topology.Topology, depType string, parent *topology.Node, template *topology.ServiceTemplate, operation string) (<-chan error, <-chan bool) {
	err, timeout := p.executeNode(operation, parent, template)
	select {
	case e := <-err:
		ec := make(chan error, 1)
		ec <- e
		return ec, timeout
	case <-timeout:
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

func (p *Manager) postTraverseNode(topo *topology.Topology, depType string, parent *topology.Node, template *topology.ServiceTemplate, operation string) (<-chan error, <-chan bool) {
	links := parent.FindLinksByType(depType)
	if links != nil && len(links) > 0 {
		dones := []<-chan error{}
		timeouts := []<-chan bool{}
		for _, link := range links {
			node := topo.GetNode(link.Target)
			done, timeout := p.postTraverseNode(topo, depType, node, template, operation)
			dones = append(dones, done)
			timeouts = append(timeouts, timeout)
		}

		var err error = nil
		for i, done := range dones {
			select {
			case e := <-done:
				if e != nil {
					err = e
				}
			case <-timeouts[i]:
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

