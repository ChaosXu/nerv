package repository

import (
	"github.com/toolkits/file"
	"github.com/ChaosXu/nerv/lib/resource/model"
)

type FileScriptRepository struct {
	Root string
}

func NewFileScriptRepository(root string) *FileScriptRepository {
	return &FileScriptRepository{root}
}

func (p *FileScriptRepository) Get(class string, path string) (*model.Script, error) {
	if content, err := file.ToString(p.Root + class + "/" + path); err != nil {
		return nil, err
	} else {
		return &model.Script{Content:content}, nil
	}
}
