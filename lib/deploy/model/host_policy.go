package model

import "strings"

func createNodesByHostTemplate(nodeTemplate *NodeTemplate) []*Node {
	addrs := nodeTemplate.getParameterValue("addresses")
	credential := nodeTemplate.getParameterValue("credential")
	nodes := []*Node{}
	for _, addr := range strings.Split(addrs, ",") {
		node := &Node{
			Name:nodeTemplate.Name,
			Template:nodeTemplate.Name,
			Address:addr,
			Credential: credential,
			Links:[]*Link{},
			Status:Status{RunStatus:RunStatusNone},
		}
		nodes = append(nodes, node)
	}
	return nodes
}

func createNodesByECHostTemplate(nodeTemplate *NodeTemplate) []*Node {
	addrs := nodeTemplate.getParameterValue("addresses")

	nodes := []*Node{}
	for _, addr := range strings.Split(addrs, ",") {
		node := &Node{
			Name:nodeTemplate.Name,
			Template:nodeTemplate.Name,
			Address:addr,
			Links:[]*Link{},
			Status:Status{RunStatus:RunStatusNone},
		}
		nodes = append(nodes, node)
	}
	return nodes
}
