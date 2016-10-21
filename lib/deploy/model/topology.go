package model

import (
	"github.com/jinzhu/gorm"
	"github.com/ChaosXu/nerv/lib/db"
	"fmt"
	"time"
	"github.com/ChaosXu/nerv/lib/log"
)

func init() {
	db.Models["Topology"] = topologyDesc()
}

func topologyDesc() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &Topology{},
		New: func() interface{} {
			return &Topology{}
		},
		NewSlice:func() interface{} {
			return &[]Topology{}
		},
	}
}

type traverseCallback func(node *Node, template *ServiceTemplate) (<-chan error, <-chan bool)

//Topology which has been deployed by the service template
type Topology struct {
	gorm.Model
	Status
	Name     string //topology name
	Template string //service template name
	Version  int32  //service template version
	Nodes    []*Node
}

//Install the topology and start to serve
func (p *Topology) Install() error {
	log.LogCodeLine()
	return p.postTraverse("contained", "Create")
}

//Uninstall the topology
func (p *Topology) Uninstall() error {
	log.LogCodeLine()
	return p.preTraverse("contained", "Delete")
}

//Configure the topology for start
func (p *Topology) Configure() error {
	return fmt.Errorf("TBD")
}

//Start the Topology
func (p *Topology) Start() error {
	log.LogCodeLine()
	return p.postTraverse("contained", "Start")
}

//Stop the Topology
func (p *Topology) Stop() error {
	log.LogCodeLine()
	return p.preTraverse("contained", "Stop")
}

func (p *Topology) preTraverse(depType string, operation string) error {
	lock := GetLock("Topology", p.ID)
	if !lock.TryLock() {
		return fmt.Errorf("topology is doing. ID=%d", p.ID)
	}
	defer lock.Unlock()

	tnodes := []*Node{}
	db.DB.Where("topology_id =?", p.ID).Preload("Links").Find(&tnodes)
	p.Nodes = tnodes

	template := ServiceTemplate{}
	if err := db.DB.Where("name=? and version=?", p.Template, p.Version).Preload("Nodes").Preload("Nodes.Parameters").First(&template).Error; err != nil {
		return err
	}

	dones := []<-chan error{}
	timeouts := []<-chan bool{}

	for _, node := range p.Nodes {
		done, timeout := p.preTraverseNode(depType, node, &template, operation)
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
		p.RunStatus = RunStatusGreen
	} else {
		p.Error = err.Error()
		p.RunStatus = RunStatusRed
	}
	db.DB.Save(p)
	return err
}

func (p *Topology) postTraverse(depType string, operation string) error {
	lock := GetLock("Topology", p.ID)
	if !lock.TryLock() {
		return fmt.Errorf("topology is doing. ID=%d", p.ID)
	}
	defer lock.Unlock()

	tnodes := []*Node{}
	db.DB.Where("topology_id =?", p.ID).Preload("Links").Find(&tnodes)
	p.Nodes = tnodes

	template := ServiceTemplate{}
	if err := db.DB.Where("name=? and version=?", p.Template, p.Version).Preload("Nodes").Preload("Nodes.Parameters").First(&template).Error; err != nil {
		return err
	}

	dones := []<-chan error{}
	timeouts := []<-chan bool{}

	for _, node := range p.Nodes {
		done, timeout := p.postTraverseNode(depType, node, &template, operation)
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
		p.RunStatus = RunStatusGreen
	} else {
		p.Error = err.Error()
		p.RunStatus = RunStatusRed
	}
	db.DB.Save(p)
	return err
}

func (p *Topology) executeNode(operation string, node *Node, template *ServiceTemplate) (<-chan error, <-chan bool) {
	done := make(chan error, 1)
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(30 * time.Second)
		timeout <- true
	}()

	go func() {
		nodeTemplate := template.findTemplate(node.Template)
		if nodeTemplate != nil {
			if err := node.Execute(operation, nodeTemplate); err != nil {
				done <- err
			} else {
				done <- nil
			}
		} else {
			node.RunStatus = RunStatusRed
			err := fmt.Errorf("template %s of node %s isn't exist", node.Template, node.Name)
			node.Error = err.Error()
			db.DB.Save(node)
			done <- err
		}
	}()

	return done, timeout
}

func (p *Topology) preTraverseNode(depType string, parent *Node, template *ServiceTemplate, operation string) (<-chan error, <-chan bool) {
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
			node := p.getNode(link.Target)
			done, timeout := p.preTraverseNode(depType, node, template, operation)
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

func (p *Topology) postTraverseNode(depType string, parent *Node, template *ServiceTemplate, operation string) (<-chan error, <-chan bool) {
	links := parent.FindLinksByType(depType)
	if links != nil && len(links) > 0 {
		dones := []<-chan error{}
		timeouts := []<-chan bool{}
		for _, link := range links {
			node := p.getNode(link.Target)
			done, timeout := p.postTraverseNode(depType, node, template, operation)
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

//Only used to add host node
func (p *Topology) addNode(node *Node) {
	p.Nodes = append(p.Nodes, node)
}

func (p *Topology) putNode(node *Node) {
	old := p.getNode(node.Name)
	if old == nil {
		p.Nodes = append(p.Nodes, node)
	}
}

func (p *Topology) getNode(name string) *Node {
	for _, node := range p.Nodes {
		if node.Name == name {
			return node
		}
	}
	return nil
}

func (p *Topology) getNodes(name string) []*Node {
	nodes := []*Node{}
	for _, node := range p.Nodes {
		if node.Name == name {
			nodes = append(nodes, node)
		}
	}
	return nodes
}
