package credential

import (
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/jinzhu/gorm"
)

func init() {
	db.Models["Credential"] = credentialDesc()
}

func credentialDesc() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &Credential{},
		New: func() interface{} {
			return &Credential{}
		},
		NewSlice: func() interface{} {
			return &[]Credential{}
		},
	}
}

//Credential is used to login a host
type Credential struct {
	gorm.Model
	Type     string `json:"type"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}
