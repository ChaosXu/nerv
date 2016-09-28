package main

import (
	"net/http"

	"github.com/pressly/chi"
	//"github.com/pressly/chi/render"
	//"github.com/jinzhu/gorm"
	//"github.com/open-falcon/common/db"
	//"reflect"
	"github.com/pressly/chi/render"
	"fmt"
	"github.com/chaosxu/nerv/lib/model"
	"github.com/chaosxu/nerv/lib/middleware"
	"strings"
	"github.com/jinzhu/gorm"
)

type User struct {
	Name string
}

func routeObj(r *chi.Mux) {
	r.Route("/objs/:class", func(r chi.Router) {
		r.Get("/", list)
		r.Post("/", create)
		r.Put("/", update)
		r.Route("/:id", func(r chi.Router) {
			r.Get("/", get)
			r.Delete("/", remove)
		})
	})
}

//func getRepository(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
//		class := chi.URLParam(req, "class")
//
//	})
//}

func handlePanic(w http.ResponseWriter, req *http.Request) {
	if r := recover(); r != nil {
		fmt.Println(r)
		render.Status(req, 500)
		render.JSON(w, req, r)
	}
}

func list(w http.ResponseWriter, req *http.Request) {
	//class := chi.URLParam(req, "class")

	//db,err :=gorm.Open("","")
	//if err!=nil {
	//	panic("failed to connect")
	//}
	//defer db.Close()
	//
	//db.Create()
}

// get one obj. query params: assocations=a,b...
func get(w http.ResponseWriter, req *http.Request) {
	defer handlePanic(w, req)

	class := middleware.CurrentParams(req).PathParam("class")
	id := middleware.CurrentParams(req).PathParam("id")
	ass := middleware.CurrentParams(req).QueryParam("associations")

	md := model.Models[class]
	if md == nil {
		render.Status(req, 400)
		render.JSON(w, req, fmt.Sprintf("class %s isn't exists", class))
		return
	}

	data := md.New()
	var db *gorm.DB
	if ass == "" {
		db = DB
	} else {
		for _, as := range strings.Split(ass, ",") {
			db = DB.Preload(as)
		}
	}
	if db.First(data, id).RecordNotFound() {
		render.Status(req, 200)
		render.JSON(w, req, nil)
		return
	}

	render.Status(req, 200)
	render.JSON(w, req, data)
}

func create(w http.ResponseWriter, req *http.Request) {
	defer handlePanic(w, req)

	class := middleware.CurrentParams(req).PathParam("class")
	md := model.Models[class]
	if md == nil {
		render.Status(req, 400)
		render.JSON(w, req, fmt.Sprintf("class %s isn't exists", class))
		return
	}

	data := md.New()
	if err := render.Bind(req.Body, data); err != nil {
		render.Status(req, 400)
		render.JSON(w, req, err.Error())
		return
	}

	if err := DB.Create(data).Error; err != nil {
		render.Status(req, 500)
		render.JSON(w, req, err.Error())
		return
	}

	render.Status(req, 200)
	render.JSON(w, req, data)
}

func remove(w http.ResponseWriter, req *http.Request) {
	defer handlePanic(w, req)

	class := middleware.CurrentParams(req).PathParam("class")
	id := middleware.CurrentParams(req).PathParam("id")

	md := model.Models[class]
	if md == nil {
		render.Status(req, 400)
		render.JSON(w, req, fmt.Sprintf("class %s isn't exists", class))
		return
	}

	data := md.New()
	if err := DB.First(data, id).Error; err != nil {
		render.Status(req, 400)
		render.JSON(w, req, err.Error())
		return
	}

	if err := DB.Unscoped().Delete(data).Error; err != nil {
		render.Status(req, 500)
		render.JSON(w, req, err.Error())
		return
	}

	render.Status(req, 200)
}

func update(w http.ResponseWriter, req *http.Request) {
	defer handlePanic(w, req)

	class := middleware.CurrentParams(req).PathParam("class")
	md := model.Models[class]
	if md == nil {
		render.Status(req, 400)
		render.JSON(w, req, fmt.Sprintf("class %s isn't exists", class))
		return
	}

	data := md.New()
	if err := render.Bind(req.Body, data); err != nil {
		render.Status(req, 400)
		render.JSON(w, req, err.Error())
		return
	}

	if err := DB.Save(data).Error; err != nil {
		render.Status(req, 500)
		render.JSON(w, req, err.Error())
		return
	}

	render.Status(req, 200)
	render.JSON(w, req, data)
}


