package init

import (
	"fmt"
	"github.com/ChaosXu/nerv/lib/resource/model"
)

// InitExecutor perform an operation when init nerv
type InitExecutor struct {

}

func (p *InitExecutor) Perform(class *model.Class, operation string, args map[string]string) error {
	fmt.Println(class.Name + "." + operation)
	return nil
}
