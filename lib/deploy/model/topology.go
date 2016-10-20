package model

import (
	"github.com/jinzhu/gorm"
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/ChaosXu/nerv/lib/log"
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
//func newTopology(serviceTemplate *ServiceTemplate, name string) *Topology {
//	topology := &Topology{Name:name, Template:serviceTemplate.Name, Version:serviceTemplate.Version, Nodes:[]*Node{}}
//
//	templates := map[string]*NodeTemplate{}
//	for _, template := range serviceTemplate.Nodes {
//		templates[template.Name] = &template
//		//TBD: don't hard code that generate host from ip pool in node template of the nerv host
//		if template.Type == "/nerv/Host" {
//			nodes := createNodesByHostTemplate(template)
//			for _, node := range nodes {
//				topology.addNode(node)
//			}
//		} else {
//			node := &Node{Name:template.Name, Template:template.Name, Links:[]*Link{}, Status:Status{RunStatus:RunStatusNone}}
//			topology.putNode(node)
//		}
//	}
//
//	for _, template := range serviceTemplate.Nodes {
//		traverse(templates, &template, topology)
//	}
//
//	return topology
//}
//
//func traverse(nodeTemplates map[string]*NodeTemplate, nodeTemplate *NodeTemplate, topology *Topology) {
//	deps := nodeTemplate.Dependencies
//	if deps == nil || len(deps) == 0 {
//		return
//	}
//
//	for _, dep := range deps {
//		target := dep.Target
//
//		if target != "" {
//			targetNodes := topology.getNodes(target)
//			sourceNode := topology.getNode(nodeTemplate.Name)
//			if sourceNode != nil {
//				for _, targetNode := range targetNodes {
//					sourceNode.Link(dep.Type, targetNode.Name)
//					targetNodeTemplate := nodeTemplates[targetNode.Template]
//					if targetNodeTemplate != nil {
//						traverse(nodeTemplates, targetNodeTemplate, topology)
//					}
//				}
//			}
//		}
//	}
//}

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
	if err := db.DB.Where("name=? and version=?", p.Template, p.Version).Preload("Nodes").Preload("Nodes.Parameters").First(&template).Error; err != nil {
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
		nodeTemplate := template.findTemplate(node.Template)
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
