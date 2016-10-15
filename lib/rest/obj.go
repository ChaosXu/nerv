package rest

import (
	"net/http"
	"fmt"
	"strings"

	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
	"github.com/jinzhu/gorm"
	"github.com/ChaosXu/nerv/lib/rest/middleware"
	"github.com/ChaosXu/nerv/lib/db"
	_ "github.com/ChaosXu/nerv/lib/deploy/model"
	"reflect"
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
			r.Post("/:method", invoke)
		})
	})
}

func handlePanic(w http.ResponseWriter, req *http.Request) {
	//if r := recover(); r != nil {
	//	fmt.Println(r)
	//	render.Status(req, 500)
	//	render.JSON(w, req, r)
	//}
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
	md := db.Models[class]
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

	md := db.Models[class]
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
	md := db.Models[class]
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

	md := db.Models[class]
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
	md := db.Models[class]
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

func invoke(w http.ResponseWriter, req *http.Request) {
	defer handlePanic(w, req)

	class := middleware.CurrentParams(req).PathParam("class")
	id := middleware.CurrentParams(req).PathParam("id")
	methodName := middleware.CurrentParams(req).PathParam("method")

	md := db.Models[class]
	if md == nil {
		render.Status(req, 404)
		render.JSON(w, req, fmt.Sprintf("class %s isn't exists", class))
		return
	}

	data := md.New()
	if err := db.DB.First(data, id).Error; err != nil {
		render.Status(req, 404)
		render.JSON(w, req, err.Error())
		return
	}

	t := reflect.TypeOf(data)
	if m, b := t.MethodByName(methodName); b != true {
		render.Status(req, 404)
		render.JSON(w, req, fmt.Sprintf("%s/%s/%s isn't exists", class, id, methodName))
		return

	} else {
		args := []interface{}{}
		if err := render.Bind(req.Body, &args); err != nil {
			render.Status(req, 400)
			render.JSON(w, req, err.Error())
			return
		}

		in := []reflect.Value{reflect.ValueOf(data)}
		for _, arg := range args {
			in = append(in, reflect.ValueOf(arg))
		}
		values := m.Func.Call(in)
		ret := []interface{}{}
		for _, value := range values {
			ret = append(ret, value.Interface())
		}
		render.Status(req, 200)
		render.JSON(w, req, ret)
	}
}


