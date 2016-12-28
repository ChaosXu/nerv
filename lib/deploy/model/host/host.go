package host

import (
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/ChaosXu/nerv/lib/deploy/model/topology"
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

// Host for service install
type Host struct {
	gorm.Model
	Name       string            `json:"name";gorm:"unique;not null"`
	IP         string            `json:"ip";gorm:"unique;not null"`
	Template   string            `json:"template";gorm:"unique;not null"` //The template of the host
	TopologyID uint               `json:"topologyID";gorm"unique"`         //The topology of components in the host
}

// Install an agent in the host
func (p *Host) Install() error {
	template, err := topology.GetTemplate(p.Template)
	if err != nil {
		return err
	}

	topology, err := template.NewTopology(p.Name)
	if err != nil {
		return err
	}

	if err = db.DB.Create(topology).Error; err != nil {
		return err
	}

	if err = topology.Install(); err != nil {
		return err
	} else {
		p.TopologyID = topology.ID
		return db.DB.Save(p).Error
	}
}

// Uninstall an agent from host
func (p *Host) Uninstall() error {
	return nil
}

// Configure an agent in the host
func (p *Host) Configure() error {
	return fmt.Errorf("TBD")
}

// Start an agent in the host
func (p *Host) Start() error {
	return fmt.Errorf("TBD")
}

// Stop an agent in the host
func (p *Host) Stop() error {
	return fmt.Errorf("TBD")
}
