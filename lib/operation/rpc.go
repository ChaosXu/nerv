package operation

import (
	"fmt"
	"github.com/ChaosXu/nerv/lib/credential"
	"github.com/ChaosXu/nerv/lib/resource/model"
	"github.com/ChaosXu/nerv/lib/resource/repository"
	crpc "github.com/ChaosXu/nerv/lib/rpc"
	"net/rpc"
)

// RpcEnvironment where nerv add worker cluster.
type RpcEnvironment struct {
	ScriptRepository repository.ScriptRepository `inject:"script_rep_standalone"`
}

func (p *RpcEnvironment) Exec(class *model.Class, operation *model.Operation, args map[string]string) error {
	fmt.Printf("Rpc.Exec %s.%s %s\n", class.Name, operation.Name, operation.Implementor)

	script, err := p.ScriptRepository.Get(operation.DefineClass, operation.Implementor)
	if err != nil {
		return err
	}

	addr := args["address"]
	if addr == "" {
		return fmt.Errorf("address is empty")
	}
	return p.call(script, args, addr, nil)
}

func (p *RpcEnvironment) call(script *model.Script, args map[string]string, addr string, cre *credential.Credential) error {
	//TBD:auth,don't hard code agent port
	client, err := rpc.DialHTTP("tcp", addr+":3334")
	if err != nil {
		return fmt.Errorf("connect host failed. %s", err.Error())
	}

	export := ""
	for k, v := range args {
		if k == "address" {
			continue
		}
		if k == "root" {
			v = "nervapp/" //The root dir of app is $HOME/nervapp
		}
		export = export + fmt.Sprintf(" %s=%s", k, v)
	}
	shell := "export " + export + " && cd ~ &&" + script.Content
	fmt.Println(shell)

	remoteScript := &crpc.RemoteScript{Content: shell}
	var reply string
	err = client.Call("Agent.Execute", remoteScript, &reply)
	if err != nil {
		return fmt.Errorf("Agent.Execute failed. %s", err.Error())
	} else {
		fmt.Println(reply)
	}

	return nil
}
