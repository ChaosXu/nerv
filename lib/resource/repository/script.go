package repository

import "github.com/ChaosXu/nerv/lib/resource/model"

// ScriptRepository manage scripts of all resources
type ScriptRepository interface {
	// Get script from {class}/{path}
	Get(class string, path string) (*model.Script, error)
}
