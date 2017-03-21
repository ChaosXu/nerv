package main

import (
	"fmt"
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/ChaosXu/nerv/lib/env"
	user "github.com/ChaosXu/nerv/lib/user/model"
	"github.com/jinzhu/gorm"
	"log"
)

func setup() {
	initDB()
	createAdmin()
	db.DB.Close()
}

func createAdmin() {
	admin := &user.Account{
		Name:     "admin",
		Nick:     "admin",
		Mail:     "admin@nerv.com",
		Phone:    11111111111,
		Password: "admin",
	}
	var count int64
	if err := db.DB.Model(&user.Account{}).Where("Name=?", "admin").Count(&count).Error; err != nil {
		log.Fatal(err.Error())
	}
	if count == 0 {
		if err := db.DB.Create(admin).Error; err != nil {
			log.Fatal(err.Error())
		}
	}
}

func initDB() {
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
	db.DB.LogMode(false)
	for _, v := range db.Models {
		db.DB.AutoMigrate(v.Type)
	}
}
