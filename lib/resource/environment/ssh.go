package environment

import (
	"fmt"
	"bytes"
	"strings"
	"golang.org/x/crypto/ssh"
	"github.com/ChaosXu/nerv/lib/db"
	"github.com/ChaosXu/nerv/lib/credential"
	"github.com/ChaosXu/nerv/lib/resource/model"
	"github.com/ChaosXu/nerv/lib/resource/repository"
)

// SshEnvironment where nerv add worker cluster.
type SshEnvironment struct {
	ScriptRepository repository.ScriptRepository `inject:"script_rep_standalone"`
}

func (p *SshEnvironment) Exec(class *model.Class, operation *model.Operation, args map[string]string) error {
	fmt.Printf("SSH.Exec %s.%s %s\n", class.Name, operation.Name, operation.Implementor)

	script, err := p.ScriptRepository.Get(operation.Implementor)
	if err != nil {
		return err;
	}

	cres := args["credential"]
	if cres == "" {
		return fmt.Errorf("credential is empty")
	}
	pairs := strings.Split(cres, ",")
	if len(pairs) < 2 {
		return fmt.Errorf("error credential: %s", cres)
	}
	cre := &credential.Credential{}
	if err := db.DB.Where("type=? and name=?", pairs[0], pairs[1]).First(cre).Error; err != nil {
		return fmt.Errorf("credential %s not found", cres, err.Error())
	}

	addr := args["address"]
	if addr == "" {
		return fmt.Errorf("address is empty")
	}
	return p.call(script, args, addr, cre)
}

func (p *SshEnvironment) call(script *model.Script, args map[string]string, addr string, cre *credential.Credential) error {
	config := &ssh.ClientConfig{
		User:cre.User,
		Auth:[]ssh.AuthMethod{
			ssh.Password(cre.Password),
		},
	}
	client, err := ssh.Dial("tcp", addr + ":22", config)
	if err != nil {
		return err
	}

	session, err := client.NewSession()
	if err != nil {
		return nil
	}
	defer session.Close()

	var stdout, stderr bytes.Buffer
	session.Stdout = &stdout
	session.Stderr = &stderr

	export := ""
	for k, v := range args {
		if k == "address" || k == "credential" {
			continue
		}
		export = export + fmt.Sprintf(" %s=%s", k, v)
	}
	shell := "export " + export + " && " + script.Content
	fmt.Println(shell)

	stdoutContent := ""
	stderrContent := ""
	if err := session.Run(shell); err != nil {
		stdoutContent = stdout.String()
		if stdoutContent != "" {
			fmt.Println("stdout\n" + stdoutContent)
		}
		stderrContent = stderr.String()
		if stderrContent != "" {
			fmt.Println("stderr\n" + stderrContent)
			return fmt.Errorf("%s\n%s", err.Error(), stderrContent)
		} else {
			return err
		}
	}
	stdoutContent = stdout.String()
	if stdoutContent != "" {
		fmt.Println(stdoutContent)
	}

	return nil
}
