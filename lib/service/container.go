package service

// Container manage all services
type Container struct {
	factories map[string]ServiceFactory
	inited    []string
}

func NewContainer() *Container {
	return &Container{factories:map[string]ServiceFactory{}, inited:[]string{}}
}

// Add factory
func (p *Container) Add(name string, factory ServiceFactory) {
	p.factories[name] = factory
}

// Start container
func (p *Container) Start() {
	if len(p.inited) > 0 {
		return
	}

	fp := map[string]ServiceFactory{}
	for n, f := range p.factories {
		p.register(n, f, fp)
	}

	for i, n := range p.inited {
		f := p.factories[n]
		if err := f.Init(); err != nil {
			p.dispose(i)
			break
		}
		svc := f.Get()
		if init, ok := svc.(Initializer); ok {
			if err := init.Init(); err != nil {
				p.dispose(i)
				break
			}
		}
	}
}

// Stop container and release all service's resource
func (p *Container) Stop() {

}

func (p *Container) register(name string, factory ServiceFactory, fp map[string]ServiceFactory) {
	if fp[name] != nil {
		return
	}

	deps := factory.Dependencies()
	if deps == nil {
		return
	}

	for _, dep := range deps {
		depF := p.factories[dep]
		if depF != nil {
			p.register(dep, depF, fp)
		}
	}

	p.inited = append(p.inited, name)
	fp[name] = factory
}

func (p *Container) dispose(l int) {
	for i := l - 1; i >= 0; i-- {
		n := p.inited[i]
		f := p.factories[n]
		svc := f.Get()
		if d, ok := svc.(Disposable); ok {
			d.Dispose()
		}
	}
}
