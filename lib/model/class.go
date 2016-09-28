package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Class is the metadata of the node
type Class struct {
	gorm.Model
	Name       string      //The name of NodeType
	Base       string      //Base type name
	Operations []Operation //Operation of type
}

// Operation is action of type
type Operation struct {
	ClassID     int        `gorm:"index"` //Foreign key of the Class
	Implementor string                    //Function implement the operation
}

func (p *Class) Invoke(name string, node *Node) (NodeStatus, error) {
	return NodeStatusRed, fmt.Errorf("TBD")
}

func init() {
	Models["Class"] = classDesc()
	Models["Operation"] = operationDesc()
}

func operationDesc() *ModelDescriptor {
	return &ModelDescriptor{
		Type: &Operation{},
		New: func() interface{} {
			return &Operation{}
		},
		NewSlice:func() interface{} {
			return &[]Operation{}
		},
	}
}

func classDesc() *ModelDescriptor {
	return &ModelDescriptor{
		Type: &Class{},
		New: func() interface{} {
			return &Class{}
		},
		NewSlice:func() interface{} {
			return &[]Class{}
		},
	}
}




