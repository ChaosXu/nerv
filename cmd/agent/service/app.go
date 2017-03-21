package service

import (
	"github.com/ChaosXu/nerv/cmd/agent/model"
	"github.com/ChaosXu/nerv/lib/brick"
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/ChaosXu/nerv/lib/util"
)

// AppService
type AppService struct {
	brick.Trigger
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
	if err := db.DB.Save(app).Error; err != nil {
		return err
	} else {
		p.Emmit("Update", name)
		return nil
	}
}
