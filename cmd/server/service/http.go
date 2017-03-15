package lib

import (
	"github.com/ChaosXu/nerv/lib/service"
	"github.com/pressly/chi"
	chim "github.com/pressly/chi/middleware"
	"github.com/ChaosXu/nerv/lib/net/http/rest/middleware"
	"github.com/ChaosXu/nerv/lib/net/http/rest"
	"github.com/ChaosXu/nerv/lib/env"
	"log"
	"net/http"
)

func init() {
	service.Registry.Put("http", &HttpServiceFactory{})
}

type HttpServiceFactory struct {
	httpService *HttpService
}

func (p *HttpServiceFactory) Init() error {
	p.httpService = newHttpService()
	return nil
}

func (p *HttpServiceFactory) Get() interface{} {
	return p.httpService
}

func newHttpService() *HttpService {
	return &HttpService{}
}

type HttpService struct {
}

func (p *HttpService) Init() error {
	r := chi.NewRouter()
	r.Use(chim.Logger)
	r.Use(middleware.ParamsParser)

	r.Route("/api/objs/:class", func(r chi.Router) {
		//TBD:don't use server rest api
		r.Get("/", rest.List)
		r.Post("/", rest.Create)
		r.Put("/", rest.Update)
		r.Route("/:id", func(r chi.Router) {
			r.Get("/", rest.Get)
			r.Delete("/", rest.Remove)
			r.Post("/", rest.InvokeService)
			r.Post("/:method", rest.InvokeObj)
		})
	})
	port := env.Config().GetMapString("http", "port", "3333")
	go func() {
		log.Fatal(http.ListenAndServe(":"+port, r))
	}()
	return nil
}
