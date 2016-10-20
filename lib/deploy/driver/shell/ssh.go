package shell

import (
	"golang.org/x/crypto/ssh"
	"bytes"
	"fmt"
	"log"
	"strings"
	"github.com/ChaosXu/nerv/lib/db"
)

//RemoteExecute a script on the host of addr
func RemoteExecute(addr string, credentialRef string, scriptUrl string, args map[string]string) error {
	script, err := loadScript(scriptUrl)
	if err != nil {
		return err
	}
	log.Println(script)


	credential := Credential{}
	pairs := strings.Split(credentialRef, ",")
	if err := db.DB.Where("type=? and name=?", pairs[0], pairs[1]).First(&credential).Error; err != nil {
		return err
	}

	config := &ssh.ClientConfig{
		User:credential.User,
		Auth:[]ssh.AuthMethod{
			ssh.Password(credential.Password),
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

	export := ""
	for k, v := range args {
		export = export + fmt.Sprintf(" %s=%s", k, v)
	}
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
