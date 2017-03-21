package repository

import (
	"github.com/ChaosXu/nerv/lib/resource/model"
	"github.com/toolkits/file"
)

type HttpScriptRepository struct {
	Root string
}

func NewHttpScriptRepository(root string) *HttpScriptRepository {
	return &HttpScriptRepository{root}
}

func (p *HttpScriptRepository) Get(class string, path string) (*model.Script, error) {
	if content, err := file.ToString(p.Root + class + "/" + path); err != nil {
		return nil, err
	} else {
		return &model.Script{Content: content}, nil
	}
}
