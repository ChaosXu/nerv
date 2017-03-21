package main

import (
	"fmt"
	"github.com/ChaosXu/nerv/cmd/agent/cmd"
	_ "github.com/ChaosXu/nerv/cmd/agent/model"
	"os"
)

var (
	Version = "main.min.build"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
