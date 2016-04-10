package auth

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"strings"
)

func AuthenticateLocal(user string, pass string) bool {

	cmd := "/bin/hostname"
	hostname := "gargivanu-Inspiron-N5010"
	username := user
	password := pass

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}

	client, err := ssh.Dial("tcp", "localhost:22", config)
	if err == nil {

		session, err := client.NewSession()
		if err != nil {
			panic("Failed to create session: " + err.Error())

		}
		defer session.Close()

		var b bytes.Buffer
		session.Stdout = &b
		if err := session.Run(cmd); err != nil {
			panic("Failed to run: " + err.Error())
		}
		// fmt.Println(b.String())
		if strings.TrimSpace(b.String()) == hostname {
			// fmt.Println("SUCCESS")
			return true
		}

	} else {
		// fmt.Println("FAILED")
		return false
	}
	return false
}

func main() {
	result := AuthenticateLocal("rupesh.thakur", "p@ssw0rdas")
	if result {
		fmt.Println("SUCCESS")
	} else {
		fmt.Println("FAILED")
	}
}
