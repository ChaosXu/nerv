package cli

import (
	"github.com/ChaosXu/nerv/test/util"
	"testing"
)

func TestWorkerCredentialDelete(t *testing.T) {

	cmd := &util.Cmd{
		Dir:   "../../release/nerv/nerv-cli/bin",
		Cli:   "./nerv-cli",
		Items: []string{"credential", "delete", "-i", "1"},
	}

	if out, err := cmd.Run(t); err != nil {
		t.Log(string(out))
		t.Errorf("%s", err.Error())
	} else {
		t.Log(string(out))
	}
}
