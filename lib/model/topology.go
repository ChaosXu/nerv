package model

import (
	"github.com/jinzhu/gorm"
	"github.com/chaosxu/nerv/lib/db"
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


//Topology which has been deployed by the service template
type Topology struct {
	gorm.Model
	Name  string
	Nodes []Node
}

//NewTopology create a topology from service template
func newTopology(t *ServiceTemplate, name string) *Topology {
	topology := &Topology{Name:name, Nodes:[]Node{}}

	tnodes := []NodeTemplate{}
	db.DB.Where("service_template_id =? ", t.ID).Preload("Dependencies").Find(&tnodes)
	tnodeMap := map[string]*NodeTemplate{}

	for _, tnode := range tnodes {
		node := &Node{Name:tnode.Name, Template:tnode.Name, Links:[]Link{}, Status:NodeStatusNew}
		topology.putNode(node)
		tnodeMap[tnode.Name] = &tnode
	}

	for _, tnode := range tnodes {
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
				//TBD: BUG
				sourceNode.link(dep.Type, targetNode.Name)
				targetNodeTemplate := tnodeMap[targetNode.Template]
				if targetNodeTemplate != nil {
					traverse(tnodeMap, targetNodeTemplate, topology)
				}
			}
		}
	}
}

//Install the topology and start to serve
func (p *Topology) Install() {
	for _, node := range p.Nodes {
		p.postTraverse("contained", &node, p.installNode)
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

func (p *Topology) installNode(node *Node) {
	//nt := node.Template
	//class := GetClassRepository().Find(nt.Type)
	//if class != nil {
	//	status, err := class.Invoke("create", node)
	//	if err != nil {
	//		node.Error = err
	//	}
	//	node.Status = status
	//} else {
	//	node.Status = NodeStatusRed
	//	node.Error = fmt.Errorf("type %s of node %s isn't exist", nt.Type, nt.Name)
	//}
}

func (p *Topology) postTraverse(depType string, parent *Node, fn func(node *Node)) {
	links := parent.findLinksByType(depType)
	if links != nil {
		for _, link := range links {
			node := p.getNode(link.Target)
			p.postTraverse(depType, node, fn)
		}
	}
	fn(parent)
}

func (p *Topology) putNode(node *Node) {
	old := p.getNode(node.Name)
	if old == nil {
		p.Nodes = append(p.Nodes, *node)
	}
}

func (p *Topology) getNode(name string) *Node {
	for _, node := range p.Nodes {
		if node.Name == name {
			return &node
		}
	}
	return nil
}
