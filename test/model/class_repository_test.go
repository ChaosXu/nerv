package model

import (
	"testing"

	"github.com/chaosxu/nerv/lib/model"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	model.InitClassRepository("classes")
	class := model.GetClassRepository().Find("/host/Host")
	assert.NotNil(t, class, "/host/Host isn't exists")
	assert.Equal(t, "/host/Host", class.Name)
	assert.Equal(t, "/Resource", class.Base)

	var ops = class.Operations;
	assert.Len(t, ops, 5)
}
