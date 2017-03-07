package service

import "github.com/ChaosXu/nerv/lib/service"

func init() {
	service.Registry.Put("LogFile", &LogConfigServiceFactory{})
}

type LogConfigServiceFactory struct {
	logConfigService *LogConfigService
}

func (p *LogConfigServiceFactory) Init() error {
	p.logConfigService = &LogConfigService{}
	return nil
}

func (p *LogConfigServiceFactory) Get() interface{} {
	return p.logConfigService
}

type LogConfigService struct {

}

func (p *LogConfigService) Add(file string) error {
	return nil
}
