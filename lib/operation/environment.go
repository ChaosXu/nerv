package operation

import "github.com/ChaosXu/nerv/lib/resource/model"

// Environment where all resources deployed
type Environment interface {

	// Exec do operation
	Exec(class *model.Class, op *model.Operation, args map[string]string) error
}
