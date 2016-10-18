package golang

//Models register all resources
var Models map[string]Resource = map[string]Resource{}

//Resource operations
type Resource interface {
	//Create a resource
	Create() error

	//Delete a resource
	Delete() error

	//Start a resource
	Start() error

	//Stop a resource
	Stop() error

	//Setup a resource
	Setup() error
}
