package model

// Class is the metadata of the node
type Class struct {
	Name       string      //The name of NodeType
	Base       string      //Base type name
	Operations []Operation //Operation of type
}

func (p *Class) GetOperation(name string) *Operation {

	for _, op := range p.Operations {
		if op.Name == name {
			return &op
		}
	}

	return nil
}

// Operation is action of type
type Operation struct {
	DefineClass string ""
	Name        string //Operation name
	Type        string //Operation type.eg.shell
	Implementor string //Function implement the operation
}
