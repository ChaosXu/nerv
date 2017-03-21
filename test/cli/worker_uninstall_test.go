package cli

import (
	"github.com/ChaosXu/nerv/test/util"
	"testing"
)

func TestWorkerUninstall(t *testing.T) {

	//start
	cmd := &util.Cmd{
		Dir:   "../../release/nerv/nerv-cli/bin",
		Cli:   "./nerv-cli",
		Items: []string{"topo", "uninstall", "-i", "2"},
	}

	if out, err := cmd.Run(t); err != nil {
		t.Log(string(out))
		t.Errorf("%s", err.Error())
	} else {
		t.Log(string(out))
	}
}
