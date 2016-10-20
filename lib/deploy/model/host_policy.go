package model

import "strings"

func createNodesByHostTemplate(nodeTemplate *NodeTemplate) []*Node {
	addrs := nodeTemplate.getParameterValue("addresses")
	credential := nodeTemplate.getParameterValue("credential")
	nodes := []*Node{}
	for _, addr := range strings.Split(addrs, ",") {
		//TBD:support network
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
