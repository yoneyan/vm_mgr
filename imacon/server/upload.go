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

func (s *server) UploadImage(ctx context.Context, in *pb.ImageTransferData) (*pb.ImageResult, error) {
	log.Println("----UploadImage----")
	log.Println("Receive GetIP : " + in.GetIp() + ", Port: " + in.GetPort())
	log.Println("Receive Path  : " + in.GetPath())
	log.Println("Receive ID  : " + strconv.Itoa(int(in.Image.GetType())))

	data, result := db.GetDBImage(in.Image.GetName(), in.Image.GetTag())
	if result == false {
		return &pb.ImageResult{Result: false}, nil
	}

	var path string

	if data.Type == 0 {
		path = etc.ConfigData.ISOPath + "/" + data.FileName
	} else if data.Type == 1 {
		path = etc.ConfigData.ImagePath + "/" + data.FileName
	}

	uuid := sftp.GenerateUUID()

	db.AddDBTransfer(db.Transfer{UUID: uuid, Status: 0})

	go sftp.DataUpload(&sftp.FileData{
		Type:         data.Type,
		LocalPath:    path,
		RemotePath:   in.GetPath(),
		ProgressUUID: uuid,
	}, &sftp.SSHInfo{
		IP:      in.GetIp(),
		Port:    in.GetPort(),
		User:    "root",
		KeyPath: etc.ConfigData.KeyPath,
	})
	return &pb.ImageResult{Result: true, Uuid: uuid}, nil
}
