package env

import (
	"os"
	"flag"
	"log"
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
		log.Println(err.Error())
		os.Exit(1)
	} else {
		Config = config
	}

	if *Debug {
		log.Printf("%+v\n", Config)
	}
}
