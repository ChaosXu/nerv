package model

import (
	"testing"
	"log"
	"github.com/chaosxu/nerv/model"
	"github.com/chaosxu/nerv/template"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestInstall(t *testing.T) {
	model.InitClassRepository("./classes")
	template, err := template.GetServiceTemplate("test_service_template.json")
	assert.Nil(t, err)
	topology := model.NewTopology(template)
	assert.NotNil(t, topology)

	topology.Install()

	for _, node := range topology.Nodes {
		assert.Nil(t, node.Error, fmt.Sprintf("node %s has error %s", node.Template.Name, node.Error))
	}
}

func TestNewTopology(t *testing.T) {
	model.InitClassRepository("./classes")
	template, err := template.GetServiceTemplate("test_service_template.json")
	assert.Nil(t, err)

	topology := model.NewTopology(template)
	nodes := topology.Nodes
	assert.NotNil(t, nodes)

	assert.NotNil(t, nodes["loadBalance"])
	assert.NotNil(t, nodes["webServer"])
	assert.NotNil(t, nodes["mysql"])
	assert.NotNil(t, nodes["host"])

	lb := nodes["loadBalance"]
	assert.Equal(t, "loadBalance", lb.Template.Name)
	assert.Nil(t, lb.Error)
	log.Printf("status %d %d", lb.Status, model.NodeStatusNew)
	log.Printf("status %t", lb.Status == model.NodeStatusNew)
	assert.EqualValues(t, model.NodeStatusNew, lb.Status)
}

