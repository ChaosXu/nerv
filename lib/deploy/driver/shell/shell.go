package shell

import (
	"net/rpc/jsonrpc"
	"github.com/ChaosXu/nerv/cmd/agent/shell"
	"fmt"
	"github.com/ChaosXu/nerv/lib/env"
	"log"
	"github.com/ChaosXu/nerv/lib/deploy/driver/util"
)

func Execute(address string, scriptUri string, args map[string]string) error {
	rep := env.Config().GetMapString("scripts", "repository")
	if rep == "" {
		return fmt.Errorf("scripts.repository isn't setted")
	}
	scriptUrl := rep + scriptUri
	log.Printf("url:%s\n", scriptUrl)
	script, err := util.LoadScript(scriptUrl)
	if err != nil {
		return err
	}

	port := env.Config().GetMapString("agent", "port", "3334")
	client, err := jsonrpc.Dial("tcp", address + ":" + port)
	if err != nil {
		return err
	}
	defer client.Close()

	remoteScript := &shell.RemoteScript{
		Content: script,
		Args:    args,
	}

	var out string
	err = client.Call("RemoteShell.Execute", remoteScript, &out)
	if err != nil {
		return err
	}
	log.Println(out)
	return nil
}
