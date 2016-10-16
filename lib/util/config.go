package util

import (
	"encoding/json"
	"github.com/toolkits/file"
	"fmt"
)

//Config read configuration from file path
type Config struct {
	props map[string]string
}

//LoadConfig return a pointer of the configuration from the json file in the path
func LoadConfig(path string) (*Config, error) {
	if !file.IsExist(path) {
		return nil, fmt.Errorf("file: %s isn't exists", path)
	}

	body, err := file.ToTrimString(path)
	if err != nil {
		return nil, err
	}

	config := &Config{props:map[string]string{}}
	err = json.Unmarshal([]byte(body), &config.props)
	if err != nil {
		return nil, err
	} else {
		return config, nil
	}
}

func (p *Config) GetProperty(name string, value... string) string {
	if r := p.props[name]; r != "" {
		return r
	} else if len(value) > 0 {
		return value[0]
	} else {
		return ""
	}
}
