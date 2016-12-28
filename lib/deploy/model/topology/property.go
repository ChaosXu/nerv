package topology

import (
	"github.com/jinzhu/gorm"
	"github.com/ChaosXu/nerv/lib/db"
)

//Link two nodes
type Property struct {
	gorm.Model
	NodeID int        `gorm:"index"` //Foreign key of the node
	Key    string
	Value  string
}

func init() {
	db.Models["Property"] = propertyDesc()
}

func propertyDesc() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &Property{},
		New: func() interface{} {
			return &Property{}
		},
		NewSlice:func() interface{} {
			return &[]Property{}
		},
	}
}
