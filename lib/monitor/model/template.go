package model

import (
	"github.com/jinzhu/gorm"
	"github.com/ChaosXu/nerv/lib/db"
)

func init() {
	db.Models["DiscoveryTemplate"] = discoveryTemplate()
	db.Models["DiscoveryItem"] = discoveryItem()
	db.Models["MonitorTemplate"] = monitorTemplate()
	db.Models["MonitorItem"] = monitorItem()
}

func discoveryTemplate() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &DiscoveryTemplate{},
		New: func() interface{} {
			return &DiscoveryTemplate{}
		},
		NewSlice:func() interface{} {
			return &[]DiscoveryTemplate{}
		},
	}
}

func discoveryItem() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &DiscoveryItem{},
		New: func() interface{} {
			return &DiscoveryItem{}
		},
		NewSlice:func() interface{} {
			return &[]DiscoveryItem{}
		},
	}
}

func monitorTemplate() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &MonitorTemplate{},
		New: func() interface{} {
			return &MonitorTemplate{}
		},
		NewSlice:func() interface{} {
			return &[]MonitorTemplate{}
		},
	}
}

func monitorItem() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &MonitorItem{},
		New: func() interface{} {
			return &MonitorItem{}
		},
		NewSlice:func() interface{} {
			return &[]MonitorItem{}
		},
	}
}

//DiscoveryTemplate controls how to discovery resource on the host of the agent
type DiscoveryTemplate struct {
	gorm.Model
	ResourceType string                 `json:"resourceType"`
	Name         string                 `json:"name"`
	Host         bool                   `json:"host"`
	Items        []DiscoveryItem        `json:"items"`
}

//DiscoveryItem controls how to discovery all configs of the resource
type DiscoveryItem struct {
	gorm.Model
	Metric  string      `json:"metric"`
	Service string      `json:"service`
}

//MonitorTemplate controls how to monitor metrics
type MonitorTemplate struct {
	gorm.Model
	ResourceType string                   `json:"resourceType"`
	Name         string                   `json:"name"`
	Items        []MonitorItem            `json:"items"`
}

//MonitorItem controls how to collect a metric and process alert
type MonitorItem struct {
	gorm.Model
	Metric string   `json:"metric"`
	Period int64    `json:"period"`
}
