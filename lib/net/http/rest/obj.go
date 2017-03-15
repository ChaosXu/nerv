package rest

import (
	"net/http"
	"fmt"
	"strings"
	"reflect"
	"strconv"
	"github.com/pressly/chi/render"
	"github.com/jinzhu/gorm"
	"github.com/ChaosXu/nerv/lib/net/http/rest/middleware"
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/ChaosXu/nerv/lib/service"
	"encoding/json"
	"log"
)

type User struct {
	Name string
}

func handlePanic(w http.ResponseWriter, req *http.Request) {
	//if r := recover(); r != nil {
	//	fmt.Println(r)
	//	render.Status(req, 500)
	//	render.JSON(w, req, r)
	//}
}



func List(w http.ResponseWriter, req *http.Request) {
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

	//count
	d := db.DB.Model(md.NewSlice());
	var count int64
	if where != "" {
		d = d.Where(where, args)
	}
	if err := d.Count(&count).Error; err != nil {
		render.Status(req, 500)
		render.JSON(w, req, err.Error())
		return
	}


	//order page
	var page, pageCount, limit int64
	limit = 10
	var err error
	paramPage := middleware.CurrentParams(req).QueryParam("page")
	paramSize := middleware.CurrentParams(req).QueryParam("pageSize")

	if paramSize != "" {
		limit, err = strconv.ParseInt(paramSize, 10, 32)
		if err != nil {
			limit = 10
		}
	}
	pageCount = count / limit
	if count % limit >= 0 {
		pageCount += 1
	}

	if paramPage != "" {
		page, err = strconv.ParseInt(paramPage, 10, 32)
		if err != nil {
			page = 0
		}
	}
	if page >= pageCount {
		page = pageCount - 1
	}

	if where != "" {
		d = db.DB.Where(where, args)
	}

	order := middleware.CurrentParams(req).QueryParam("order")
	if order != "" {
		d = d.Order(order);
	}
	data := md.NewSlice()
	if d.Offset(page * limit).Limit(limit).Find(data).RecordNotFound() {
		render.Status(req, 200)
		render.JSON(w, req, data)
		return
	}

	render.Status(req, 200)
	render.JSON(w, req, map[string]interface{}{"data":data, "page":page, "pageSize":limit, "pageCount":pageCount})
}

// get one obj. query params: assocations=a,b...
func Get(w http.ResponseWriter, req *http.Request) {
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

func Create(w http.ResponseWriter, req *http.Request) {
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

func Remove(w http.ResponseWriter, req *http.Request) {
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

func Update(w http.ResponseWriter, req *http.Request) {
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

func InvokeService(w http.ResponseWriter, req *http.Request) {
	defer handlePanic(w, req)

	class := middleware.CurrentParams(req).PathParam("class")
	methodName := middleware.CurrentParams(req).PathParam("id")

	log.Println("InvokeService")
	svc := service.Registry.Get(class)
	if svc == nil {
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

func InvokeObj(w http.ResponseWriter, req *http.Request) {
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


