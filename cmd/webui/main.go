package main

import (
	"github.com/ChaosXu/nerv/lib/env"
	"github.com/ChaosXu/nerv/lib/net/http/proxy"
	"github.com/pressly/chi"
	"log"
	"net/http"
	"os"
)

var (
	Version = "main.min.build"
)

func main() {
	log.Println(os.Getwd())
	log.Println("Version:" + Version)
	env.Init()

	port := env.Config().GetMapString("http", "port")
	if port == "" {
		log.Fatalln("http_port isn't setted")
	}

	r := chi.NewRouter()
	for source, target := range env.Config().GetMap("proxies") {
		remoteUrl, ok := target.(string)
		if !ok {
			log.Fatalf("get remote url failed. source url:%s\n", source)
		}
		serverProxy, err := proxy.NewProxy(remoteUrl)
		if err != nil {
			log.Fatalln(err.Error())
		}
		r.Route(source, func(r chi.Router) {
			r.Post("/*", func(w http.ResponseWriter, r *http.Request) {
				serverProxy.Handle(w, r)
			})
			r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
				serverProxy.Handle(w, r)
			})
			r.Delete("/*", func(w http.ResponseWriter, r *http.Request) {
				serverProxy.Handle(w, r)
			})
			r.Put("/*", func(w http.ResponseWriter, r *http.Request) {
				serverProxy.Handle(w, r)
			})
		})
	}

	for url, file := range env.Config().GetMap("files") {
		log.Printf("file router: %s -> %s", url, file)
		r.FileServer(url, http.Dir(file.(string)))
	}

	log.Fatalln(http.ListenAndServe(":"+port, r))
}
