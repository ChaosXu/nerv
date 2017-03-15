package main

import (
	"fmt"
	"os"
	"github.com/ChaosXu/nerv/cmd/agent/cmd"
	_ "github.com/ChaosXu/nerv/cmd/agent/model"
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

