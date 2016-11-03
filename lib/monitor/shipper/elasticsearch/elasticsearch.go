package elasticsearch

import (
	"log"
	"fmt"
	"encoding/json"
	"github.com/ChaosXu/nerv/lib/env"
	"github.com/ChaosXu/nerv/lib/monitor/model"
	"github.com/ChaosXu/nerv/lib/monitor/shipper"
	"github.com/go-resty/resty"
	"time"
)

type ElasticsearchShipper struct {
	server string
	cfg    *env.Properties
}

func NewShipper(cfg *env.Properties) shipper.Shipper {
	address := cfg.GetMapString("shipper", "server", "3334")
	return &ElasticsearchShipper{server:address, cfg:cfg}
}

func (p *ElasticsearchShipper) Init() error {
	return nil
}

func (p *ElasticsearchShipper) Send(v interface{}) {
	switch v.(type) {
	case *model.Sample:
		s, _ := v.(*model.Sample)
		p.sendSample(s)
	case *model.Resource:
		log.Printf("TDB:send resource: %+v\n", v)
	}
}

func (p *ElasticsearchShipper)sendSample(sample *model.Sample) {
	body, err := json.Marshal(sample)
	if err != nil {
		log.Printf("send sample error. %s\n", err.Error())
	}

	resType := sample.Tags["resourceType"]
	metric := sample.Metric
	templateName := getTemplateName(resType, metric)

	res, err := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(string(body)).
			Post(fmt.Sprintf("http://%s/%s/%s?pretty=true", p.server, p.index(templateName), templateName))
	if err != nil {
		log.Printf("create schema error. %s %s \n%s", resType, metric, err.Error())
	} else if res.StatusCode() >= 200 && res.StatusCode() < 300 {
		log.Printf("create schema. %s %s %d \n%s", resType, metric, res.StatusCode(), string(res.Body()))
	} else {
		log.Printf("create schema error. %s %s %d \n%s", resType, metric, res.StatusCode(), string(res.Body()))
	}
}

func (p *ElasticsearchShipper)index(template string) string {
	year, month, day := time.Now().Date()
	return fmt.Sprintf("%s_%d_%d_%d", template, year, month, day)
}

