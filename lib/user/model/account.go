package model

import (
	"github.com/jinzhu/gorm"
	"github.com/ChaosXu/nerv/lib/db"
)

// Account is a principal of the user
type Account struct {
	gorm.Model
	Name string	`gorm:"not null;unique"`
	Nick string `gorm:"not null;unique"`
	Mail string	`gorm:"not null;unique"`
	Phone int64 `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}

func init() {
	db.Models["Account"] = accountDesc()
}

func accountDesc() *db.ModelDescriptor {
	return &db.ModelDescriptor{
		Type: &Account{},
		New: func() interface{} {
			return &Account{}
		},
		NewSlice:func() interface{} {
			return &[]Account{}
		},
	}
}
