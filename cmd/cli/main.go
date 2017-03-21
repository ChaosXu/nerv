package main

import (
	"fmt"
	"github.com/ChaosXu/nerv/cmd/cli/cmd"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
