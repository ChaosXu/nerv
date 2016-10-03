package main

import (
	"flag"
	"net/http"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/pressly/chi"
	"github.com/chaosxu/nerv/lib/middleware"
	chim "github.com/pressly/chi/middleware"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/chaosxu/nerv/lib/db/rest"
	"github.com/chaosxu/nerv/lib/db"
	"github.com/chaosxu/nerv/lib/model"
)

//var routes = flag.Bool("routes", false, "Generate router documentation")


func main() {
	flag.Parse()

	gdb, err := gorm.Open("mysql", "root:root@/nerv?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.DB = gdb
	db.DB.LogMode(true)
	for _, v := range model.Models {
		db.DB.AutoMigrate(v.Type)
	}
	defer db.DB.Close()

	r := chi.NewRouter()
	r.Use(chim.Logger)
	r.Use(middleware.ParamsParser)

	rest.RouteObj(r)

	log.Fatal(http.ListenAndServe(":3333", r))
}




