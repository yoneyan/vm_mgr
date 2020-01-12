package main

import (
	"bytes"
	"golang.org/x/crypto/ssh"
	"log"
)

func sshTest() string{
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
		//log.Fatal("Failed to run: " + err.Error())
		return "Failed to run: " + err.Error()
	}
	//log.Println(remoteCmd + ":" + output.String())
	return remoteCmd + ":" + output.String()

}
