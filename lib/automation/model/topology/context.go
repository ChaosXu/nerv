package topology

import (
	"k8s.io/kubernetes/pkg/util/json"
	"fmt"
	"regexp"
)

type Context struct {
	inputsDef []Input
	inputs    map[string]interface{}
}

func NewContext(def []Input, inputs map[string]interface{}) *Context {
	if len(def) > 0 && inputs != nil {
		buf, _ := json.Marshal(inputs)
		if buf != nil {
			fmt.Println(string(buf))
		}
	}
	return &Context{inputsDef:def, inputs:inputs}
}

func (p *Context) GetValue(name string) interface{} {
	reg := regexp.MustCompile(`^\$\{(.+)\}$`)
	match := reg.FindStringSubmatch(name)
	if len(match) < 1 {
		return nil
	}
	key := match[1]
	//fmt.Printf("GetValue %s\n", key)
	if key == "" {
		return nil
	} else {
		return p.inputs[key]
	}
}
