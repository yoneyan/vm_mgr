package sftp

import (
	"fmt"
	"github.com/pkg/sftp"
	"github.com/yoneyan/vm_mgr/imacon/db"
	"github.com/yoneyan/vm_mgr/imacon/etc"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func DataDownload(filedata *FileData, sshdata *SSHInfo) {
	uuid := GenerateUUID()
	if filedata.Type == 0 {
		filedata.LocalPath = etc.ConfigData.ImagePath + "/" + uuid + ".iso"
		filedata.Name = uuid + ".iso"
	} else if filedata.Type == 1 {
		filedata.LocalPath = etc.ConfigData.ImagePath + "/" + uuid + ".img"
		filedata.Name = uuid + ".img"
	}

	buf, err := ioutil.ReadFile(etc.ConfigData.KeyPath)
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

	dFile, err := client.Create(filedata.RemotePath)
	if err != nil {
		log.Println(err)
	}
	defer dFile.Close()

	sFile, err := os.Open(filedata.LocalPath)
	if err != nil {
		log.Println(err)
	}

	bytes, err := io.Copy(dFile, sFile)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%d bytes copied\n", bytes)

	//fmt.Printf("RemoveDBTransfer: ")
	//fmt.Println(db.RemoveDBTransfer(filedata.ProgressUUID))
	fmt.Printf("AddDBImage: ")
	db.AddDBImage(db.Image{
		FileName:  filedata.Name,
		Name:      "",
		Tag:       "",
		Type:      filedata.Type,
		Capacity:  etc.FileSize(filedata.LocalPath),
		AddTime:   int(time.Now().Unix()),
		Authority: filedata.Authority,
		MinMem:    filedata.MinMem,
		Status:    0,
	})
}
