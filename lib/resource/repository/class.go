package repository

import (
	"github.com/ChaosXu/nerv/lib/resource/model"
	"github.com/toolkits/file"
	"encoding/json"
)

// ClassRepository  manage class of resource
type ClassRepository interface {
	// Get one class
	Get(path string) (*model.Class, error)
}

type LocalClassRepository struct {
	root string
}

func NewLocalClassRepository(rootPath string) *LocalClassRepository {
	return &LocalClassRepository{root:rootPath}
}

func (p *LocalClassRepository) Get(path string) (*model.Class, error) {
	content, err := file.ToBytes(p.root + path + "/type.json")
	if err != nil {
		return nil, err
	}

	class := &model.Class{}
	if err = json.Unmarshal(content, class); err != nil {
		return nil, err
	}

	return class, nil
}
