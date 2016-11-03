package store

import (
	"os"
	"log"
	"github.com/ChaosXu/nerv/lib/env"
	"github.com/ChaosXu/nerv/lib/monitor/model"
	"github.com/ChaosXu/nerv/lib/monitor/shipper/elasticsearch"
)

var (
	Version = "main.min.build"
)

func main() {
	log.Println("Version:" + Version)
	env.Init()

	if env.Setup {
		log.Println("Setup...")
		setup()
	}
}

func setup() {
	path := env.Config().GetMapString("metrics", "path", "resources/path")
	if metrics, err := model.LoadMetrics(path); err != nil {
		log.Panicln(err.Error())
		os.Exit(1)
	}else{
		elasticsearch.CreateSchemas(env.Config().GetMapString("shipper","server","localhost:9200"),metrics)
	}
}
