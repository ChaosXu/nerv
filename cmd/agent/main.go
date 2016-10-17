package main

import (
	"fmt"
	"github.com/ChaosXu/nerv/lib/env"
)

var (
	Version = "main.min.build"
)

func main() {
	fmt.Println("Version:" + Version)
	env.Init()

	select {}
}
