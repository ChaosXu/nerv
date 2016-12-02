package model

import (
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/jinzhu/gorm"
)

func init() {
	db.Models["Script"] = scriptDesc()
}

func scriptDesc() *db.ModelDescriptor {
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

type Script struct {
	gorm.Model
	Name       string           `json:"name";gorm:"unique;not null"`
	Path       string            `json:"path";gorm:"not null"`
	Type       string            `json:"type";gorm:"not null"`
	File       string            `json:"file";gorm:"not null"`
}

