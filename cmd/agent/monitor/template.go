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

func LoadMonitorTemplates(path string) []*model.MonitorTemplate {
	return nil
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

			log.Printf("DiscoveryTemplates:%s", path)
		}
		return nil
	})
	return templates, err
}

func readMetric(resourceType string, metricName string) (*model.Metric, error) {
	root := env.Config().GetMapString("metrics", "path", "../config/metrics")
	file := path.Join(root, strings.ToLower(resourceType), metricName) + ".json"
	metric := &model.Metric{}
	err := json.FromPath(file, metric)
	if err != nil {
		return nil, err
	} else {
		return metric, nil
	}
}