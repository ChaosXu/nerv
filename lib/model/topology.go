package model

import (
	"github.com/jinzhu/gorm"
	"github.com/chaosxu/nerv/lib/db"
	"github.com/chaosxu/nerv/lib/log"
	"fmt"
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
		node := &Node{Name:tnode.Name, Template:tnode.Name, Links:[]*Link{}, Status:NodeStatusNew}
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
	Name     string //topology name
	Template string //service template name
	Version  int32  //service template version
	Nodes    []*Node
}

//Install the topology and start to serve
func (p *Topology) Install() {
	log.LogCodeLine()
	tnodes := []*Node{}
	db.DB.Where("topology_id =? ", p.ID).Preload("Links").Find(&tnodes)
	p.Nodes = tnodes

	template := ServiceTemplate{}
	if err := db.DB.Where("name=? and version=?", p.Template, p.Version).Preload("Nodes").First(&template).Error; err != nil {
		//TBD
		fmt.Sprintln(err.Error())
	}

	for _, node := range p.Nodes {
		p.postTraverse("contained", node, &template, p.installNode)
	}
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

func (p *Topology) installNode(node *Node, template *ServiceTemplate) {
	nodeTemplate := template.FindNode(node.Template)
	if nodeTemplate != nil {
		node.Execute("Create", nodeTemplate)
	} else {
		node.Status = NodeStatusRed
		node.Error = fmt.Sprintf("template %s of node %s isn't exist", node.Template, node.Name)
		db.DB.Save(node)
	}
}

func (p *Topology) postTraverse(depType string, parent *Node, template *ServiceTemplate, fn func(node *Node, template *ServiceTemplate)) {
	links := parent.FindLinksByType(depType)
	if links != nil {
		for _, link := range links {
			node := p.getNode(link.Target)
			p.postTraverse(depType, node, template, fn)
		}
	}
	fn(parent, template)
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
