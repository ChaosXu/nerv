package topology

import (
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/jinzhu/gorm"
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
		NewSlice: func() interface{} {
			return &[]Topology{}
		},
	}
}

type traverseCallback func(node *Node, template *ServiceTemplate) (<-chan error, <-chan bool)

//Topology which has been deployed by the service template
type Topology struct {
	gorm.Model
	Status
	Name     string  `json:"name"`     //topology name
	Template string  `json:"template"` //service template name
	Version  int     `json:"version"`  //service template version
	Nodes    []*Node `json:"nodes"`
}

//Only used to add host node
func (p *Topology) AddNode(node *Node) {
	p.Nodes = append(p.Nodes, node)
}

func (p *Topology) PutNode(node *Node) {
	old := p.GetNode(node.Name)
	if old == nil {
		p.Nodes = append(p.Nodes, node)
	}
}

func (p *Topology) GetNode(name string) *Node {
	for _, node := range p.Nodes {
		if node.Name == name {
			return node
		}
	}
	return nil
}

func (p *Topology) GetNodes(name string) []*Node {
	nodes := []*Node{}
	for _, node := range p.Nodes {
		if node.Name == name {
			nodes = append(nodes, node)
		}
	}
	return nodes
}
