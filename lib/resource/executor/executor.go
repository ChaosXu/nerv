package executor

import (
	"github.com/ChaosXu/nerv/lib/resource/model"
	"fmt"
)

// Executor perform an operation of class.
type Executor interface {
	// Perform an operation of class.
	Perform(class *model.Class, operation string, args map[string]string) error
}

type ExecutorImpl struct {
	Local Executor `inject:"local"`
}

func (p *ExecutorImpl) Perform(class *model.Class, operation string, args map[string]string) error {
	op := class.GetOperation(operation)
	if op == nil {
		return fmt.Errorf("unsupported operation. class:%s,operation:%s", class.Name, operation)
	}

	worker := p.findExecutor(op.Type)
	if worker == nil {
		return fmt.Errorf("unsupported operation type. class:%s,operation:%s,type:%s", class.Name, operation, op.Type)
	}
	return worker.Perform(class, operation, args)
}
func (p *ExecutorImpl)findExecutor(opType string) Executor {
	switch opType {
	case "shell":
		return p.Local
	default:
		return nil
	}
}
