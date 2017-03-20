package service

import (
	"log"
	"net/http"
	"github.com/pressly/chi/render"
	"github.com/pressly/chi"
	"github.com/ChaosXu/nerv/lib/net/http/rest/middleware"
	"github.com/ChaosXu/nerv/lib/net/http/rest"
	"github.com/ChaosXu/nerv/lib/env"
	"github.com/ChaosXu/nerv/lib/db"
	chim "github.com/pressly/chi/middleware"
	user "github.com/ChaosXu/nerv/lib/user/model"
	_ "github.com/ChaosXu/nerv/lib/automation/model"
	_ "github.com/ChaosXu/nerv/lib/monitor/model"
	_ "github.com/ChaosXu/nerv/lib/user/model"
)

// HttpService
type HttpService struct {
	Controller *rest.RestController `inject:RestController`
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
			r.Post("/", p.Controller.InvokeService)
			r.Post("/:method", p.Controller.InvokeObj)
		})
	})
	port := env.Config().GetMapString("http", "port", "3333")
	go func() {
		log.Fatal(http.ListenAndServe(":" + port, r))
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
		render.JSON(w, req, map[string]string{"Name":account.Name})
	} else {
		render.Status(req, 200)
		render.JSON(w, req, map[string]string{"Name":account.Name})
	}
}



