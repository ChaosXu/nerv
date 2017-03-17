package service

import (
	"github.com/ChaosXu/nerv/lib/service"
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/ChaosXu/nerv/cmd/agent/model"
	"github.com/ChaosXu/nerv/lib/util"
)

func init() {
	service.Registry.Put("App", &AppServiceFactory{})
}

// AppServiceFactory
type AppServiceFactory struct {
	appService *AppService
}

func (p *AppServiceFactory) Init() error {
	p.appService = &AppService{}
	return nil
}

func (p *AppServiceFactory) Get() interface{} {
	return p.appService
}

func (p *AppServiceFactory) Dependencies() []string {
	return nil
}

// AppService
type AppService struct {

}

func (p *AppService) Update(name string, attrs []string, values []interface{}) error {
	app := &model.App{}
	if err := db.DB.First(app, "Name = ?", name).Error; err != nil {
		return nil
	}

	err := util.SetValue(app, attrs, values)
	if err != nil {
		return err
	}
	if err := db.DB.Save(app).Error; err!=nil {
		return err
	} else {
		return nil
	}
}
