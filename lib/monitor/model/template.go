package model

import (
	"github.com/jinzhu/gorm"
	"github.com/ChaosXu/nerv/lib/db"
)

func init() {
	db.Models["MonitorTemplate"] = monitorTemplate()
	db.Models["MonitorItem"] = monitorItem()
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

//MonitorTemplate controls how to monitor metrics
type MonitorTemplate struct {
	gorm.Model
	ResourceType string
	Items        []MonitorItem
}

//MonitorItem controls how to collect a metric and process alert
type MonitorItem struct {
	gorm.Model
	Metric string
	Period int64
}
