package model

import (
	"github.com/jinzhu/gorm"
	"github.com/chaosxu/nerv/lib/db"
	"github.com/chaosxu/nerv/lib/log"
	"fmt"
	"time"
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

//NewTopology create a topology from service template
func newTopology(t *ServiceTemplate, name string) *Topology {
	topology := &Topology{Name:name, Template:t.Name, Version:t.Version, Nodes:[]*Node{}}

	tnodeMap := map[string]*NodeTemplate{}
	for _, tnode := range t.Nodes {
		node := &Node{Name:tnode.Name, Template:tnode.Name, Links:[]*Link{}, Status:Status{RunStatus:RunStatusNone}}
		topology.putNode(node)
		tnodeMap[tnode.Name] = &tnode
	}

	for _, tnode := range t.Nodes {
		traverse(tnodeMap, &tnode, topology)
	}

	return topology
}

func traverse(tnodeMap map[string]*NodeTemplate, tnode *NodeTemplate, topology *Topology) {
	deps := tnode.Dependencies
	if deps == nil || len(deps) == 0 {
		return
	}

	for _, dep := range deps {
		target := dep.Target

		if target != "" {
			targetNode := topology.getNode(target)
			sourceNode := topology.getNode(tnode.Name)
			if targetNode != nil && sourceNode != nil {
				sourceNode.Link(dep.Type, targetNode.Name)
				targetNodeTemplate := tnodeMap[targetNode.Template]
				if targetNodeTemplate != nil {
					traverse(tnodeMap, targetNodeTemplate, topology)
				}
			}
		}
	}
}

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

	lock := GetLock("Topology", p.ID)
	if !lock.TryLock() {
		return fmt.Errorf("topology is installing. ID=%d", p.ID)
	}
	defer lock.Unlock()

	tnodes := []*Node{}
	db.DB.Where("topology_id =?", p.ID).Preload("Links").Find(&tnodes)
	p.Nodes = tnodes

	template := ServiceTemplate{}
	if err := db.DB.Where("name=? and version=?", p.Template, p.Version).Preload("Nodes").First(&template).Error; err != nil {
		//TBD
		fmt.Sprintln(err.Error())
	}

	dones := []<-chan error{}
	timeouts := []<-chan bool{}

	for _, node := range p.Nodes {
		done, timeout := p.postTraverse("contained", node, &template, p.installNode)
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
		p.RunStatus = RunStatusRed
	}
	db.DB.Save(p)
	return nil
}

//Uninstall the topology
func (p *Topology) Uninstall() {

}

//Configure the topology for start
func (p *Topology) Configure() {

}

//Start the Topology
func (p *Topology) Start() {

}

//Stop the Topology
func (p *Topology) Stop() {

}

func (p *Topology) installNode(node *Node, template *ServiceTemplate) (<-chan error, <-chan bool) {
	done := make(chan error, 1)
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(30 * time.Second)
		timeout <- true
	}()

	go func() {
		nodeTemplate := template.FindNode(node.Template)
		if nodeTemplate != nil {
			if err := node.Execute("Create", nodeTemplate); err != nil {
				done <- err
			} else {
				done <- nil
			}
		} else {
			node.RunStatus = RunStatusRed
			node.Error = fmt.Sprintf("template %s of node %s isn't exist", node.Template, node.Name)
			done <- db.DB.Save(node).Error
		}
	}()

	return done, timeout
}

func (p *Topology) postTraverse(depType string, parent *Node, template *ServiceTemplate, fn func(node *Node, template *ServiceTemplate) (<-chan error, <-chan bool)) (<-chan error, <-chan bool) {
	links := parent.FindLinksByType(depType)
	if links != nil && len(links) > 0 {
		dones := []<-chan error{}
		timeouts := []<-chan bool{}
		for _, link := range links {
			node := p.getNode(link.Target)
			done, timeout := p.postTraverse(depType, node, template, fn)
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
			return fn(parent, template)
		} else {
			done := make(chan error, 1)
			timeout := make(chan bool, 1)
			done <- err
			return done, timeout
		}
	} else {
		return fn(parent, template)
	}
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
