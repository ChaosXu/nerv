package repository

import (
	"fmt"
	"github.com/go-resty/resty"
	"github.com/ChaosXu/nerv/lib/deploy/model/topology"
	"github.com/ChaosXu/nerv/lib/env"
	"encoding/json"
	"github.com/toolkits/file"
)


// TemplateRepository manage all templates
type TemplateRepository interface {
	// GetTemplate load template from storage
	GetTemplate(path string) (*topology.ServiceTemplate, error)
}

// RemoteTemplateRepository load template from remote server
type RemoteTemplateRepository struct{}

func (p *RemoteTemplateRepository) GetTemplate(path string) (*topology.ServiceTemplate, error) {
	baseUrl := env.Config().GetMapString("templates", "repository", "http://localhost:3332/api/templates")
	url := fmt.Sprintf("%s%s", baseUrl, path)
	res, err := resty.R().
			SetHeader("Content-Type", "application/json").
			Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		return nil, fmt.Errorf("%s", string(res.Body()))
	}
	template := &topology.ServiceTemplate{}
	if err = json.Unmarshal(res.Body(), template); err != nil {
		return nil, err
	}
	template.Path = path
	return template, nil
}

// LocalTemplateRepository load template from local file system.
type LocalTemplateRepository struct{}

func (p *LocalTemplateRepository) GetTemplate(path string) (*topology.ServiceTemplate, error) {
	content, err := file.ToBytes(path)
	if err != nil {
		return nil, err
	}

	template := &topology.ServiceTemplate{}
	if err = json.Unmarshal(content, template); err != nil {
		return nil, err
	}
	template.Path = path
	return template, nil
}
