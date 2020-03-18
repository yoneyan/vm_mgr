package server

import (
	"context"
	"github.com/yoneyan/vm_mgr/imacon/db"
	"github.com/yoneyan/vm_mgr/imacon/etc"
	"github.com/yoneyan/vm_mgr/imacon/sftp"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"log"
	"strconv"
)

func (s *server) DownloadImage(ctx context.Context, in *pb.ImageTransferData) (*pb.ImageResult, error) {
	log.Println("----DownloadImage----")
	log.Println("Receive GetIP : " + in.GetIp() + ", Port: " + in.GetPort())
	log.Println("Receive Path  : " + in.GetPath())
	log.Println("Receive Type  : " + strconv.Itoa(int(in.Image.GetType())))

	var path, name string
	uuid := sftp.GenerateUUID()

	if in.Image.GetType() == 0 {
		path = etc.ConfigData.ISOPath + "/" + in.Image.GetName()
		name = in.Image.GetName()
	} else if in.Image.GetType() == 1 {
		path = etc.ConfigData.ImagePath + "/" + uuid + ".img"
		name = uuid + ".img"
	}

	uuid = sftp.GenerateUUID()

	db.AddDBTransfer(db.Transfer{UUID: uuid, Status: 0})

	go sftp.DataDownload(&sftp.FileData{
		Name:         name,
		Type:         int(in.Image.GetType()),
		LocalPath:    path,
		RemotePath:   in.GetPath(),
		ProgressUUID: uuid,
		Authority:    int(in.Image.GetAuthority()),
		MinMem:       int(in.Image.GetMinmem()),
	}, &sftp.SSHInfo{
		IP:      in.GetIp(),
		Port:    in.GetPort(),
		User:    "root",
		KeyPath: etc.ConfigData.KeyPath,
	})

	return &pb.ImageResult{Result: true, Uuid: uuid}, nil
}
