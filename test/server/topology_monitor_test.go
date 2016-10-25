package server

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
)

func TestInstallMonitor(t *testing.T) {
	//create(t, "Class", "../../classes/nerv/agent.json")
	//create(t, "Class", "../../classes/nerv/echost.json")
	//create(t, "Class", "../../classes/nerv/monitor.json")
	template := create(t, "ServiceTemplate", "templates/monitor_template.json")
	assert.NotNil(t, template)
	id := reflect.ValueOf(template).Elem().FieldByName("ID").Interface()
	ret, err := invoke(t, "ServiceTemplate", id, "CreateTopology", "monitorTopology")
	if err != nil {
		assert.Nil(t, err, err.Error())
	}

	retObj := reflect.ValueOf(ret).Elem().Index(0).Interface().(map[string]interface{})
	invoke(t, "Topology", int(retObj["ID"].(float64)), "Install")
	//invoke(t, "Topology", int(retObj["ID"].(float64)), "Stop")
	//invoke(t, "Topology", int(retObj["ID"].(float64)), "Start")
	//invoke(t, "Topology", int(retObj["ID"].(float64)), "Stop")
	//invoke(t, "Topology", int(retObj["ID"].(float64)), "Uninstall")
}
