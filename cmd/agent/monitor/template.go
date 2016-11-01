package monitor

import (
	"github.com/ChaosXu/nerv/lib/monitor/model"
	"path/filepath"
	"os"
	"github.com/toolkits/file"
	"log"
	"github.com/ChaosXu/nerv/lib/json"
	"github.com/ChaosXu/nerv/lib/env"
	"path"
	"strings"
)

func LoadMonitorTemplates(path string) ([]*model.MonitorTemplate, error) {
	templates := []*model.MonitorTemplate{}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if file.Ext(path) == ".json" {
			template := &model.MonitorTemplate{}
			if err := json.FromPath(path, template); err != nil {
				return err
			}
			templates = append(templates, template)

			log.Printf("LoadMonitorTemplates: %s %s", template.ResourceType, path)
		}
		return nil
	})
	return templates, err
}

func LoadDiscoveryTemplates(path string) ([]*model.DiscoveryTemplate, error) {
	templates := []*model.DiscoveryTemplate{}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if file.Ext(path) == ".json" {
			template := &model.DiscoveryTemplate{}
			if err := json.FromPath(path, template); err != nil {
				return err
			}
			templates = append(templates, template)

			log.Printf("LoadDiscoveryTemplates: %s %s", template.ResourceType, path)
		}
		return nil
	})
	return templates, err
}

func loadMetric(cfg *env.Properties,resourceType string, metricName string) (*model.Metric, error) {
	root := cfg.GetMapString("metrics", "path", "../config/metrics")
	file := path.Join(root, strings.ToLower(resourceType), metricName) + ".json"
	metric := &model.Metric{}
	err := json.FromPath(file, metric)
	if err != nil {
		return nil, err
	} else {
		return metric, nil
	}
}