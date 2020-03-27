package sftp

import (
	"fmt"
	"github.com/google/uuid"
)

//type 0:ISO 1: DiskImage
type FileData struct {
	Type         int
	Name         string
	Path         string
	LocalPath    string
	RemotePath   string
	IsLock       bool
	Authority    int
	ProgressUUID string
	MinMem       int
}

type SSHInfo struct {
	IP      string
	Port    string
	User    string
	KeyPath string
}

func GenerateUUID() string {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	uuid := u.String()
	return uuid
}
