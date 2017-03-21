package db

import (
	"fmt"
	"github.com/ChaosXu/nerv/lib/env"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

type DBService struct{}

func (p *DBService) GetDB() *gorm.DB {
	return DB
}

func (p *DBService) Init() error {
	url := fmt.Sprintf(
		"%s:%s@%s",
		env.Config().GetMapString("db", "user", "root"),
		env.Config().GetMapString("db", "password", "root"),
		env.Config().GetMapString("db", "url"),
	)
	gdb, err := gorm.Open("mysql", url)
	if err != nil {
		return err
	}
	DB = gdb
	DB.LogMode(false)
	for _, v := range Models {
		DB.AutoMigrate(v.Type)
	}

	return nil
}

func (p *DBService) Dispose() error {
	return DB.Close()
}
