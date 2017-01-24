package topology


// ServiceTemplate is a prototype of service.
type ServiceTemplate struct {
	Path        string
	Name        string           	`json:"name"`
	Version     int32            	`json:"version"`
	Environment string         		`json:"environment"` //standalone|distributed
	Nodes       []NodeTemplate   	`json:"nodes"`
}

// NodeTemplate is a prototype of service node.
type NodeTemplate struct {
	Name         string        `json:"name"`        //Node name
	Type         string        `json:"type"`        //The name of NodeType
	Parameters   []Parameter `json:parameters`      //parameters of NodeTemplate
	Dependencies []Dependency `json:"dependencies"` //The dependencies of node
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
func (p *ServiceTemplate) NewTopology(name string) *Topology {

	topology := &Topology{Name:name, Template:p.Path, Version:p.Version, Nodes:[]*Node{}}

	for _, template := range p.Nodes {
		p.createNode(&template, topology)
	}

	return topology;
}

func (p *ServiceTemplate) createNode(nodeTemplate *NodeTemplate, topology *Topology) []*Node {
	deps := nodeTemplate.Dependencies

	targetNodes := []*Node{}
	if deps == nil || len(deps) == 0 {
		targetNodes = topology.GetNodes(nodeTemplate.Name)
		if len(targetNodes) == 0 {
			//TBD: optimize
			if nodeTemplate.Type == "/nerv/cluster/Host" {
				targetNodes = newNodesByHostTemplate(nodeTemplate)
			} else if nodeTemplate.Type == "/nerv/ECHost" {
				targetNodes = newNodesByECHostTemplate(nodeTemplate)
			} else {
				targetNodes = append(targetNodes, newNodeByTemplate(nodeTemplate))
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
			targetNodes = p.createNode(targetTemplate, topology)
			for _, targetNode := range targetNodes {
				sourceNode := &Node{
					Name:nodeTemplate.Name,
					Template:nodeTemplate.Name,
					Class:nodeTemplate.Type,
					Address:targetNode.Address,
					Credential:targetNode.Credential,
					Links:[]*Link{},
					Properties:newConfigs(nodeTemplate),
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
