package model

import (
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/jinzhu/gorm"
)

func init() {
	db.Models["ResourceType"] = resourceTypeDesc()
	db.Models["Operation"] = operationDesc()
}

func resourceTypeDesc() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &ResourceType{},
		New: func() interface{} {
			return &ResourceType{}
		},
		NewSlice:func() interface{} {
			return &[]ResourceType{}
		},
	}
}

func operationDesc() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &Operation{},
		New: func() interface{} {
			return &Operation{}
		},
		NewSlice:func() interface{} {
			return &[]Operation{}
		},
	}
}

type ResourceType struct {
	gorm.Model
	Name       string           `json:"name";gorm:"unique;not null"`
	Version    int64            `json:"version";gorm:"not null"`
	Operations []Operation   `json:"operations"`
}

type Operation struct {
	gorm.Model
	ResourceTypeID int           `json:"resourceTypeID";gorm:"index"` //Foreign key
	Name              string           `json:"name";gorm:"unique;not null"`
	Type              string            `json:"type";gorm:"not null"`
	Implementor       string            `json:"implementor";gorm:"not null"`
}
