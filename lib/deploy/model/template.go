package model

import (
	"github.com/jinzhu/gorm"
	"github.com/ChaosXu/nerv/lib/db"
)

func init() {
	db.Models["ServiceTemplate"] = templateDesc()
	db.Models["NodeTemplate"] = nodeTemplateDesc()
	db.Models["Dependency"] = dependencyDesc()
	db.Models["Parameter"] = parameterDesc()
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

func parameterDesc() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &Parameter{},
		New: func() interface{} {
			return &Parameter{}
		},
		NewSlice:func() interface{} {
			return &[]Parameter{}
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

// ServiceTemplate is a prototype of service.
type ServiceTemplate struct {
	gorm.Model
	Name    string           `json:"name"`
	Version int32            `json:"version"`
	Nodes   []NodeTemplate   `json:"nodes"`
}

// NodeTemplate is a prototype of service node.
type NodeTemplate struct {
	gorm.Model
	ServiceTemplateID int           `gorm:"index"`       //Foreign key of the service template
	Name              string        `json:"name"`        //Node name
	Type              string        `json:"type"`        //The name of NodeType
	Parameters        []Parameter `json:parameters`      //parameters of NodeTemplate
	Dependencies      []Dependency `json:"dependencies"` //The dependencies of node
}

func (p *NodeTemplate) getParameterValue(name string) string {
	for _, param := range p.Parameters {
		if param.Name == name {
			return param.Value
		}
	}
	return ""
}

// Dependency is relationship  between two node
type Dependency struct {
	gorm.Model
	NodeTemplateID int       `gorm:"index"`  //Foreign key of the node template
	Type           string    `json:"type"`   //The type of dependency: connect;contained
	Target         string    `json:"target"` //The name of target node
}

// Parameter is used to generate the node of template
type Parameter struct {
	gorm.Model
	NodeTemplateID int       `gorm:"index;unique_index:idx_parameter_idn"` //Foreign key of the node template
	Name           string    `json:"name";gorm:"unique_index:idx_parameter_idn"`
	Value          string    `json:"value"`
}


// CreateTopology create a topology by the service template.
func (p *ServiceTemplate) CreateTopology(name string) (*Topology, error) {
	nodeTemplates := []NodeTemplate{}
	db.DB.Where("service_template_id =? ", p.ID).Preload("Dependencies").Preload("Parameters").Find(&nodeTemplates)
	p.Nodes = nodeTemplates

	topology := &Topology{Name:name, Template:p.Name, Version:p.Version, Nodes:[]*Node{}}

	for _, template := range p.Nodes {
		p.createNode(&template, topology)
	}

	if err := db.DB.Create(topology).Error; err != nil {
		return nil, err
	} else {
		return topology, nil
	}
}

func (p *ServiceTemplate) createNode(nodeTemplate *NodeTemplate, topology *Topology) []*Node {
	deps := nodeTemplate.Dependencies

	targetNodes := []*Node{}
	if deps == nil || len(deps) == 0 {
		targetNodes = topology.getNodes(nodeTemplate.Name)
		if len(targetNodes) == 0 {
			if nodeTemplate.Type == "/nerv/Host" {
				targetNodes = createNodesByHostTemplate(nodeTemplate)
			} else {
				targetNodes = append(targetNodes, &Node{Name:nodeTemplate.Name, Template:nodeTemplate.Name, Links:[]*Link{}, Status:Status{RunStatus:RunStatusNone}})
			}
			for _, targetNode := range targetNodes {
				topology.addNode(targetNode)
			}
		}
		return targetNodes
	}

	sourceNodes := []*Node{}
	for _, dep := range deps {
		if dep.Type == "contained" {
			targetTemplate := p.findTemplate(dep.Target)
			targetNodes = p.createNode(targetTemplate, topology)
			for _, targetNode := range targetNodes {
				sourceNode := &Node{
					Name:nodeTemplate.Name,
					Template:nodeTemplate.Name,
					Address:targetNode.Address,
					Credential:targetNode.Credential,
					Links:[]*Link{},
					Status:Status{RunStatus:RunStatusNone},
				}
				sourceNode.Link(dep.Type, targetNode.Name)
				sourceNodes = append(sourceNodes, sourceNode)
				topology.addNode(sourceNode)
			}
			return sourceNodes
		}
	}
	return sourceNodes
}

// FindNode return the node
func (p *ServiceTemplate) findTemplate(name string) *NodeTemplate {
	for _, node := range p.Nodes {
		if node.Name == name {
			return &node
		}
	}
	return nil
}
