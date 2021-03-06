package service

import (
	_ "github.com/ChaosXu/nerv/lib/automation/model"
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/ChaosXu/nerv/lib/env"
	_ "github.com/ChaosXu/nerv/lib/monitor/model"
	"github.com/ChaosXu/nerv/lib/net/http/rest"
	"github.com/ChaosXu/nerv/lib/net/http/rest/middleware"
	_ "github.com/ChaosXu/nerv/lib/user/model"
	user "github.com/ChaosXu/nerv/lib/user/model"
	"github.com/pressly/chi"
	chim "github.com/pressly/chi/middleware"
	"github.com/pressly/chi/render"
	"log"
	"net/http"
)

// HttpService
type HttpService struct {
	Controller *rest.RestController `inject:"RestController"`
}

func (p *HttpService) Init() error {
	r := chi.NewRouter()
	r.Use(chim.Logger)
	r.Use(middleware.ParamsParser)

	r.Route("/api/objs/Login", func(r chi.Router) {
		r.Post("/", login)
	})

	r.Route("/api/objs/:class", func(r chi.Router) {
		//TBD:don't use server rest api
		r.Get("/", p.Controller.List)
		r.Post("/", p.Controller.Create)
		r.Put("/", p.Controller.Update)
		r.Route("/:id", func(r chi.Router) {
			r.Get("/", p.Controller.Get)
			r.Delete("/", p.Controller.Remove)
			r.Post("/", p.Controller.InvokeServiceFunc())
			r.Post("/:method", p.Controller.InvokeObj)
		})
	})
	port := env.Config().GetMapString("http", "port", "3333")
	go func() {
		log.Fatal(http.ListenAndServe(":"+port, r))
	}()
	return nil
}

func handlePanic(w http.ResponseWriter, req *http.Request) {
	//if r := recover(); r != nil {
	//	fmt.Println(r)
	//	render.Status(req, 500)
	//	render.JSON(w, req, r)
	//}
}

func login(w http.ResponseWriter, req *http.Request) {
	defer handlePanic(w, req)

	account := &user.Account{}
	if err := render.Bind(req.Body, account); err != nil {
		render.Status(req, 400)
		render.JSON(w, req, err.Error())
		return
	}

	var ret user.Account
	//TBD: using hash
	db := db.DB.Where("name=? and password=?", account.Name, account.Password).First(&ret)
	if err := db.Error; err != nil {
		render.Status(req, 500)
		render.JSON(w, req, err.Error())
		return
	}

	if db.RecordNotFound() {
		render.Status(req, 404)
		render.JSON(w, req, map[string]string{"Name": account.Name})
	} else {
		render.Status(req, 200)
		render.JSON(w, req, map[string]string{"Name": account.Name})
	}
}
