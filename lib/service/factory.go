package service

// ServiceFactory create a service
type ServiceFactory interface {
	// Init factory
	Init() error

	// Get return a service
	Get() interface{}
}
