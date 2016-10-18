package main

import (
	"net/http"
	"log"

	"github.com/pressly/chi"
	"github.com/ChaosXu/nerv/lib/env"
	"os"
)

var (
	Version = "main.min.build"
)

func main() {
	log.Println(os.Getwd())
	log.Println("Version:" + Version)
	env.Init()

	port := env.Config.GetString("http_port")
	if port == "" {
		log.Fatalln("http_port isn't setted")
	}

	r := chi.NewRouter()

	for url, file := range env.Config.GetMap("files") {
		log.Printf("file router: %s -> %s", url, file)
		r.FileServer(url, http.Dir(file.(string)))
	}

	log.Fatalln(http.ListenAndServe(":" + port, r))
}