package main

import (
	"fmt"
	"github.com/ChaosXu/nerv/lib/env"
	"github.com/ChaosXu/nerv/lib/rpc"
	"log"
	"os"
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
