package shipper

//Shipper upload data to the server
type Shipper interface {

	//Init shipper
	Init() error

	//Send data to the server
	Send(v interface{})
}
