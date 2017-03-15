package model

import (
	"github.com/jinzhu/gorm"
	"github.com/ChaosXu/nerv/lib/db"
)

func init() {
	db.Models["App"] = appDesc()
}

func appDesc() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &App{},
		New: func() interface{} {
			return &App{}
		},
		NewSlice:func() interface{} {
			return &[]App{}
		},
	}
}


// App managed by agent
type App struct {
	gorm.Model
	Name    string    `json:"name"`
	Path    string    `json:"path"`
	Version string    `json:"version"`
}
