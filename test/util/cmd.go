package util

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

type Cmd struct {
	Dir   string
	Cli   string
	Items []string
}

func (p *Cmd) Run(t *testing.T) ([]byte, error) {
	if wd, err := os.Getwd(); err != nil {
		t.Error(err)
	} else {
		defer os.Chdir(wd)
	}
	os.Chdir(p.Dir)
	args := p.Cli + " " + strings.Join(p.Items, " ")
	t.Logf("%s\n", args)

	return exec.Command("/bin/bash", "-c", args).Output()
}
