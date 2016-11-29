package main

import (
	"net/http"
	"log"

	chim "github.com/pressly/chi/middleware"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pressly/chi"
	"github.com/ChaosXu/nerv/lib/rest/middleware"
	"github.com/ChaosXu/nerv/lib/rest"
	"github.com/ChaosXu/nerv/lib/db"
	user "github.com/ChaosXu/nerv/lib/user/model"
	"fmt"
	"github.com/ChaosXu/nerv/lib/env"
	"os"
)

var (
	Version = "main.min.build"
)

func main() {
	fmt.Println("Workspace:" + os.Args[0])
	fmt.Println("Version:" + Version)
	env.Init()

	if *env.Setup {
		log.Println("setup...")
		setup()
		log.Println("setup success")
	} else {
		log.Println("run")
		initDB()
		defer db.DB.Close()

		r := initRouter()
		port := env.Config().GetMapString("http", "port", "3333")
		log.Fatal(http.ListenAndServe(":" + port, r))
	}
}

func setup() {
	initDB();
	createAdmin();
	db.DB.Close();
}

func createAdmin() {
	admin := &user.Account{
		Name:"admin",
		Nick:"admin",
		Mail: "admin@nerv.com",
		Phone: 11111111111,
		Password:"admin",
	}
	if err := db.DB.Create(admin).Error; err != nil {
		log.Fatal(err.Error());
	}
}

func initRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(chim.Logger)
	r.Use(middleware.ParamsParser)

	rest.RouteObj(r)
	return r
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
	db.DB.LogMode(true)
	for _, v := range db.Models {
		db.DB.AutoMigrate(v.Type)
	}
}




