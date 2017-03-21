package topology

// newNodesByHostTemplate create host nodes for nerv worker cluster
func newNodesByHostTemplate(nodeTemplate *NodeTemplate, ctx *Context) []*Node {
	configs := newConfigs(nodeTemplate, ctx)
	addrs := nodeTemplate.getParameterValue("address", ctx)
	credential := nodeTemplate.getParameterValue("credential", ctx)

	nodes := []*Node{}
	if addrs == nil {
		return nodes
	}
	//BUG? if addr is constant that wiil return empty nodes
	ipList, ok := addrs.([]interface{})
	if !ok {
		return nodes
	}
	cre, ok := credential.(string)
	if !ok {
		return nodes
	}
	for _, addr := range ipList {
		ip, ok := addr.(string)
		if ok {
			node := &Node{
				Name:       nodeTemplate.Name,
				Template:   nodeTemplate.Name,
				Class:      nodeTemplate.Type,
				Address:    ip,
				Credential: cre,
				Links:      []*Link{},
				Properties: configs,
				Status:     Status{RunStatus: RunStatusNone},
			}
			nodes = append(nodes, node)
		}
	}
	return nodes
}

//func newNodesByECHostTemplate(nodeTemplate *NodeTemplate) []*Node {
//	configs := newConfigs(nodeTemplate)
//	addrs := nodeTemplate.getParameterValue("addresses")
//
//	nodes := []*Node{}
//	for _, addr := range strings.Split(addrs, ",") {
//		node := &Node{
//			Name:nodeTemplate.Name,
//			Template:nodeTemplate.Name,
//			Class:nodeTemplate.Type,
//			Address:addr,
//			Links:[]*Link{},
//			Properties:configs,
//			Status:Status{RunStatus:RunStatusNone},
//		}
//		nodes = append(nodes, node)
//	}
//	return nodes
//}

func newNodeByTemplate(nodeTemplate *NodeTemplate, ctx *Context) *Node {
	configs := newConfigs(nodeTemplate, ctx)
	return &Node{Name: nodeTemplate.Name, Template: nodeTemplate.Name, Class: nodeTemplate.Type, Links: []*Link{}, Properties: configs, Status: Status{RunStatus: RunStatusNone}}
}

func newConfigs(nodeTempalte *NodeTemplate, ctx *Context) []*Property {
	var configs []*Property
	if nodeTempalte.Parameters == nil {
		return configs
	}

	for _, param := range nodeTempalte.Parameters {
		//TBD: don't hard code
		if param.Name == "address" {
			continue
		}
		v := nodeTempalte.formatParameterValue(param.Name, ctx)
		if v != nil {
			value, ok := v.(string)
			if ok {
				config := &Property{Key: param.Name, Value: value}
				configs = append(configs, config)
			}
		}
	}

	return configs
}
