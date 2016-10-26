package main

import (
	"log"
	"github.com/ChaosXu/nerv/lib/env"
	"github.com/ChaosXu/nerv/cmd/agent/deploy"
	"github.com/ChaosXu/nerv/cmd/agent/monitor"
)

var (
	Version = "main.min.build"
)

func main() {
	log.Println("Version:" + Version)
	env.Init()

	agent, err := deploy.NewAgent()
	if err != nil {
		log.Fatalln(err.Error())
	}

	monitor := monitor.NewMonitor()
	if err := monitor.Start(); err != nil {
		log.Fatalln(err.Error())
	}

	if err := agent.Start(); err != nil {
		log.Fatalln(err.Error())
	}
}

