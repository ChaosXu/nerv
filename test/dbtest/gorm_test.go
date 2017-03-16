package dbtest

import (
	"testing"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/ChaosXu/nerv/lib/db"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var models []interface{} = []interface{}{
	&Host{},
	&Address{},
}

type Host struct {
	gorm.Model
	Name        string
	Description string
	Address     *[]Address
}

type Address struct {
	gorm.Model
	IP string
}

func TestDB(t *testing.T) {
	if err := initDB(); err != nil {
		t.Fatal(err.Error())
	}
	defer db.DB.Close()

	host := Host{
		Name:"host1",
		Description:"host1-desc",
		Address:&[]Address{
			{IP:"1.1.1.1"},
			{IP:"1.1.1.2"},
		},
	}
	if err := db.DB.Create(&host).Error; err != nil {
		t.Error(err)
	}
	host2 := &Host{}
	if err := db.DB.First(host2, host.ID).Error; err != nil {
		t.Error(err)
	}

	host3 := Host{
		Model: gorm.Model{ID:host2.ID},
		Description:"update",
	}

	if err := db.DB.Save(host3).Error; err != nil {
		t.Error(err)
	}

	//host4 := &Host{}
	//if err := db.DB.First(host4, host.ID).Error; err != nil {
	//	t.Error(err)
	//}
	//assert.Equal(t, "update", host4.Description)
	//assert.Equal(t, "host", host4.Name)
}

func initDB() error {
	url := fmt.Sprintf(
		"%s:%s@%s",
		"root",
		"root",
		"/nerv?charset=utf8&parseTime=True&loc=Local",
	)
	gdb, err := gorm.Open("mysql", url)
	if err != nil {
		return err
	}
	db.DB = gdb
	db.DB.LogMode(false)
	for _, v := range models {
		db.DB.AutoMigrate(v)
	}
	return nil
}
