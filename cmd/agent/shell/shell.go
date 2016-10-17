package ssh

import (
	"github.com/ChaosXu/nerv/lib/rpc"
	"fmt"
)

func init() {
	rpc.Register(&Shell{})
}

//Shell execute a script
type Shell int

//Execute a script
func (p *Shell) Execute(script string, vars map[string]string, out string) error {
	return fmt.Errorf("TBD")
}
