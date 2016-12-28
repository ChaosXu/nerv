package topology

import (
	"github.com/go-resty/resty"
	"fmt"
	"github.com/ChaosXu/nerv/lib/env"
	"k8s.io/kubernetes/pkg/util/json"
)

func GetTemplate(path string) (*ServiceTemplate, error) {
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
	template := &ServiceTemplate{}
	if err = json.Unmarshal(res.Body(), template); err != nil {
		return nil, err
	}
	return template, nil
}
