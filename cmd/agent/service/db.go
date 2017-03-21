package service

import (
	"fmt"
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

type DBService struct {
}

func (p *DBService) Init() error {
	if _, err := os.Stat("../data"); err != nil {
		if err := os.MkdirAll("../data", os.ModeDir|os.ModePerm); err != nil {
			return fmt.Errorf("create dir ../data failed. %s", err.Error())
		}
	}

	gdb, err := gorm.Open("sqlite3", "../data/agent.db")
	if err != nil {
		return err
	}
	db.DB = gdb
	db.DB.LogMode(false)
	for _, v := range db.Models {
		db.DB.AutoMigrate(v.Type)
	}
	return nil
}

func (p *DBService) Dispose() error {
	return db.DB.Close()
}
