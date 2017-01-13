package repository

import (
	"github.com/toolkits/file"
	"github.com/ChaosXu/nerv/lib/resource/model"
)

type StandaloneRepository struct {
	Root string
}

func NewStandaloneScriptRepository(root string) *StandaloneRepository {
	return &StandaloneRepository{root}
}

func (p *StandaloneRepository) Get(path string) (*model.Script, error) {
	if content, err := file.ToString(p.Root + "/scripts" + path); err != nil {
		return nil, err
	} else {
		return &model.Script{Content:content}, nil
	}
}
