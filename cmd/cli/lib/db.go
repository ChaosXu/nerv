package lib

import (
	"fmt"
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/ChaosXu/nerv/lib/env"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB() *gorm.DB {
	url := fmt.Sprintf(
		"%s:%s@%s",
		env.Config().GetMapString("db", "user", "root"),
		env.Config().GetMapString("db", "password", "root"),
		env.Config().GetMapString("db", "url"),
	)
	gdb, err := gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	db.DB = gdb
	//db.DB.LogMode(true)
	for _, v := range db.Models {
		db.DB.AutoMigrate(v.Type)
	}

	return db.DB
}
