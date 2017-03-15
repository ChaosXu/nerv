package model

import "github.com/jinzhu/gorm"

// App managed by agent
type App struct {
	gorm.Model
	Name    string
	Path    string
	Version string
}
