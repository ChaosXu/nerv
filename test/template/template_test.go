package template

import (
	"testing"

	"github.com/chaosxu/nerv/lib/template"
	"github.com/stretchr/testify/assert"
)

func TestGetServiceTemplate(t *testing.T) {
	template, err := template.GetServiceTemplate("test_service_template.json")
	assert.Nil(t, err)

	nodes := template.Nodes
	assert.NotNil(t, nodes)
	assert.Len(t, nodes, 4)

	n0 := nodes[0]
	assert.Equal(t, "loadBalance", n0.Name)
	assert.Equal(t, "/middleware/OpenResty", n0.Type)

	deps := n0.Dependencies
	assert.NotNil(t,deps)
	assert.Equal(t,"contained",deps[0].Type)
	assert.Equal(t,"host",deps[0].Target)
	assert.Equal(t,"connect",deps[1].Type)
	assert.Equal(t,"webServer",deps[1].Target)

	//n1 := nodes[1]
	//n2 := nodes[2]
	//n3 := nodes[3]
}
