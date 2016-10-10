package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/chaosxu/nerv/lib/db"
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
	gorm.Model
	ClassID     int        `gorm:"index"` //Foreign key of the Class
	Name        string                    //Operation name
	Type        string                    //Operation type.eg.shell
	Implementor string                    //Function implement the operation
}

func (p *Class) Invoke(name string, node *Node, template *NodeTemplate) (NodeStatus, error) {
	return NodeStatusRed, fmt.Errorf("TBD")
}

func init() {
	db.Models["Class"] = classDesc()
	db.Models["Operation"] = operationDesc()
}

func operationDesc() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &Operation{},
		New: func() interface{} {
			return &Operation{}
		},
		NewSlice:func() interface{} {
			return &[]Operation{}
		},
	}
}

func classDesc() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &Class{},
		New: func() interface{} {
			return &Class{}
		},
		NewSlice:func() interface{} {
			return &[]Class{}
		},
	}
}




