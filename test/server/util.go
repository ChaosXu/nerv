package server

import (
	"github.com/go-resty/resty"
	"fmt"
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"reflect"
	"github.com/toolkits/file"
	"strings"
	"github.com/ChaosXu/nerv/lib/db"
	_"github.com/ChaosXu/nerv/lib/automation/model"
	"log"
)

func find(t *testing.T, class string) {
	logCodeLine()
	md := db.Models[class]

	url := fmt.Sprintf("%s/objs/%s?where=name=?&values=/host/Host", ServerUrl, class)
	fmt.Println(url)
	res, err := resty.R().
			SetHeader("Content-Type", "application/json").
			Get(url)
	if err != nil {
		assert.Nil(t, err, err.Error())
	}
	assert.Equal(t, 200, res.StatusCode())
	b := res.Body();
	//fmt.Printf("body\n %s\n", string(b))

	data := md.NewSlice()
	err = json.Unmarshal(b, data)
	assert.Nil(t, err)

	fmt.Println("find")
	fmt.Println(data)
}

func remove(t *testing.T, class string, id interface{}) {
	logCodeLine()
	var (
		err error
		res *resty.Response
	)

	res, err = resty.R().
			SetHeader("Content-Type", "application/json").
			Delete(fmt.Sprintf("%s/objs/%s/%d", ServerUrl, class, id))

	if err != nil {
		assert.Nil(t, err, err.Error())
	}
	assert.Equal(t, 200, res.StatusCode(), string(res.Body()))
}

func update(t *testing.T, class string, data interface{}) interface{} {
	logCodeLine()
	var (
		body []byte
		err error
		res *resty.Response
	)

	if body, err = json.Marshal(data); err != nil {
		assert.NotNil(t, err, err.Error())
	}
	res, err = resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Put(fmt.Sprintf("%s/objs/%s", ServerUrl, class))

	if err != nil {
		assert.Nil(t, err, err.Error())
	}
	assert.Equal(t, 200, res.StatusCode())
	b := res.Body();
	//fmt.Printf("body\n %s\n", string(b))
	md := db.Models[class]
	updated := md.New()
	err = json.Unmarshal(b, updated)
	assert.Nil(t, err)
	return data
}

func createObj(t *testing.T, class string, obj interface{}) interface{} {
	logCodeLine()
	var (
		body []byte
		err error
		res *resty.Response
	)
	if body, err = json.Marshal(obj); err != nil {
		assert.NotNil(t, err, err.Error())
	}

	res, err = resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post(fmt.Sprintf("%s/objs/%s", ServerUrl, class))

	if err != nil {
		assert.Nil(t, err, err.Error())
	}
	assert.Equal(t, 200, res.StatusCode(), string(res.Body()))
	b := res.Body();
	//fmt.Printf("body\n %s\n", string(b))
	md := db.Models[class]
	data := md.New()
	err = json.Unmarshal(b, data)
	assert.Nil(t, err)
	return data
}

func create(t *testing.T, class string, dataPath string) interface{} {
	logCodeLine()
	var (
		body string
		err error
		res *resty.Response
	)
	body, err = file.ToTrimString(dataPath)
	assert.Nil(t, err)
	res, err = resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post(fmt.Sprintf("%s/objs/%s", ServerUrl, class))

	if err != nil {
		assert.Nil(t, err, err.Error())
	}
	assert.Equal(t, 200, res.StatusCode(), string(res.Body()))
	b := res.Body();
	//fmt.Printf("body\n %s\n", string(b))
	md := db.Models[class]
	data := md.New()
	err = json.Unmarshal(b, data)
	assert.Nil(t, err)
	return data
}

func getAndPreLoad(t *testing.T, class string, id interface{}) interface{} {
	logCodeLine()
	md := db.Models[class]
	assList := associations(reflect.TypeOf(md.Type).Elem(), "", []string{})
	ass := strings.Join(assList, ",")
	var url string
	if ass == "" {
		url = fmt.Sprintf("%s/objs/%s", ServerUrl, class, id)
	} else {
		url = fmt.Sprintf("%s/objs/%s/%d?associations=%s", ServerUrl, class, id, ass)
	}
	res, err := resty.R().
			SetHeader("Content-Type", "application/json").
			Get(url)
	if err != nil {
		assert.Nil(t, err, err.Error())
	}
	assert.Equal(t, 200, res.StatusCode())
	b := res.Body();
	//fmt.Printf("body\n %s\n", string(b))

	data := md.New()
	err = json.Unmarshal(b, data)
	assert.Nil(t, err)

	v := reflect.ValueOf(data).Elem()
	for _, as := range assList {
		assertAssociations(t, v, as)
	}
	return data
}

func invoke(t *testing.T, class string, id interface{}, method string, args ...interface{}) (interface{}, error) {
	logCodeLine()
	var (
		body string
		err error
		res *resty.Response
	)

	if b, err := json.Marshal(args); err != nil {
		assert.Nil(t, err, err.Error())
	} else {
		body = string(b)
	}
	assert.Nil(t, err)
	res, err = resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(body).
			Post(fmt.Sprintf("%s/objs/%s/%d/%s", ServerUrl, class, id, method))

	if err != nil {
		assert.Nil(t, err, err.Error())
	}
	assert.Equal(t, 200, res.StatusCode(), string(res.Body()))
	b := res.Body();
	log.Printf("%+vs", string(b))
	ro := &[]interface{}{}
	if err := json.Unmarshal(b, ro); err != nil {
		assert.Nil(t, err, err.Error())
	}
	return ro, err
}

func getNil(t *testing.T, class string, id interface{}) interface{} {
	logCodeLine()
	md := db.Models[class]
	assList := associations(reflect.TypeOf(md.Type).Elem(), "", []string{})
	ass := strings.Join(assList, ",")
	var url string
	if ass == "" {
		url = fmt.Sprintf("%s/objs/%s", ServerUrl, class, id)
	} else {
		url = fmt.Sprintf("%s/objs/%s/%d?associations=%s", ServerUrl, class, id, ass)
	}
	res, err := resty.R().
			SetHeader("Content-Type", "application/json").
			Get(url)
	if err != nil {
		assert.Nil(t, err, err.Error())
	}
	assert.Equal(t, 200, res.StatusCode())
	b := res.Body();
	//fmt.Printf("body\n %s\n", string(b))

	data := md.New()
	err = json.Unmarshal(b, data)
	assert.Nil(t, err)
	assert.Equal(t, "null\n", string(b))
	fmt.Println("return body:" + string(b))

	return data
}

func assertAssociations(t *testing.T, v reflect.Value, ass string) {
	assList := strings.Split(ass, ".")
	as := v.Type().Name()
	cv := v
	for _, ass := range assList {
		switch  cv.Kind(){
		case reflect.Slice:
			cv = cv.FieldByName(ass)
			as = as + "." + ass
			fmt.Printf("assert %s\n", as)
			if cv.Len() > 0 {
				cv = cv.Index(0)
				assert.NotNil(t, cv.Interface(), as + " is nil")
			}
		case reflect.Struct:
			cv = cv.FieldByName(ass)
			as = as + "." + ass
			fmt.Printf("assert %s\n", as)
			assert.NotNil(t, cv.Interface(), as + " is nil")
		}
	}
}

func associations(t reflect.Type, parent string, ass []string) []string {
	n := t.NumField()
	if parent != "" {
		parent = parent + "."
	}
	for i := 0; i < n; i++ {
		f := t.Field(i)

		switch f.Type.Kind() {
		case reflect.Slice:
			fmt.Printf("Name:%s,Type:%s,Kind:%s\n", f.Name, f.Type.Elem().Name(), f.Type.Kind())
			ass = append(ass, f.Name)
			return associations(f.Type.Elem(), parent + f.Name, ass)
		case reflect.Struct:
			fmt.Printf("Name:%s,Type:%s,Kind:%s\n", f.Name, f.Type.Name(), f.Type.Kind())
			//Not embedded type
			if f.Name != f.Type.Name() {
				ass = append(ass, f.Name)
				return associations(f.Type, parent + f.Name, ass)
			}
		}
	}

	return ass
}

