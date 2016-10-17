package env

import (
	"fmt"
	"os"
	"flag"
)

var (
	Config *Properties
	Debug *bool
)

func Init() {
	configPath := flag.String("c", "../config/config.json", "configuration file")
	Debug = flag.Bool("d", false, "show debug info")

	flag.Parse()

	if config, err := LoadConfig(*configPath); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		Config = config
	}

	if *Debug {
		fmt.Printf("%+v\n", Config)
	}
}
