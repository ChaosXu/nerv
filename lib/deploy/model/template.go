package model

import (
	"github.com/jinzhu/gorm"
	"github.com/ChaosXu/nerv/lib/db"
)

func init() {
	db.Models["ServiceTemplate"] = templateDesc()
	db.Models["NodeTemplate"] = nodeTemplateDesc()
	db.Models["Dependency"] = dependencyDesc()
}

func templateDesc() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &ServiceTemplate{},
		New: func() interface{} {
			return &ServiceTemplate{}
		},
		NewSlice:func() interface{} {
			return &[]ServiceTemplate{}
		},
	}
}

func nodeTemplateDesc() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &NodeTemplate{},
		New: func() interface{} {
			return &NodeTemplate{}
		},
		NewSlice:func() interface{} {
			return &[]NodeTemplate{}
		},
	}
}

func dependencyDesc() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &Dependency{},
		New: func() interface{} {
			return &Dependency{}
		},
		NewSlice:func() interface{} {
			return &[]Dependency{}
		},
	}
}

// Dependency is relationship bettwen two node
type Dependency struct {
	gorm.Model
	NodeTemplateID int       `gorm:"index"`  //Foreign key of the node template
	Type           string    `json:"type"`   //The type of dependency: connect;contained
	Target         string    `json:"target"` //The name of target node
}

// NodeTemplate is a prototype of service node.
type NodeTemplate struct {
	gorm.Model
	ServiceTemplateID int           `gorm:"index"`       //Foreign key of the service template
	Name              string        `json:"name"`        //Node name
	Type              string        `json:"type"`        //The name of NodeType
	Dependencies      []Dependency `json:"dependencies"` //The dependencies of node
}

// ServiceTemplate is a prototype of service.
type ServiceTemplate struct {
	gorm.Model
	Name    string           `json:"name"`
	Version int32            `json:"version"`
	Nodes   []NodeTemplate   `json:"nodes"`
}

// CreateTopology create a topology by the service template.
func (p *ServiceTemplate) CreateTopology(name string) (*Topology, error) {
	tnodes := []NodeTemplate{}
	db.DB.Where("service_template_id =? ", p.ID).Preload("Dependencies").Find(&tnodes)
	p.Nodes = tnodes

	topology := newTopology(p, name)
	if err := db.DB.Create(topology).Error; err != nil {
		return nil, err
	} else {
		return topology, nil
	}
}

// FindNode return the node
func (p *ServiceTemplate) FindNode(name string) *NodeTemplate {
	for _, node := range p.Nodes {
		if node.Name==name {
			return &node
		}
	}
	return nil
}