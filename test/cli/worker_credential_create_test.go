package cli

import (
	"testing"
	"github.com/ChaosXu/nerv/test/util"
)


func TestWorkerCredentialCreate(t *testing.T) {

	cmd := &util.Cmd{
		Dir: "../../release/nerv/nerv-cli/bin",
		Cli:"./nerv-cli",
		Items:[]string{"credential", "create", "-D", "../../../../test/cli/worker_credential.json"},
	}

	if out, err := cmd.Run(t); err != nil {
		t.Log(string(out))
		t.Errorf("%s", err.Error())
	} else {
		res := string(out)
		t.Log(res)
	}
}
