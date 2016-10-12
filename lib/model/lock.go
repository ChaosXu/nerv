package model

//Lock for mutex
type Lock struct {

}

//TryLock return true if the lock has been acquired
func (p *Lock) TryLock() bool {
	//TBD
	return false
}

//Lock will be wait until the lock has been acquired
func (p *Lock) Lock() {

}

//Unlock release the lock.Do nothing if no lock
func (p *Lock) Unlock() {

}


