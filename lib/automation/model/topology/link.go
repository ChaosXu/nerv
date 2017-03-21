package topology

import (
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/jinzhu/gorm"
)

//Link two nodes
type Link struct {
	gorm.Model
	NodeID int    `gorm:"index"` //Foreign key of the node
	Type   string //link type
	Source string //source node name
	Target string //target node name
}

func init() {
	db.Models["Link"] = linkDesc()
}

func linkDesc() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &Link{},
		New: func() interface{} {
			return &Link{}
		},
		NewSlice: func() interface{} {
			return &[]Link{}
		},
	}
}
