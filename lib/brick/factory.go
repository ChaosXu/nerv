package brick

// Factory create a service
type Factory interface {

	// New a service instance
	New() interface{}
}
