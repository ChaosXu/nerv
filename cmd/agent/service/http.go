package service

import (
	"log"
	"net/http"
	"github.com/pressly/chi"
	chim "github.com/pressly/chi/middleware"
	"github.com/ChaosXu/nerv/lib/env"
	"github.com/ChaosXu/nerv/lib/net/http/rest"
	"github.com/ChaosXu/nerv/lib/net/http/rest/middleware"
)

type HttpService struct {
	Controller *rest.RestController `inject:"RestController"`
}

func (p *HttpService) Init() error {
	r := chi.NewRouter()
	r.Use(chim.Logger)
	r.Use(middleware.ParamsParser)

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
	port := env.Config().GetMapString("http", "port", "3335")
	go func() {
		log.Fatal(http.ListenAndServe(":" + port, r))
	}()
	return nil
}

