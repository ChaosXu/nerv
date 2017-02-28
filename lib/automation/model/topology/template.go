package topology

import "strings"


// ServiceTemplate is a prototype of service.
type ServiceTemplate struct {
	Path        string
	Name        string             `json:"name"`
	Version     int32              `json:"version"`
	Environment string             `json:"environment"` //standalone|distributed
	Inputs      []Input            `json:"inputs"`
	Nodes       []NodeTemplate     `json:"nodes"`
}

// Input define arguments that used by nodes
type Input struct {
	Name  string
	Type  string
	Value string
}

// NodeTemplate is a prototype of service node.
type NodeTemplate struct {
	Name         string        `json:"name"`        //Node name
	Type         string        `json:"type"`        //The name of NodeType
	Parameters   []Parameter `json:parameters`      //parameters of NodeTemplate
	Dependencies []Dependency `json:"dependencies"` //The dependencies of node
}

func (p *NodeTemplate) formatParameterValue(name string, ctx *Context) interface{} {
	var pv string
	for _, param := range p.Parameters {
		if param.Name == name {
			pv = param.Value
			break
		}
	}
	return ctx.FormatValue(pv)
}

func (p *NodeTemplate) getParameterValue(name string, ctx *Context) interface{} {
	var pv string
	for _, param := range p.Parameters {
		if param.Name == name {
			pv = param.Value
			break
		}
	}
	if pv == "" {
		return ""
	}
	if p.isVar(pv) {
		return ctx.GetValue(pv)
	} else {
		return pv
	}
}

func (p *NodeTemplate) isVar(pv string) bool {
	return strings.Index(pv, "${") >= 0 && strings.LastIndex(pv, "}") > 0
}

// Dependency is relationship  between two node
type Dependency struct {
											 //	gorm.Model
	NodeTemplateID int                       //`gorm:"index"`  //Foreign key of the node template
	Type           string    `json:"type"`   //The type of dependency: connect;contained
	Target         string    `json:"target"` //The name of target node
}

// Parameter is used to generate the node of template
type Parameter struct {
					   //	gorm.Model
	NodeTemplateID int //`gorm:"index"` //Foreign key of the node template
	Name           string    `json:"name"`
	Value          string    `json:"value"`
}


// CreateTopology create a topology by the service template.
func (p *ServiceTemplate) NewTopology(name string, version int, ctx *Context) *Topology {

	topology := &Topology{Name:name, Template:p.Path, Version:version, Nodes:[]*Node{}}

	for _, template := range p.Nodes {
		p.createNode(&template, topology, ctx)
	}

	return topology;
}

func (p *ServiceTemplate) createNode(nodeTemplate *NodeTemplate, topology *Topology, ctx *Context) []*Node {
	deps := nodeTemplate.Dependencies

	targetNodes := []*Node{}
	if deps == nil || len(deps) == 0 {
		targetNodes = topology.GetNodes(nodeTemplate.Name)
		if len(targetNodes) == 0 {
			//TBD: optimize
			if nodeTemplate.Type == "/nerv/cluster/Host" {
				targetNodes = newNodesByHostTemplate(nodeTemplate, ctx)
			} else if nodeTemplate.Type == "/nerv/compute/Host" {
				targetNodes = newNodesByHostTemplate(nodeTemplate, ctx)
			} else {
				targetNodes = append(targetNodes, newNodeByTemplate(nodeTemplate, ctx))
			}
			for _, targetNode := range targetNodes {
				topology.AddNode(targetNode)
			}
		}
		return targetNodes
	}

	sourceNodes := []*Node{}
	for _, dep := range deps {
		if dep.Type == "contained" {
			targetTemplate := p.FindTemplate(dep.Target)
			targetNodes = p.createNode(targetTemplate, topology, ctx)
			for _, targetNode := range targetNodes {
				sourceNode := &Node{
					Name:nodeTemplate.Name,
					Template:nodeTemplate.Name,
					Class:nodeTemplate.Type,
					Address:targetNode.Address,
					Credential:targetNode.Credential,
					Links:[]*Link{},
					Properties:newConfigs(nodeTemplate, ctx),
					Status:Status{RunStatus:RunStatusNone},
				}
				sourceNode.Link(dep.Type, targetNode.Name)
				sourceNodes = append(sourceNodes, sourceNode)
				topology.AddNode(sourceNode)
			}
			return sourceNodes
		}
	}
	return sourceNodes
}

// FindNode return the node
func (p *ServiceTemplate) FindTemplate(name string) *NodeTemplate {
	for _, node := range p.Nodes {
		if node.Name == name {
			return &node
		}
	}
	return nil
}
