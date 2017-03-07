package service

// Initializer init a service
type Initializer interface {
	Init() error
}

// Disposable dispose a service that release all resources which used
type Disposable interface {
	Dispose() error
}
