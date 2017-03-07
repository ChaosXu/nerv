package service

import (
	_ "github.com/ChaosXu/nerv/lib/service"
	"github.com/pressly/chi"
	"net/http"
	"github.com/ChaosXu/nerv/lib/env"
	"log"
	"github.com/ChaosXu/nerv/lib/service"
	"github.com/pressly/chi/render"
	"fmt"
	"github.com/ChaosXu/nerv/lib/net/http/rest/middleware"
	"reflect"
	"encoding/json"
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
	r.Route("/api/objs/:class", func(r chi.Router) {
		//r.Get("/", rest.List)
		//r.Post("/", rest.Create)
		//r.Put("/", rest.Update)
		r.Route("/:id", func(r chi.Router) {
			//r.Get("/", rest.Get)
			//r.Delete("/", rest.Remove)
			r.Post("/", invokeService)
			//r.Post("/:method", rest.InvokeObj)
		})
	})
	port := env.Config().GetMapString("http", "port", "3335")
	go func() {
		log.Fatal(http.ListenAndServe(":" + port, r))
	}()
	return nil
}

func invokeService(w http.ResponseWriter, req *http.Request) {

	class := middleware.CurrentParams(req).PathParam("class")
	methodName := middleware.CurrentParams(req).PathParam("id")

	svc := service.Registry.Get(class)
	if svc != nil {
		render.Status(req, 404)
		render.JSON(w, req, fmt.Sprintf("service %s isn't exists", class))
		return
	}

	t := reflect.TypeOf(svc)
	if m, b := t.MethodByName(methodName); b != true {
		render.Status(req, 404)
		render.JSON(w, req, fmt.Sprintf("method %s.%s isn't exists.from %v", class, methodName, t))
		return

	} else {
		args := []json.RawMessage{}
		if err := render.Bind(req.Body, &args); err != nil {
			render.Status(req, 400)
			render.JSON(w, req, err.Error())
			return
		}

		in := []reflect.Value{reflect.ValueOf(svc)}
		funcType := m.Func.Type()

		for i, arg := range args {
			argType := funcType.In(i + 1)
			argValue := reflect.New(argType)
			if err := json.Unmarshal(arg, argValue.Interface()); err == nil {
				in = append(in, argValue.Elem())
			} else {
				render.Status(req, 500)
				render.JSON(w, req, err.Error())
				return
			}
		}

		values := m.Func.Call(in)
		ret := []interface{}{}
		httpCode := 200
		for _, value := range values {
			rawValue := value.Interface()
			if e, ok := rawValue.(error); ok {
				httpCode = 500
				ret = append(ret, e.Error())
			} else {
				ret = append(ret, rawValue)
			}

		}
		render.Status(req, httpCode)
		render.JSON(w, req, ret)
	}
}
