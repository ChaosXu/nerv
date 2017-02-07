package repository

import (
	"fmt"
	"github.com/ChaosXu/nerv/lib/resource/model"
	"github.com/toolkits/file"
	"encoding/json"
)

type FileClassRepository struct {
	root string
}

func NewFileClassRepository(rootPath string) *FileClassRepository {
	return &FileClassRepository{root:rootPath}
}

func (p *FileClassRepository) Get(path string) (*model.Class, error) {
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

func (p *FileClassRepository) InheritFrom(class *model.Class, base string) (*model.Class, error) {
	if class.Base == "" {
		return nil, fmt.Errorf("class.Base is empty.class=%s", class.Name)
	}

	baseClass, err := p.Get(class.Base)
	if err != nil {
		return nil, fmt.Errorf("class.Base isn't exist,%s. class.Base=%s", err.Error(), class.Base)
	}

	if baseClass.Name == base {
		return baseClass, nil
	} else {
		return p.InheritFrom(baseClass, base)
	}
}

func (p *FileClassRepository) GetOperation(class *model.Class, name string) (*model.Operation, error) {
	op := class.GetOperation(name)
	if op != nil {
		op.DefineClass = class.Name
		return op, nil
	}
	base := class.Base
	if base == "" {
		return nil, fmt.Errorf("operation isn't exists.class=%s,operation=%s", class.Name, name)
	}
	baseClass, err := p.Get(base)
	if err != nil {
		return nil, fmt.Errorf("could not get operation, because %s.class=%s,operation=%", err.Error(), class.Name, name)
	}
	return p.GetOperation(baseClass, name)
}
