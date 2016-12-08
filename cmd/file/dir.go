package main

import (
	"net/http"
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
)

func InitDir(r *chi.Mux) {

	r.Route("/api/files", func(r chi.Router) {
		r.Get("/", list)
		r.Get("/*", list)
		r.Post("/*", create)
		r.Put("/*", save)
		r.Delete("/", remove)
	})
}

func list(w http.ResponseWriter, req *http.Request) {

	render.Status(req, 200)
	render.JSON(w, req, map[string]interface{}{"name":"file1", "type":"dir"})
}

func create(w http.ResponseWriter, req *http.Request) {
	render.Status(req, 200)
	render.JSON(w, req, map[string]interface{}{"name":"create", "type":"dir"})
}

func save(w http.ResponseWriter, req *http.Request) {
	render.Status(req, 200)
	render.JSON(w, req, map[string]interface{}{"name":"file1", "type":"dir"})
}

func remove(w http.ResponseWriter, req *http.Request) {
	render.Status(req, 200)
	render.JSON(w, req, map[string]interface{}{"name":"file1", "type":"dir"})
}
