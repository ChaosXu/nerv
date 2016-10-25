package model

import "strings"

func newNodesByHostTemplate(nodeTemplate *NodeTemplate) []*Node {
	configs:=newConfigs(nodeTemplate)
	addrs := nodeTemplate.getParameterValue("addresses")
	credential := nodeTemplate.getParameterValue("credential")

	nodes := []*Node{}
	for _, addr := range strings.Split(addrs, ",") {
		node := &Node{
			Name:nodeTemplate.Name,
			Template:nodeTemplate.Name,
			Class:nodeTemplate.Type,
			Address:addr,
			Credential: credential,
			Links:[]*Link{},
			Properties:configs,
			Status:Status{RunStatus:RunStatusNone},
		}
		nodes = append(nodes, node)
	}
	return nodes
}

func newNodesByECHostTemplate(nodeTemplate *NodeTemplate) []*Node {
	configs := newConfigs(nodeTemplate)
	addrs := nodeTemplate.getParameterValue("addresses")

	nodes := []*Node{}
	for _, addr := range strings.Split(addrs, ",") {
		node := &Node{
			Name:nodeTemplate.Name,
			Template:nodeTemplate.Name,
			Class:nodeTemplate.Type,
			Address:addr,
			Links:[]*Link{},
			Properties:configs,
			Status:Status{RunStatus:RunStatusNone},
		}
		nodes = append(nodes, node)
	}
	return nodes
}

func newNodeByTemplate(nodeTemplate *NodeTemplate) *Node {
	configs := newConfigs(nodeTemplate)
	return &Node{Name:nodeTemplate.Name, Template:nodeTemplate.Name, Class:nodeTemplate.Type, Links:[]*Link{}, Properties:configs, Status:Status{RunStatus:RunStatusNone}}
}

func newConfigs(nodeTempalte *NodeTemplate) []*Property {
	var configs []*Property
	if nodeTempalte.Parameters == nil {
		return configs
	}

	for _, param := range nodeTempalte.Parameters {
		config := &Property{Key:param.Name, Value:param.Value}
		configs = append(configs, config)
	}

	return configs;
}