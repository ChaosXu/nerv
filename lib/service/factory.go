package service

// ServiceFactory create a service
type ServiceFactory interface {
	// Init factory
	Init() error

	// Get return a service
	Get() interface{}

	// Dependencies return all dependency service name
	Dependencies() []string
}

// Factory create a service
type Factory interface {

	// New a service instance
	New() interface{}
}
