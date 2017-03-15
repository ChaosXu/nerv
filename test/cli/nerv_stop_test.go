package cli

import (
	"testing"
	"github.com/ChaosXu/nerv/test/util"
)


func TestNervStop(t *testing.T) {

	//start
	cmd := &util.Cmd{
		Dir: "../../release/nerv/nerv-cli/bin",
		Cli:"./nerv-cli",
		Items:[]string{"nerv", "stop", "-i", "1"},
	}

	if out, err := cmd.Run(t); err != nil {
		t.Log(string(out))
		t.Errorf("%s", err.Error())
	} else {
		t.Log(string(out))
	}
}

