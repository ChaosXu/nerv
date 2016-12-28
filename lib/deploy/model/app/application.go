package app

import (
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/jinzhu/gorm"
	"fmt"
)

func init() {
	db.Models["Application"] = appDesc()
}

func appDesc() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &Application{},
		New: func() interface{} {
			return &Application{}
		},
		NewSlice:func() interface{} {
			return &[]Application{}
		},
	}
}

// Host for service install
type Application struct {
	gorm.Model
	Name       string            `json:"name";gorm:"unique;not null"`
	IP         string            `json:"ip";gorm:"unique;not null"`
	Template   string            `json:"template";gorm:"unique;not null"` //The template of the host
	TopologyID int               `json:"topologyID";gorm"unique"`	//The topology of components in the host
}

// Install an agent in the host
func (p *Application) Install() error {

	return nil
}

// Uninstall an agent from host
func (p *Application) Uninstall() error {
	return nil
}

// Configure an agent in the host
func (p *Application) Configure() error {
	return fmt.Errorf("TBD")
}

// Start an agent in the host
func (p *Application) Start() error {
	return fmt.Errorf("TBD")
}

// Stop an agent in the host
func (p *Application) Stop() error {
	return fmt.Errorf("TBD")
}
