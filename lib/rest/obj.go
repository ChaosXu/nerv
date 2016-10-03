package rest

import (
	"net/http"
	"fmt"
	"strings"

	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
	"github.com/jinzhu/gorm"
	"github.com/chaosxu/nerv/lib/model"
	"github.com/chaosxu/nerv/lib/middleware"
	"github.com/chaosxu/nerv/lib/db"
)

type User struct {
	Name string
}

func RouteObj(r *chi.Mux) {
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

func handlePanic(w http.ResponseWriter, req *http.Request) {
	if r := recover(); r != nil {
		fmt.Println(r)
		render.Status(req, 500)
		render.JSON(w, req, r)
	}
}

func list(w http.ResponseWriter, req *http.Request) {
	defer handlePanic(w, req)

	class := middleware.CurrentParams(req).PathParam("class")

	where := middleware.CurrentParams(req).QueryParam("where")
	var args []string
	if where != "" {
		values := middleware.CurrentParams(req).QueryParam("values")
		args = strings.Split(values, ",")
		if values == "" {
			render.Status(req, 400)
			render.JSON(w, req, fmt.Sprintf("the values query param must be provided if the where query param is exists"))
		}
	}
	md := model.Models[class]
	if md == nil {
		render.Status(req, 400)
		render.JSON(w, req, fmt.Sprintf("class %s isn't exists", class))
		return
	}

	d := db.DB
	if where != "" {
		d = db.DB.Where(where, args)
	}

	data := md.NewSlice()
	if d.Find(data).RecordNotFound() {
		render.Status(req, 200)
		render.JSON(w, req, data)
		return
	}

	render.Status(req, 200)
	render.JSON(w, req, data)
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
	var d *gorm.DB
	if ass == "" {
		d = db.DB
	} else {
		for _, as := range strings.Split(ass, ",") {
			d = db.DB.Preload(as)
		}
	}
	if d.First(data, id).RecordNotFound() {
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

	if err := db.DB.Create(data).Error; err != nil {
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
	if err := db.DB.First(data, id).Error; err != nil {
		render.Status(req, 400)
		render.JSON(w, req, err.Error())
		return
	}

	if err := db.DB.Unscoped().Delete(data).Error; err != nil {
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

	if err := db.DB.Save(data).Error; err != nil {
		render.Status(req, 500)
		render.JSON(w, req, err.Error())
		return
	}

	render.Status(req, 200)
	render.JSON(w, req, data)
}


