package driver

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/ChaosXu/nerv/lib/deploy/driver/ssh"
	"github.com/ChaosXu/nerv/lib/deploy/driver/shell"
)

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

// Class is the metadata of the node
type Class struct {
	gorm.Model
	Name       string      `gorm:"unique"` //The name of NodeType
	Base       string                      //Base type name
	Operations []Operation                 //Operation of type
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
func (p *Class) Invoke(operation string, address string, credential string, args map[string]string) error {
	op := p.findOperation(operation)
	if op == nil {
		return fmt.Errorf("unsupported operation %s", operation)
	}
	switch op.Type {
	case "ssh":
		return ssh.Execute(address, op.Implementor, args, credential)
	case "shell":
		return shell.Execute(address, op.Implementor, args)
	//case "go":
	//	m := golang.Models
	//	res := m[op.Implementor]
	//	if res == nil {
	//		return fmt.Errorf("TBD operation type %s", op.Type)
	//	} else {
	//		res.Create()
	//		return nil
	//	}

	default:
		return fmt.Errorf("unsupported operation type %s", op.Type)
	}
	return fmt.Errorf("invoke%s %s", operation, address)
}

func (p *Class)findOperation(opName string) *Operation {
	for _, op := range p.Operations {
		if op.Name == opName {
			return &op
		}
	}
	return nil
}






