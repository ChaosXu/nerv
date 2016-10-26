package monitor

//Transfer upload samples to server
type Transfer interface {
	Send(sample *Sample)
}

type DefaultTransfer struct{

}

func NewTransfer() Transfer {
	return &DefaultTransfer{}
}

func (p *DefaultTransfer) Send(sample *Sample) {
	//TBD
}
