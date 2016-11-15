package env

import (
	"os"
	"flag"
	"log"
)

var (
	config *Properties
	Debug *bool
	Setup *bool
)

func SetConfig(c *Properties) {
	config = c
}

func Config() *Properties {
	if config == nil {
		panic("config is nil")
	} else {
		return config
	}
}

func Init() {
	configPath := flag.String("c", "../config/config.json", "configuration file")
	Debug = flag.Bool("d", false, "show debug info")
	Setup = flag.Bool("s", false, "setup schema")

	flag.Parse()

	if c, err := LoadConfig(*configPath); err != nil {
		log.Println(err.Error())
		os.Exit(1)
	} else {
		config = c
	}

	if *Debug {
		log.Printf("%+v\n", config)
	}
}
