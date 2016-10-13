package server

import (
	"testing"
	"reflect"
	"time"
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/chaosxu/nerv/lib/automation/model"
)

func TestClassRest(t *testing.T) {
	testCRUD(t, "Class", "classes/host/host.json")
}

func TestUpdateAddChild(t *testing.T) {
	data := create(t, "Class", "classes/host/host.json").(*model.Class)
	data.Operations = append(data.Operations, model.Operation{Name:"updateAddChile", Type:"go", Implementor:"test"})
	update(t, "Class", data)
}

func TestUpdateRemoveChild(t *testing.T) {
	data := create(t, "Class", "classes/host/host.json").(*model.Class)
	data.Operations = append(data.Operations, model.Operation{Name:"updateAddChile", Type:"go", Implementor:"test"})
	data = update(t, "Class", data).(*model.Class)

	fmt.Printf("%d\n", len(data.Operations))
	data.Operations = data.Operations[:len(data.Operations) - 1]
	fmt.Printf("%d\n", len(data.Operations))
	update(t, "Class", data)
}

func TestUpdateAddRemoveAddChild(t *testing.T) {
	data := create(t, "Class", "classes/host/host.json").(*model.Class)
	data.Operations = append(data.Operations, model.Operation{Name:"uar-a", Type:"go", Implementor:"test"})

	time.Sleep(time.Second)
	data = update(t, "Class", data).(*model.Class)

	fmt.Printf("%d\n", len(data.Operations))
	data.Operations = data.Operations[:len(data.Operations) - 1]
	fmt.Printf("%d\n", len(data.Operations))
	data.Operations = append(data.Operations, model.Operation{Name:"uar-ara", Type:"go", Implementor:"test"})

	time.Sleep(time.Second)
	update(t, "Class", data)
}

func TestUpdateUpdateChild(t *testing.T) {

}

func TestInvoke(t *testing.T) {
	template := create(t, "ServiceTemplate", "templates/test_service_template.json")
	assert.NotNil(t, template)

	//update(t, "Topology", topology)
	v := reflect.ValueOf(template).Elem()
	id := v.FieldByName("ID").Interface()
	ret, err := invoke(t, "ServiceTemplate", id, "CreateTopology", "topology1")
	if err != nil {
		assert.Nil(t, err, err.Error())
	}
	assert.NotNil(t, ret)
	fmt.Printf("%+v\n", ret)
	retObj := reflect.ValueOf(ret).Elem().Index(0).Interface().(map[string]interface{})
	fmt.Printf("%+v\n", retObj)
	fmt.Printf("%+v\n", retObj["ID"])
	remove(t, "Topology", int(retObj["ID"].(float64)))
	remove(t, "ServiceTemplate", id)
}

func testCRUD(t *testing.T, class string, dataPath string) {
	data := create(t, class, dataPath)

	find(t, class)

	v := reflect.ValueOf(data).Elem()
	id := v.FieldByName("ID").Interface()
	data = getAndPreLoad(t, class, id)

	time := v.FieldByName("UpdatedAt").Interface().(time.Time)
	fmt.Println(time)
	data = update(t, class, data)

	remove(t, class, id)

	getNil(t, class, id)
}



