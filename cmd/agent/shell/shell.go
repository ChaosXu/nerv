package shell

import (
	"os/exec"
	"github.com/ChaosXu/nerv/lib/rpc"
	"fmt"
)

func init() {
	//var shell RemoteShell
	rpc.Register(new(RemoteShell))
}

//Shell execute a script
type RemoteShell int

type RemoteScript struct {
	Content string
	Args    map[string]string
}
//Execute a script in the host of the agent
func (p *RemoteShell) Execute(script *RemoteScript, response *string) error {
	out, err := exec.Command("/bin/bash","-c",script.Content).Output()
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}
