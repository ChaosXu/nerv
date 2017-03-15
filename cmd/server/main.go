package main

import (

	"log"
	"fmt"
	"os"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/ChaosXu/nerv/lib/env"
	libsvc "github.com/ChaosXu/nerv/lib/service"
	_ "github.com/ChaosXu/nerv/cmd/server/service"
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

		initServices()
	}
}

func initServices() {
	for _, factory := range libsvc.Registry.Services {
		if err := factory.Init(); err != nil {
			log.Fatalln(err.Error())
		}

		svc := factory.Get()
		if svc != nil {
			initializer, ok := svc.(libsvc.Initializer)
			if ok {
				if err := initializer.Init(); err != nil {
					log.Fatalln(err.Error())
				}
			}
		}
	}
	select {}
}





