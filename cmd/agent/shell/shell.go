package shell

import (
	"os/exec"
	"github.com/ChaosXu/nerv/lib/rpc"
	"fmt"
	"log"
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
	export := ""
	for k, v := range script.Args {
		export = export + fmt.Sprintf(" %s=%s", k, v)
	}
	shell := "export " + export + " && " + script.Content
	log.Println(shell)

	out, err := exec.Command("/bin/bash","-c",shell).Output()
	if err != nil {
		return err
	}
	log.Println(string(out))

	return nil
}
