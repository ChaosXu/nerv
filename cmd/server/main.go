package main

import (
	"fmt"
	"github.com/ChaosXu/nerv/cmd/server/service"
	"github.com/ChaosXu/nerv/lib/automation/manager"
	"github.com/ChaosXu/nerv/lib/brick"
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/ChaosXu/nerv/lib/env"
	"github.com/ChaosXu/nerv/lib/net/http/rest"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
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

		container := brick.NewContainer()
		container.Add(&db.DBService{}, "DBService", nil)
		container.Add(&service.HttpService{}, "HTTP", nil)
		container.Add(&rest.RestController{}, "RestController", nil)
		container.Add(&manager.Deployer{}, "Topology", &service.TopologyServiceFactory{})
		container.Build()
		defer container.Dispose()
		select {}
	}
}
