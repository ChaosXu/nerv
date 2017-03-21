package env

import (
	"encoding/json"
	"fmt"
	"github.com/toolkits/file"
)

//Properties read configuration from file path
type Properties struct {
	data map[string]interface{}
}

func NewProperties(data map[string]interface{}) *Properties {
	return &Properties{data: data}
}

//LoadConfig return a pointer of the configuration from the json file in the path
func LoadConfig(path string) (*Properties, error) {
	if !file.IsExist(path) {
		return nil, fmt.Errorf("file: %s isn't exists", path)
	}

	body, err := file.ToTrimString(path)
	if err != nil {
		return nil, err
	}

	config := NewProperties(map[string]interface{}{})
	err = json.Unmarshal([]byte(body), &config.data)
	if err != nil {
		return nil, err
	} else {
		return config, nil
	}
}

//GetString return string from config
func (p *Properties) GetString(name string, value ...string) string {
	if r := p.data[name]; r != nil {
		return r.(string)
	} else if len(value) > 0 {
		return value[0]
	} else {
		return ""
	}
}

//GetMapString return string from map in the config
func (p *Properties) GetMapString(name string, field string, value ...string) string {
	if r := p.data[name]; r != nil {
		if v := r.(map[string]interface{})[field]; v != nil {
			return v.(string)
		}
	}

	if len(value) > 0 {
		return value[0]
	} else {
		return ""
	}
}

//GetMap return map from the config
func (p *Properties) GetMap(name string) map[string]interface{} {
	if r := p.data[name]; r != nil {
		return r.(map[string]interface{})
	} else {
		return map[string]interface{}{}
	}
}
