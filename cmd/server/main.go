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

	initDB()
	defer db.DB.Close()

	r := initRouter()
	log.Fatal(http.ListenAndServe(":3333", r))
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
		env.Config().GetString("user", "root"),
		env.Config().GetString("password", "root"),
		env.Config().GetString("url"),
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




