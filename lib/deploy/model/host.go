package model

import (
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/jinzhu/gorm"
)

func init() {
	db.Models["Host"] = hostDesc()
}

func hostDesc() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &Host{},
		New: func() interface{} {
			return &Host{}
		},
		NewSlice:func() interface{} {
			return &[]Host{}
		},
	}
}

type Host struct {
	gorm.Model
	Name    string           `json:"name";gorm:"unique;not null"`
	IP 		string            `json:"ip";gorm:"unique;not null"`
}
