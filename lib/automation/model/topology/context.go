package topology

import (
	"encoding/json"
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
	return &Context{inputsDef: def, inputs: inputs}
}

func (p *Context) FormatValue(value string) interface{} {
	reg := regexp.MustCompile(`\$\{(.+)\}`)
	return reg.ReplaceAllStringFunc(value, func(name string) string {
		key := name[2 : len(name)-1]
		v := p.inputs[key]
		if v == nil {
			return p.GetDefaultValue(key)
		} else {
			text, ok := v.(string)
			if ok {
				return text
			} else {
				return ""
			}
		}
	})
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
		ret := p.inputs[key]
		if ret == nil {
			ret = p.GetDefaultValue(key)
		}
		return ret
	}
}

func (p *Context) GetDefaultValue(name string) string {
	for _, input := range p.inputsDef {
		if input.Name == name {
			return input.Value
		}
	}
	return ""
}
