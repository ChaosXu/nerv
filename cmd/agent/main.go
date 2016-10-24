package main

import (
	"fmt"
	"log"
	"os"
	"github.com/ChaosXu/nerv/lib/env"
	"github.com/ChaosXu/nerv/lib/rpc"
	_ "github.com/ChaosXu/nerv/cmd/agent/shell"
)

var (
	Version = "main.min.build"
)

func main() {
	fmt.Println("Version:" + Version)
	env.Init()
	if err := rpc.Start(); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
