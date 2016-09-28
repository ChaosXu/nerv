package model


type NodeStatus int

const (
	NodeStatusNew = iota //when new
	NodeStatusGreen        //all element ok
	NodeStatusYellow    //some ok,some failed
	NodeStatusRed        //all element failed
)

//Node is element of topology
type Node struct {
	Template *NodeTemplate
	Links    map[string][]*Node
	Status   NodeStatus //status of node
	Error    error      //error during the node process
}

//func init() {
//	Models["Class"] = classDesc()
//	Models["Operation"] = operationDesc()
//}
//
//func operationDesc() *ModelDescriptor {
//	return &ModelDescriptor{
//		Type: &Operation{},
//		New: func() interface{} {
//			return &Operation{}
//		},
//	}
//}
//
//func classDesc() *ModelDescriptor {
//	return &ModelDescriptor{
//		Type: &Class{},
//		New: func() interface{} {
//			return &Class{}
//		},
//	}
//}
