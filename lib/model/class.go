package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//func init() {
//	db.Models["Class"] = classDesc()
//	db.Models["Operation"] = operationDesc()
//}
//
//func operationDesc() *db.ModelDescriptor {
//	return &db.ModelDescriptor{
//		Type: &Operation{},
//		New: func() interface{} {
//			return &Operation{}
//		},
//		NewSlice:func() interface{} {
//			return &[]Operation{}
//		},
//	}
//}
//
//func classDesc() *db.ModelDescriptor {
//	return &db.ModelDescriptor{
//		Type: &Class{},
//		New: func() interface{} {
//			return &Class{}
//		},
//		NewSlice:func() interface{} {
//			return &[]Class{}
//		},
//	}
//}

// Class is the metadata of the node
type Class struct {
	gorm.Model
	Name       string      //The name of NodeType
	Base       string      //Base type name
	Operations []Operation //Operation of type
}

// Operation is action of type
type Operation struct {
	gorm.Model
	ClassID     int        `gorm:"index"` //Foreign key of the Class
	Name        string                    //Operation name
	Type        string                    //Operation type.eg.shell
	Implementor string                    //Function implement the operation
}

// Invoke the operation of template's type on node
func (p *Class) Invoke(operation string, node *Node, template *NodeTemplate) (NodeStatus, error) {
	op := p.findOperation(operation)
	if op == nil {
		return NodeStatusRed, fmt.Errorf("unsupported operation %s", operation)
	}
	switch op.Type {
	case "shell":
		return NodeStatusRed, fmt.Errorf("TBD operation type %s", op.Type)
	case "go":
		return NodeStatusRed, fmt.Errorf("TBD operation type %s", op.Type)
	default:
		return NodeStatusRed, fmt.Errorf("unsupported operation type %s", op.Type)
	}
	return NodeStatusRed, fmt.Errorf("invoke%s %s %s", operation, node.Name, template.Type)
}

func (p *Class)findOperation(opName string) *Operation {
	for _, op := range p.Operations {
		if op.Name == opName {
			return &op
		}
	}
	return nil
}






