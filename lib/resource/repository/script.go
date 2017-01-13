package repository

import "github.com/ChaosXu/nerv/lib/resource/model"

// ScriptRepository manage scripts of all resources
type ScriptRepository interface {
	Get(path string) (*model.Script, error)
}
