package etc

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"strconv"
)

func sshTest(addr string, port int, user, pass string) bool {

	config := &ssh.ClientConfig{User: user, HostKeyCallback: ssh.InsecureIgnoreHostKey(), Auth: []ssh.AuthMethod{ssh.Password(pass)}}

	fmt.Println(addr + ":" + strconv.Itoa(port))
	conn, err := ssh.Dial("tcp", addr+":"+strconv.Itoa(port), config)
	//conn, err := ssh.Dial("tcp", addr+":"+string(port), config)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		log.Println(err)
	}
	defer session.Close()

	return true
}

func sshTest1() string {
	ip := "127.0.0.1"
	port := "22"
	user := "user"
	password := "pass"

	config := &ssh.ClientConfig{User: user, HostKeyCallback: ssh.InsecureIgnoreHostKey(), Auth: []ssh.AuthMethod{ssh.Password(password)}}

	conn, err := ssh.Dial("tcp", ip+":"+port, config)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		log.Println(err)
	}
	defer session.Close()

	var output bytes.Buffer
	session.Stdout = &output
	remoteCmd := "ls -la"
	if err := session.Run(remoteCmd); err != nil {
		return "Failed to run: " + err.Error()
	}
	return remoteCmd + ":" + output.String()

}
