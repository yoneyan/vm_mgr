package sftp

import (
	"fmt"
	"github.com/pkg/sftp"
	"github.com/yoneyan/vm_mgr/imacon/db"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func DataUpload(filedata *FileData, sshdata *SSHInfo) {

	buf, err := ioutil.ReadFile(sshdata.KeyPath)
	if err != nil {
		log.Println(err)
	}
	key, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		log.Println(err)
	}

	config := &ssh.ClientConfig{
		User:            sshdata.User,
		HostKeyCallback: nil,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
	}
	config.SetDefaults()
	sshConn, err := ssh.Dial("tcp", sshdata.IP+":"+sshdata.Port, config)
	if err != nil {
		log.Println(err)
	}
	defer sshConn.Close()

	client, err := sftp.NewClient(sshConn)
	if err != nil {
		log.Println(err)
	}
	defer client.Close()

	dFile, err := os.Create(filedata.RemotePath)
	if err != nil {
		log.Println(err)
	}
	defer dFile.Close()

	sFile, err := client.Open(filedata.LocalPath)
	if err != nil {
		log.Println(err)
	}

	bytes, err := io.Copy(dFile, sFile)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%d bytes copied\n", bytes)

	fmt.Println(db.RemoveDBTransfer(filedata.ProgressUUID))
}
