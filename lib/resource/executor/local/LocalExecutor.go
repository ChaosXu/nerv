package local

import (
	"os/exec"
	"fmt"
	"github.com/ChaosXu/nerv/lib/resource/model"
	"github.com/toolkits/file"
)

type Script struct {
	Content string
	Args    map[string]string
}
// LocalExecutor perform an operation when init nerv
type LocalExecutor struct {

}

func (p *LocalExecutor) Perform(class *model.Class, operation string, args map[string]string) error {
	fmt.Println("Perform " + class.Name + "." + operation)
	op := class.GetOperation(operation)
	if op == nil {
		return fmt.Errorf("unsupported operation. class:%s,operation:%s", class.Name, operation)
	}
	script := &Script{}
	script.Args = args
	err := p.loadScript(op.Implementor, script)
	if err != nil {
		return err;
	}
	return p.exec(script)
}

func (p *LocalExecutor) loadScript(path string, script *Script) error {
	if content, err := file.ToString("../../resources/scripts" + path); err != nil {
		return err
	} else {
		script.Content = content
		return nil
	}
}

func (p *LocalExecutor) exec(script *Script) error {
	export := ""
	for k, v := range script.Args {
		export = export + fmt.Sprintf(" %s=%s", k, v)
	}
	export = export + " APP_ROOT=" + script.Args["root"]

	shell := "export " + export + " && " + script.Content
	fmt.Println(shell)

	out, err := exec.Command("/bin/bash", "-c", shell).Output()
	if err != nil {
		return err
	}
	res := string(out)
	fmt.Println(res)

	return nil
}
