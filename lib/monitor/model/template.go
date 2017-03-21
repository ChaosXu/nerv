package model

import (
	"github.com/ChaosXu/nerv/lib/json"
	"github.com/jinzhu/gorm"
	"github.com/toolkits/file"
	"log"
	"os"
	"path/filepath"
)

//func init() {
//	db.Models["DiscoveryTemplate"] = discoveryTemplate()
//	db.Models["DiscoveryItem"] = discoveryItem()
//	db.Models["MonitorTemplate"] = monitorTemplate()
//	db.Models["MonitorItem"] = monitorItem()
//}
//
//func discoveryTemplate() *db.ModelDescriptor {
//	return &db.ModelDescriptor{
//		Type: &DiscoveryTemplate{},
//		New: func() interface{} {
//			return &DiscoveryTemplate{}
//		},
//		NewSlice:func() interface{} {
//			return &[]DiscoveryTemplate{}
//		},
//	}
//}
//
//func discoveryItem() *db.ModelDescriptor {
//	return &db.ModelDescriptor{
//		Type: &DiscoveryItem{},
//		New: func() interface{} {
//			return &DiscoveryItem{}
//		},
//		NewSlice:func() interface{} {
//			return &[]DiscoveryItem{}
//		},
//	}
//}
//
//func monitorTemplate() *db.ModelDescriptor {
//	return &db.ModelDescriptor{
//		Type: &MonitorTemplate{},
//		New: func() interface{} {
//			return &MonitorTemplate{}
//		},
//		NewSlice:func() interface{} {
//			return &[]MonitorTemplate{}
//		},
//	}
//}
//
//func monitorItem() *db.ModelDescriptor {
//	return &db.ModelDescriptor{
//		Type: &MonitorItem{},
//		New: func() interface{} {
//			return &MonitorItem{}
//		},
//		NewSlice:func() interface{} {
//			return &[]MonitorItem{}
//		},
//	}
//}

//DiscoveryTemplate controls how to discovery resource on the host of the agent
type DiscoveryTemplate struct {
	gorm.Model
	ResourceType string          `json:"resourceType"`
	Name         string          `json:"name"`
	Host         bool            `json:"host"`
	Period       int64           `json:"period"`
	Items        []DiscoveryItem `json:"items"`
}

//DiscoveryItem controls how to discovery all configs of the resource
type DiscoveryItem struct {
	gorm.Model
	Metric  string `json:"metric"`
	Service string `json:"service`
}

//MonitorTemplate controls how to monitor metrics
type MonitorTemplate struct {
	gorm.Model
	ResourceType string        `json:"resourceType"`
	Name         string        `json:"name"`
	Items        []MonitorItem `json:"items"`
}

//MonitorItem controls how to collect a metric and process alert
type MonitorItem struct {
	gorm.Model
	Metric string `json:"metric"`
	Period int64  `json:"period"`
}

func LoadMonitorTemplates(path string) ([]*MonitorTemplate, error) {
	templates := []*MonitorTemplate{}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if file.Ext(path) == ".json" {
			template := &MonitorTemplate{}
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

func LoadDiscoveryTemplates(path string) ([]*DiscoveryTemplate, error) {
	templates := []*DiscoveryTemplate{}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if file.Ext(path) == ".json" {
			template := &DiscoveryTemplate{}
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
