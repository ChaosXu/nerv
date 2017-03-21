package repository

import (
	"encoding/json"
	"fmt"
	"github.com/ChaosXu/nerv/lib/automation/model/topology"
	"github.com/toolkits/file"
)

// FileTemplateRepository load template from local file system.
type FileTemplateRepository struct{}

func (p *FileTemplateRepository) GetTemplate(path string) (*topology.ServiceTemplate, error) {
	content, err := file.ToBytes(path)
	if err != nil {
		return nil, fmt.Errorf("load file failed. file: %s. %s", path, err.Error())
	}

	template := &topology.ServiceTemplate{}
	if err = json.Unmarshal(content, template); err != nil {
		return nil, err
	}
	template.Path = path
	return template, nil
}
