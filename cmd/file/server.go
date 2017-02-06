package main

import (
	"github.com/pressly/chi"
	"net/http"
	"strings"
	httpf "github.com/ChaosXu/nerv/lib/net/http/file"
)

func FileServer(mx *chi.Mux, path string, root httpf.FileSystem) {
	if strings.ContainsAny(path, ":*") {
		panic("chi: FileServer does not permit URL parameters.")
	}

	fs := httpf.FileServer(root)
	prefix := path
	path += "*"
	mx.Get(path, exec(prefix, func(w http.ResponseWriter, r *http.Request) {
		fs.Get(w, r)
	}))

	mx.Post(path, exec(prefix, func(w http.ResponseWriter, r *http.Request) {
		fs.Post(w, r)
	}))
	mx.Put(path, exec(prefix, func(w http.ResponseWriter, r *http.Request) {
		fs.Put(w, r)
	}))
	mx.Delete(path, exec(prefix, func(w http.ResponseWriter, r *http.Request) {
		fs.Delete(w, r)
	}))
}

func exec(prefix string, fn http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if p := strings.TrimPrefix(r.URL.Path, prefix); len(p) < len(r.URL.Path) {
			r.URL.Path = p
			fn(w, r)
		} else {
			http.NotFound(w, r)
		}
	})
}
