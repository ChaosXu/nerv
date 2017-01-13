package environment

import (
	"github.com/ChaosXu/nerv/lib/resource/model"
	"github.com/ChaosXu/nerv/lib/resource/repository"
)

// Executor perform an operation of class.
type Executor interface {
	// Perform an operation of class.
	Perform(class *model.Class, operation string, args map[string]string) error
}

// ExecutorImpl select the environment by class and invoke it
type ExecutorImpl struct {
	Standalone Environment `inject:"env_standalone"`
	ClassRep   repository.ClassRepository `inject:""`
}

func (p *ExecutorImpl) Perform(class *model.Class, operation string, args map[string]string) error {
	env, err := p.findEnvironment(class)
	if err != nil {
		return err
	}

	op, err := p.ClassRep.GetOperation(class, operation)
	if err != nil {
		return err
	}

	return env.Exec(class, op, args)
}

func (p *ExecutorImpl)findEnvironment(class *model.Class) (Environment, error) {
	if _, err := p.ClassRep.InheritFrom(class, "/nerv/StandaloneProcess"); err != nil {
		return nil, err
	} else {
		return p.Standalone, nil
	}
}
