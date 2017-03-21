package repository

import (
	"github.com/ChaosXu/nerv/lib/automation/model/topology"
)

// TemplateRepository manage all templates
type TemplateRepository interface {
	// GetTemplate load template from storage
	GetTemplate(path string) (*topology.ServiceTemplate, error)
}
