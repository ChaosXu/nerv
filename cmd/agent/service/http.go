package service

import "github.com/ChaosXu/nerv/lib/service"

func init() {
	service.Registry.Put("http", &HttpServiceFactory{})
}

type HttpServiceFactory struct {
	httpService *HttpService
}

func (p *HttpServiceFactory) Init() error {
	p.httpService = newHttpService()
	return nil
}

func (p *HttpServiceFactory) Get() interface{} {
	return p.httpService
}

func newHttpService() *HttpService {
	return &HttpService{}
}

type HttpService struct {

}

func (p *HttpService) Start() error {
	return nil
}
