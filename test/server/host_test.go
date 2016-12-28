package server

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
)

func TestHostInstall(t *testing.T) {

	host := map[string]interface{}{
		"name":"host-1",
		"ip":"127.0.0.1",
		"template":"/nerv/host_template.json",
	}
	created := createObj(t, "Host", host)
	assert.NotNil(t, created)
	id := reflect.ValueOf(created).Elem().FieldByName("ID").Interface()
	_, err := invoke(t, "Host", id, "Install")
	if err != nil {
		assert.Nil(t, err, err.Error())
	}
}

//func TestNervInstall(t *testing.T) {
//	nerv := create(t, "Topology", "data/nerv_topo.json")
//	assert.NotNil(t, nerv)
//	id := reflect.ValueOf(nerv).Elem().FieldByName("ID").Interface()
//	_, err := invoke(t, "Nerv", id, "Install")
//	if err != nil {
//		assert.Nil(t, err, err.Error())
//	}
//}

