package model

import (
	"fmt"
)

//Topology which has been deployed by the service template
type Topology struct {
	Nodes map[string]*Node
}

//NewTopology create a topology from service template
func NewTopology(t *ServiceTemplate) *Topology {
	topology := &Topology{map[string]*Node{}}

	for _, tnode := range t.Nodes {
		node := &Node{tnode, map[string][]*Node{}, NodeStatusNew, nil}
		topology.Nodes[tnode.Name] = node
	}

	for _, tnode := range t.Nodes {
		traverse(tnode, topology)
	}

	return topology
}

func traverse(tnode *NodeTemplate, topology *Topology) {
	deps := tnode.Dependencies
	if deps != nil || len(deps) == 0 {
		return
	}

	for _, dep := range deps {
		target := dep.Target

		if target != "" {
			targetNode := topology.Nodes[target]
			sourceNode := topology.Nodes[tnode.Name]
			if targetNode != nil && sourceNode != nil {
				depNodes := sourceNode.Links[dep.Type]
				if depNodes == nil || len(depNodes) == 0 {
					depNodes = []*Node{}
					sourceNode.Links[dep.Type] = depNodes
				}
				sourceNode.Links[dep.Type] = append(depNodes, targetNode)
				traverse(targetNode.Template, topology)
			}
		}
	}
}

//Install the topology and start to serve
func (p *Topology) Install() {
	for _, node := range p.Nodes {
		p.postTraverse("contained", node, p.installNode)
	}
}

func (p *Topology) installNode(node *Node) {
	nt := node.Template
	class := GetClassRepository().Find(nt.Type)
	if class != nil {
		status, err := class.Invoke("create", node)
		if err != nil {
			node.Error = err
		}
		node.Status = status
	} else {
		node.Status = NodeStatusRed
		node.Error = fmt.Errorf("type %s of node %s isn't exist", nt.Type, nt.Name)
	}
}

func (p *Topology) postTraverse(depType string, parent *Node, fn func(node *Node)) {
	nodes := parent.Links[depType]
	if nodes != nil {
		for _, node := range nodes {
			p.postTraverse(depType, node, fn)
		}
	}
	fn(parent)
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
