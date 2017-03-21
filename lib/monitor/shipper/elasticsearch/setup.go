package elasticsearch

import (
	"encoding/json"
	"fmt"
	"github.com/ChaosXu/nerv/lib/monitor/model"
	"github.com/go-resty/resty"
	"log"
)

//CreateSchema generate schemas for all metrics in es
func CreateSchemas(server string, metrics []*model.Metric) {

	for _, metric := range metrics {
		templateName := getTemplateName(metric.ResourceType, metric.Name)
		template := metricToTemplate(metric, templateName)

		body, err := json.Marshal(template)
		if err != nil {
			log.Printf("marshal schema error. %s %s %s\n", metric.ResourceType, metric.Name, err.Error())
			continue
		} else {
			log.Printf("metric schema:%s %s\n%s\n", metric.ResourceType, metric.Name, body)
		}

		res, err := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(string(body)).
			Put(fmt.Sprintf("http://%s/_template/%s?pretty=true", server, templateName))
		if err != nil {
			log.Printf("create schema error. %s %s \n%s", metric.ResourceType, metric.Name, err.Error())
		} else if res.StatusCode() != 200 {
			log.Printf("create schema error. %s %s %d \n%s", metric.ResourceType, metric.Name, res.StatusCode(), string(res.Body()))
		} else {
			log.Printf("create schema. %s %s %d \n%s", metric.ResourceType, metric.Name, res.StatusCode(), string(res.Body()))
		}

	}
}

func getTemplateName(resType string, metric string) string {
	//return strings.Replace(strings.ToLower(resType), "/", ".", -1)[1:] + "." + metric
	return "metrics"
}

func metricToTemplate(metric *model.Metric, templateName string) map[string]interface{} {
	template := map[string]interface{}{
		"template": templateName + "_*",
		"order":    0,
		"settings": map[string]string{
			"index.refresh_interval": "5s",
		},
		"mappings": map[string]interface{}{
			templateName: metricToMapping(metric),
		},
	}

	return template
}

func metricToMapping(metric *model.Metric) map[string]interface{} {
	properties := map[string]interface{}{}
	mapping := map[string]interface{}{
		"_source":    map[string]interface{}{"enabled": false},
		"_all":       map[string]interface{}{"enabled": false},
		"properties": properties,
	}
	for _, field := range metric.Fields {
		item := map[string]interface{}{
			"type":       getDataType(field.DataType),
			"doc_values": true,
			"index":      "no",
		}
		properties[field.Name] = item
	}
	return mapping
}
func getDataType(dataType model.MetricDataType) string {
	switch dataType {
	case model.MetricDataTypeDouble:
		return "double"
	case model.MetricDataTypeLong:
		return "long"
	case model.MetricDataTypeBool:
		return "bool"
	default:
		return "string"
	}
}
