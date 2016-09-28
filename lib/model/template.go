package model

import (
	"fmt"
	"encoding/json"

	"github.com/toolkits/file"
	"github.com/jinzhu/gorm"
)

// Dependency is relationship bettwen two node
type Dependency struct {
	Type   string `json:"type"`   //The type of dependency: connect;contained
	Target string `json:"target"` //The name of target node
}

// NodeTemplate is a prototype of service node.
type NodeTemplate struct {
	Name         string        `json:"name"`         //Node name
	Type         string        `json:"type"`         //The name of NodeType
	Dependencies []*Dependency `json:"dependencies"` //The dependencies of node
}

// ServiceTemplate is a prototype of service.
type ServiceTemplate struct {
	gorm.Model
	Name    string          `json:"name"`
	Version int32           `json:"version"`
	Nodes   []*NodeTemplate `json:"nodes"`
}

// GetServiceTemplate read the json file of service template from path.
func GetServiceTemplate(path string) (*ServiceTemplate, error) {
	if path == "" {
		return nil, fmt.Errorf("empty path")
	}

	if !file.IsExist(path) {
		return nil, fmt.Errorf("file: %s isn't exists", path)
	}

	templateContent, err := file.ToTrimString(path)
	if err != nil {
		return nil, err
	}

	var template ServiceTemplate
	err = json.Unmarshal([]byte(templateContent), &template)
	if err != nil {
		return nil, err
	}

	return &template, nil
}
