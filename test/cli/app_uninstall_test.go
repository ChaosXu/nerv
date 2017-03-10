package cli

import (
	"testing"
)


func TestAppUninstall(t *testing.T) {

	//start
	cmd := &Cmd{
		Dir: "../../release/nerv/nerv-cli/bin",
		Cli:"./nerv-cli",
		Items:[]string{"topo", "uninstall", "-i", "3"},
	}

	if out, err := cmd.Run(t); err != nil {
		t.Log(string(out))
		t.Errorf("%s", err.Error())
	} else {
		t.Log(string(out))
	}

	//delete
	cmd = &Cmd{
		Dir: "../../release/nerv/nerv-cli/bin",
		Cli:"./nerv-cli",
		Items:[]string{"topo", "delete", "-i", "4"},
	}

	if out, err := cmd.Run(t); err != nil {
		t.Log(string(out))
		t.Errorf("%s", err.Error())
	} else {
		t.Log(string(out))
	}
}
