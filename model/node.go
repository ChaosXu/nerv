package model

import "github.com/chaosxu/nerv/template"

type NodeStatus int

const (
	NodeStatusNew = iota //when new
	NodeStatusGreen        //all element ok
	NodeStatusYellow    //some ok,some failed
	NodeStatusRed        //all element failed
)

//Node is element of topology
type Node struct {
	Template *template.NodeTemplate
	Links    map[string][]*Node
	Status   NodeStatus //status of node
	Error    error      //error during the node process
}
