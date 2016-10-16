package main

import (
	"flag"
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
	"github.com/ChaosXu/nerv/lib/util"
	"os"
)

var (
	Version = "main.min.build"
	Config *util.Config
)

func main() {
	fmt.Println("Version:" + Version)

	configPath := flag.String("c", "../config/config.json", "configuration file")
	debug := flag.Bool("d", false, "show debug info")

	flag.Parse()

	config, err := util.LoadConfig(*configPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if *debug {
		fmt.Printf("%+v\n", config)
	}
	Config = config

	initDB()
	defer db.DB.Close()

	r := chi.NewRouter()
	r.Use(chim.Logger)
	r.Use(middleware.ParamsParser)

	rest.RouteObj(r)

	log.Fatal(http.ListenAndServe(":3333", r))
}

func initDB() {
	url := fmt.Sprintf(
		"%s:%s@%s",
		Config.GetProperty("user", "root"),
		Config.GetProperty("password", "root"),
		Config.GetProperty("url"),
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




