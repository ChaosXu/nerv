package model

import (
	"github.com/jinzhu/gorm"
	"github.com/ChaosXu/nerv/lib/db"
	"k8s.io/kubernetes/pkg/registry/core/service/portallocator"
	"github.com/cloudinsight/cloudinsight-agent/common/log"
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
	Name     string    `json:"name" gorm:"not null;unique"`
	Path     string    `json:"path" gorm:"not null;unique"`
	Version  string    `json:"version"`
	Services []Service `json:"services"`
	Log      *Log
}

// Service provide by an app
type Service struct {
	gorm.Model
	Name string    `json:"name" gorm:"not null;unique"` //service name
	Type string    `json:"type" gorm:"not null"`        //service type:e.g. http | https
	Port int32     `json:"port" gorm:"not null"`        //service port
	Uri  string    `json:"uri" gorm:"not null;unique"`  //service uri
}

// Log config the information of app
type Log struct {

}
