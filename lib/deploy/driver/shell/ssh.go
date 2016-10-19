package shell

import (
	"golang.org/x/crypto/ssh"
	"bytes"
	"fmt"
	"log"
)

//RemoteExecute a script on the host of addr
func RemoteExecute(addr string, user string, password string, scriptUrl string, args map[string]string) error {
	script, err := loadScript(scriptUrl)
	if err != nil {
		return err
	}

	export := ""
	for k, v := range args {
		export = export + fmt.Sprintf(" %s=%s", k, v)
	}

	config := &ssh.ClientConfig{
		User:user,
		Auth:[]ssh.AuthMethod{
			ssh.Password(password),
		},
	}
	client, err := ssh.Dial("tcp", addr, config)
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
	script = "export " + export + " && " + script
	log.Println(script)
	if err := session.Run(script); err != nil {
		log.Println("stdout\n" + stdout.String())
		log.Println("stderr\n" + stderr.String())
		return err
	} else {
		log.Println("stdout\n" + stdout.String())
		log.Println("stderr\n" + stderr.String())
	}

	return nil
}
