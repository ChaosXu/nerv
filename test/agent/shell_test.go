package shell_test

import (
	"testing"
	"net/rpc/jsonrpc"
	"log"
	"github.com/ChaosXu/nerv/cmd/agent/shell"
)

func TestExecute(t *testing.T) {
	client, err := jsonrpc.Dial("tcp", "localhost:3334")
	if err != nil {
		t.Fatal(err.Error())
	}

	remoteScript := &shell.RemoteScript{
		Content:"echo hi",
		Args:map[string]string{},
	}

	var out string
	err = client.Call("RemoteShell.Execute", remoteScript, &out)
	if err != nil {
		t.Fatal(err.Error())
	}
	log.Println(out)
}