package main

import (
	"github.com/pressly/chi"
	"net/http"
	"strings"
	chttp "github.com/ChaosXu/nerv/lib/net/http"
)

func FileServer(mx *chi.Mux, path string, root http.FileSystem) {
	if strings.ContainsAny(path, ":*") {
		panic("chi: FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, chttp.FileServer(root))

	if path != "/" && path[len(path) - 1] != '/' {
		mx.Get(path, http.RedirectHandler(path + "/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	mx.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
