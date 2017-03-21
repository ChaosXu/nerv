package repository

import (
	"encoding/json"
	"fmt"
	"github.com/ChaosXu/nerv/lib/automation/model/topology"
	"github.com/ChaosXu/nerv/lib/env"
	"github.com/go-resty/resty"
)

// HttpTemplateRepository load template from remote server
type HttpTemplateRepository struct{}

func (p *HttpTemplateRepository) GetTemplate(path string) (*topology.ServiceTemplate, error) {
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
