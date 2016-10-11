package main

import (
	"flag"
	"net/http"
	"log"

	chim "github.com/pressly/chi/middleware"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pressly/chi"
	"github.com/chaosxu/nerv/lib/middleware"
	"github.com/chaosxu/nerv/lib/rest"
	"github.com/chaosxu/nerv/lib/db"
)

//var routes = flag.Bool("routes", false, "Generate router documentation")


func main() {
	flag.Parse()

	initDB()
	defer db.DB.Close()

	r := chi.NewRouter()
	r.Use(chim.Logger)
	r.Use(middleware.ParamsParser)

	rest.RouteObj(r)

	log.Fatal(http.ListenAndServe(":3333", r))
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




