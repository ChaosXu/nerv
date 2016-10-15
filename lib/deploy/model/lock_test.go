package model_test

import (
	"testing"
	"github.com/ChaosXu/nerv/lib/deploy"
	"github.com/stretchr/testify/assert"
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestLock_TryLock(t *testing.T) {

	initDB()

	lock1 := deploy.GetLock("obj", 1)
	lock2 := deploy.GetLock("obj", 1)
	defer lock1.Unlock()
	defer lock2.Unlock()

	ok1 := lock1.TryLock()
	assert.Equal(t, ok1, true)

	ok2 :=lock2.TryLock()
	assert.Equal(t, ok2, false)
}

func initDB() {
	gdb, err := gorm.Open("mysql", "root:root@/nerv?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.DB = gdb
	db.DB.LogMode(true)
	for _, v := range db.Models {
		db.DB.AutoMigrate(v.Type)
	}
}
