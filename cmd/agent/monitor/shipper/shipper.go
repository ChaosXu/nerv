package shipper

//Shipper upload data to the server
type Shipper interface {
	Send(v interface{})
}
