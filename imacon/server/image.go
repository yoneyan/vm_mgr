package server

import (
	"context"
	"fmt"
	"github.com/yoneyan/vm_mgr/imacon/db"
	"github.com/yoneyan/vm_mgr/imacon/etc"
	"github.com/yoneyan/vm_mgr/imacon/sftp"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"log"
	"strconv"
	"time"
)

func (s *server) AddImage(ctx context.Context, in *pb.ImageData) (*pb.ImageResult, error) {
	log.Println("----AddImage----")
	log.Println("Receive FileName : " + in.GetFilename())
	log.Println("Receive ID     : " + strconv.Itoa(int(in.GetType())))
	log.Println("Receive Name     : " + in.GetName())

	_, result := db.GetDBImageFileName(in.GetFilename())
	if result == false {
		return &pb.ImageResult{Result: false, Info: "DB Error"}, nil
	}
	path := etc.GeneratePath(int(in.GetType()), in.GetFilename())
	if etc.FileExists(path) == false {
		return &pb.ImageResult{Result: false, Info: "File not exist..."}, nil
	}

	if in.GetType() == 0 || in.GetType() == 1 {
		sftp.DataDownload(&sftp.FileData{
			Type:       0,
			Name:       "",
			Path:       in.GetPath(),
			RemotePath: in.GetPath(),
			Authority:  int(in.GetAuthority()),
			MinMem:     int(in.GetMinmem()),
		}, &sftp.SSHInfo{IP: in.GetRaddr(), Port: "22", User: "root"})
	}

	if in.GetType() == 10 || in.GetType() == 11 {
		if etc.FileExists(etc.GeneratePath(int(in.GetType()%10), in.GetFilename())) == false {
			return &pb.ImageResult{Result: false, Info: "File Not Exists!!"}, nil
		}
		if db.AddDBImage(db.Image{
			FileName:  in.GetFilename(),
			Name:      in.GetName(),
			Tag:       in.GetTag(),
			Type:      int(in.GetType() % 10),
			Capacity:  etc.FileSize(path),
			AddTime:   int(time.Now().Unix()),
			Authority: int(in.GetAuthority()),
			MinMem:    int(in.GetMinmem()),
			Status:    0,
		}) == false {
			return &pb.ImageResult{Result: false, Info: "DB Error: Change Error!!"}, nil
		}
	}
	return &pb.ImageResult{Result: true, Info: "Request acceptance!!"}, nil
}

func (s *server) DeleteImage(ctx context.Context, in *pb.ImageData) (*pb.ImageResult, error) {
	log.Println("----DeleteImage----")
	log.Println("Receive FileName : " + in.GetFilename())
	log.Println("Receive ID     : " + strconv.Itoa(int(in.GetType())))

	d, result := db.GetDBImageFileName(in.GetFilename())
	if result == false {
		return &pb.ImageResult{Result: false, Info: "DB Error"}, nil
	}
	if d.Type != int(in.GetType()) {
		return &pb.ImageResult{Result: false, Info: "image type is wrong..."}, nil
	}
	if db.RemoveDBImage(d.ID) == false {
		return &pb.ImageResult{Result: false, Info: "DB Error: Change Error!!"}, nil
	}

	return &pb.ImageResult{Result: true, Info: "ok"}, nil
}

func (s *server) ChangeNameImage(ctx context.Context, in *pb.ImageData) (*pb.ImageResult, error) {
	log.Println("----ChangeNameImage----")
	log.Println("Receive FileName : " + in.GetFilename())
	log.Println("Receive Name     : " + in.GetName())

	d, result := db.GetDBImageFileName(in.GetFilename())
	if result == false {
		return &pb.ImageResult{Result: false, Info: "DB Error"}, nil
	}
	if db.ChangeDBImageName(d.ID, in.GetName()) == false {
		return &pb.ImageResult{Result: false, Info: "DB Error: Change Error!!"}, nil
	}

	return &pb.ImageResult{Result: true, Info: "ok"}, nil
}

func (s *server) ChangeTagImage(ctx context.Context, in *pb.ImageData) (*pb.ImageResult, error) {
	log.Println("----ChangeTagImage----")
	log.Println("Receive FileName : " + in.GetFilename())
	log.Println("Receive ID     : " + strconv.Itoa(int(in.GetType())))
	log.Println("Receive Tag      : " + in.GetTag())

	d, result := db.GetDBImageFileName(in.GetFilename())
	if result == false {
		return &pb.ImageResult{Result: false, Info: "DB Error"}, nil
	}
	if db.ChangeDBImageTag(d.ID, in.GetTag()) == false {
		return &pb.ImageResult{Result: false, Info: "DB Error: Change Error!!"}, nil
	}

	return &pb.ImageResult{Result: true, Info: "ok"}, nil
}

func (s *server) ExistImage(ctx context.Context, in *pb.ImageData) (*pb.ImageResult, error) {
	log.Println("----ExistImage----")
	log.Println("Receive ID : " + strconv.Itoa(int(in.GetType())))
	log.Println("Receive Name : " + in.GetName())
	log.Println("Receive Tag  : " + in.GetTag())

	var path string

	r, result := db.GetDBImage(in.GetName(), in.GetTag())
	if result == false {
		return &pb.ImageResult{Result: false, Info: "no found"}, nil
	}
	if r.Type == 0 {
		path = "iso/" + r.FileName
	} else if r.Type == 1 {
		path = "image/" + r.FileName
	}
	return &pb.ImageResult{Result: true, Info: "ok", Path: path}, nil
}
func (s *server) ProgressImage(ctx context.Context, in *pb.ImageData) (*pb.ImageResult, error) {
	log.Println("----ProgressImage----")
	log.Println("Receive UUID : " + in.GetUuid())

	_, result := db.GetDBTransfer(in.GetUuid())
	if result == false {
		return &pb.ImageResult{Result: false, Info: "Not Found!!"}, nil
	}

	return &pb.ImageResult{Result: true, Info: "ok"}, nil
}

func (s *server) GetAllImage(d *pb.Base, stream pb.Grpc_GetAllImageServer) error {
	log.Println("----GetAllImage----")

	result := db.GetAllDBImage()
	fmt.Printf("DBstruct: ")
	fmt.Println(result)
	for _, a := range result {

		if err := stream.Send(&pb.ImageData{
			Id:        int32(a.ID),
			Path:      etc.GeneratePath(a.Type, a.FileName),
			Name:      a.Name,
			Tag:       a.Tag,
			Type:      int32(a.Type),
			Capacity:  int64(a.Capacity),
			Addtime:   int64(a.AddTime),
			Minmem:    int32(a.MinMem),
			Authority: int32(a.Authority),
			Filename:  a.FileName,
		}); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) GetAllProgressImage(d *pb.Base, stream pb.Grpc_GetAllProgressImageServer) error {
	log.Println("----GetAllProgressImage----")

	result := db.GetAllDBTransfer()
	fmt.Printf("DBstruct: ")
	fmt.Println(result)
	for _, a := range result {

		if err := stream.Send(&pb.ImageResult{Uuid: a.UUID}); err != nil {
			return err
		}
	}
	return nil
}
