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
)

//var routes = flag.Bool("routes", false, "Generate router documentation")

var (
	DB *gorm.DB
)

func main() {
	flag.Parse()

	db, err := gorm.Open("mysql", "root:root@/nerv?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DB = db
	AutoMigrate()
	defer DB.Close()

	r := chi.NewRouter()
	r.Use(chim.Logger)
	r.Use(middleware.ParamsParser)

	routeObj(r)

	log.Fatal(http.ListenAndServe(":3333", r))
}



