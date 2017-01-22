package environment

import (
	"github.com/ChaosXu/nerv/lib/resource/model"
	"github.com/ChaosXu/nerv/lib/resource/repository"
	"fmt"
)

// Executor perform an operation of class.
type Executor interface {
	// Perform an operation of class.
	Perform(env string, class *model.Class, operation string, args map[string]string) error
}

// ExecutorImpl select the environment by class and invoke it
type ExecutorImpl struct {
	Standalone Environment `inject:"env_standalone"`
	Ssh        Environment `inject:"env_ssh"`
	ClassRep   repository.ClassRepository `inject:""`
}

func (p *ExecutorImpl) Perform(envType string, class *model.Class, operation string, args map[string]string) error {
	fmt.Println("Executor.Perform " + class.Name + "." + operation)
	env, err := p.findEnvironment(envType)
	if err != nil {
		return err
	}

	op, err := p.ClassRep.GetOperation(class, operation)
	if err != nil {
		return err
	}

	return env.Exec(class, op, args)
}

func (p *ExecutorImpl)findEnvironment(envType string) (Environment, error) {
	switch envType {
	case "standalone":
		return p.Standalone, nil
	case "ssh":
		return p.Ssh, nil
	default:
		return nil, fmt.Errorf("unsupported environment")
	}
}
