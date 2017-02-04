package rest

var Services *ServiceRegistry

func init() {
	Services = &ServiceRegistry{services:map[string]interface{}{}}
}

// ServiceRegistry provide all services that will be called through the rest api
type ServiceRegistry struct {
	services map[string]interface{}
}

func (p *ServiceRegistry) Get(name string) interface{} {
	return p.services[name]
}

func (p *ServiceRegistry) Put(name string, svc interface{}) {
	p.services[name] = svc
}


