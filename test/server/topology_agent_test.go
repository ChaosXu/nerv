package server

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
)

func TestInstallAgent(t *testing.T) {
	create(t, "Class", "classes/nerv/agent.json")
	create(t, "Class", "classes/host/host.json")
	template := create(t, "ServiceTemplate", "templates/agent_template.json")
	assert.NotNil(t, template)
	id := reflect.ValueOf(template).Elem().FieldByName("ID").Interface()
	ret, err := invoke(t, "ServiceTemplate", id, "CreateTopology", "agentTopology")
	if err != nil {
		assert.Nil(t, err, err.Error())
	}

	retObj := reflect.ValueOf(ret).Elem().Index(0).Interface().(map[string]interface{})
	invoke(t, "Topology", int(retObj["ID"].(float64)), "Install")
}
