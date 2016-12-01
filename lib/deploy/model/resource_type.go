package model

import (
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/jinzhu/gorm"
)

func init() {
	db.Models["ResourceType"] = resourceTypeDesc()
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

type ResourceType struct {
	gorm.Model
	Name    string           `json:"name";gorm:"unique;not null"`
	Version int64            `json:"version";gorm:"not null"`
}
