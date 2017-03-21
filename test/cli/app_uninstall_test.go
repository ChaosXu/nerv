package cli

import (
	"github.com/ChaosXu/nerv/test/util"
	"testing"
)

func TestAppUninstall(t *testing.T) {
	id := "6"
	//start
	cmd := &util.Cmd{
		Dir:   "../../release/nerv/nerv-cli/bin",
		Cli:   "./nerv-cli",
		Items: []string{"topo", "uninstall", "-i", id},
	}

	if out, err := cmd.Run(t); err != nil {
		t.Log(string(out))
		t.Errorf("%s", err.Error())
	} else {
		t.Log(string(out))
	}

	//delete
	cmd = &util.Cmd{
		Dir:   "../../release/nerv/nerv-cli/bin",
		Cli:   "./nerv-cli",
		Items: []string{"topo", "delete", "-i", id},
	}

	if out, err := cmd.Run(t); err != nil {
		t.Log(string(out))
		t.Errorf("%s", err.Error())
	} else {
		t.Log(string(out))
	}
}
