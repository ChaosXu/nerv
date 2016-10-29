package main

import (
	"log"
	"github.com/ChaosXu/nerv/lib/env"
_	"github.com/ChaosXu/nerv/cmd/agent/deploy"
	"github.com/ChaosXu/nerv/cmd/agent/monitor"
	"github.com/ChaosXu/nerv/cmd/agent/deploy"
)

var (
	Version = "main.min.build"
)

func main() {
	log.Println("Version:" + Version)
	env.Init()

	monitor := monitor.NewMonitor()
	monitor.Start()

	agent, err := deploy.NewAgent()
	if err != nil {
		log.Fatalln(err.Error())
	}
	if err := agent.Start(); err != nil {
		log.Fatalln(err.Error())
	}
}

