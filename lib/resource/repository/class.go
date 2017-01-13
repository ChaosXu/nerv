package repository

import (
	"github.com/ChaosXu/nerv/lib/resource/model"
)

// ClassRepository  manage class of all resources
type ClassRepository interface {
	// Get one class
	Get(path string) (*model.Class, error)

	// InheritFrom return the base class from the current class
	InheritFrom(class *model.Class, base string) (*model.Class, error)

	// GetOperation return operation by name from class or it's ancestor
	GetOperation(class *model.Class, name string) (*model.Operation, error)
}
