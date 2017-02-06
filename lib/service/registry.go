package service

import (
	"fmt"
)

var Registry *ServiceRegistry

func init() {
	Registry = &ServiceRegistry{Services:map[string]ServiceFactory{}}
}

// ServiceRegistry provide all services that will be called through the rest api
type ServiceRegistry struct {
	Services map[string]ServiceFactory
}

func (p *ServiceRegistry) Get(name string) (interface{}, error) {
	if sf := p.Services[name]; sf != nil {
		return sf.Get()
	} else {
		return nil, fmt.Errorf("%s isn't exists", name)
	}
}

func (p *ServiceRegistry) Put(name string, factory ServiceFactory) {
	p.Services[name] = factory
}


