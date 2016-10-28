package model

import (
	"github.com/jinzhu/gorm"
	"github.com/ChaosXu/nerv/lib/db"
)

func init() {
	db.Models["Metric"] = metric()
	db.Models["MetricField"] = metricField()
	db.Models["Probe"] = probe()
}

func metric() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &Metric{},
		New: func() interface{} {
			return &Metric{}
		},
		NewSlice:func() interface{} {
			return &[]Metric{}
		},
	}
}

func metricField() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &MetricField{},
		New: func() interface{} {
			return &MetricField{}
		},
		NewSlice:func() interface{} {
			return &[]MetricField{}
		},
	}
}

func probe() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &Probe{},
		New: func() interface{} {
			return &Probe{}
		},
		NewSlice:func() interface{} {
			return &[]Probe{}
		},
	}
}

//MetricType define the type of metric
type MetricType string

const (
	MetricTypeStruct MetricType = "struct"
	MetricTypeTable MetricType = "table"
)

//MetricDataType define the type of metric field's data
type MetricDataType string

const (
	MetricDataTypeString MetricDataType = "string"
	MetricDataTypeDouble MetricDataType = "double"
	MetricDataTypeLong MetricDataType = "long"
)

//MetricSampleType define the type fo metric's sample
type MetricSampleType string

const (
	MetricSampleTypeGauge MetricSampleType = "gauge"        //raw data
	MetricSampleTypeCounter MetricSampleType = "counter"    //V2>=V1 ? (V2-V1)/(T2-T1) : (MAX-V1+V2)/(T2-T1)
)

//ProbeType define the sampling method
type ProbeType string

const (
	ProbeTypeShell ProbeType = "shell"
)

//Metric define the KPI of resource
type Metric struct {
	gorm.Model
	ResourceType string                `json:"resourceType"`
	Type         MetricType            `json:"type"`
	Name         string                `json:"name"`
	IsService    bool                  `json:"isService`
	Fields       []MetricField         `json:"fields"`
}

func (p *Metric) Key() string {
	for _, m := range p.Fields {
		if m.Key {
			return m.Name
		}
	}
	return ""
}

//MetricField define the filed of KPI
type MetricField struct {
	gorm.Model
	Name       string           `json:"name"`
	Key        bool             `json:"key"`
	DataType   MetricDataType   `json:"dataType"`
	SampleType MetricSampleType `json:"sampleType"`
	Probe      Probe            `json:"probe"`
}

//Probe define the sampling info
type Probe struct {
	Type     ProbeType    `json:"type"`
	Provider string       `json:"provider"`
}

