package env

import (
	"encoding/json"
	"github.com/toolkits/file"
	"fmt"
)

//Properties read configuration from file path
type Properties struct {
	data map[string]string
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

	config := &Properties{data:map[string]string{}}
	err = json.Unmarshal([]byte(body), &config.data)
	if err != nil {
		return nil, err
	} else {
		return config, nil
	}
}

func (p *Properties) GetProperty(name string, value... string) string {
	if r := p.data[name]; r != "" {
		return r
	} else if len(value) > 0 {
		return value[0]
	} else {
		return ""
	}
}
