package service

import (
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/ChaosXu/nerv/lib/db"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DBService struct {

}

func (p *DBService) Init() error {
	if _, err := os.Stat("../data"); err != nil {
		if err := os.MkdirAll("../data", os.ModeDir | os.ModePerm); err != nil {
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



