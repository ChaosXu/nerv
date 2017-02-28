package operation

import (
	"fmt"
	"net/rpc"
	"strings"
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/ChaosXu/nerv/lib/credential"
	"github.com/ChaosXu/nerv/lib/resource/model"
	"github.com/ChaosXu/nerv/lib/resource/repository"
	crpc "github.com/ChaosXu/nerv/lib/rpc"
)

// RpcEnvironment where nerv add worker cluster.
type RpcEnvironment struct {
	ScriptRepository repository.ScriptRepository `inject:"script_rep_standalone"`
}

func (p *RpcEnvironment) Exec(class *model.Class, operation *model.Operation, args map[string]string) error {
	fmt.Printf("Rpc.Exec %s.%s %s\n", class.Name, operation.Name, operation.Implementor)

	script, err := p.ScriptRepository.Get(operation.DefineClass, operation.Implementor)
	if err != nil {
		return err;
	}

	cres := args["credential"]
	if cres == "" {
		return fmt.Errorf("credential is empty")
	}
	pairs := strings.Split(cres, ",")
	if len(pairs) < 2 {
		return fmt.Errorf("error credential: %s", cres)
	}
	cre := &credential.Credential{}
	if err := db.DB.Where("type=? and name=?", pairs[0], pairs[1]).First(cre).Error; err != nil {
		return fmt.Errorf("credential %s not found", cres, err.Error())
	}

	addr := args["address"]
	if addr == "" {
		return fmt.Errorf("address is empty")
	}
	return p.call(script, args, addr, cre)
}

func (p *RpcEnvironment) call(script *model.Script, args map[string]string, addr string, cre *credential.Credential) error {
	//TBD:auth,don't hard code agent port
	client, err := rpc.DialHTTP("tcp", addr + ":3334")
	if err != nil {
		return fmt.Errorf("connect host failed. %s", err.Error())
	}

	export := ""
	for k, v := range args {
		if k == "address" || k == "credential" {
			continue
		}
		export = export + fmt.Sprintf(" %s=%s", k, v)
	}
	shell := "export " + export + " && " + script.Content
	fmt.Println(shell)

	remoteScript := &crpc.RemoteScript{Content:shell, Args:map[string]string{}}
	var reply string
	err = client.Call("Agent.Execute", remoteScript, &reply)
	if err != nil {
		return fmt.Errorf("Agent.Execute failed. %s", err.Error())
	} else {
		fmt.Println(reply)
	}

	return nil
}
