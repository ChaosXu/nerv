package db

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)

type DBService struct{}

func (p *DBService) GetDB() *gorm.DB {
	return DB
}

