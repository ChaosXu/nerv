package model

import "fmt"

// NodeType is the metadata of the node
type NodeType struct {
	Name       string                //The name of NodeType
	Base       string                //Base type name
	Operations map[string]*Operation //Operation of type
}


// Operation is action of type
type Operation struct {
	Implementor string //Function implement the operation
}

func (p *NodeType) Invoke(name string, node *Node) (NodeStatus, error) {
	return NodeStatusRed, fmt.Errorf("TBD")
}


