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

}

func (p *ExecutorImpl) Perform(class *model.Class, operation string, args map[string]string) error {
	fmt.Println(class.Name + "." + operation)
	return nil
}
