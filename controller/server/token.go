package server

import (
	"context"
	"fmt"
	"github.com/yoneyan/vm_mgr/controller/data"
	"github.com/yoneyan/vm_mgr/controller/db"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"log"
)

func (s *server) GenerateToken(ctx context.Context, in *pb.Base) (*pb.AuthResult, error) {
	log.Println("----GenerateToken----")
	log.Println("Receive AuthUser : " + in.GetUser() + ", AuthPass: " + in.GetPass())
	log.Println("Receive Token    : " + in.GetToken())

	if data.UserCertification(in.GetUser(), in.GetPass()) == false {
		return &pb.AuthResult{Result: false, Token: "Auth Failed!!"}, nil
	}
	uuid, result := data.NewToken(in.GetUser())
	if result {
		return &pb.AuthResult{Result: true, Token: uuid}, nil
	} else {
		return &pb.AuthResult{Result: false, Token: uuid}, nil
	}
}

func (s *server) DeleteToken(ctx context.Context, in *pb.Base) (*pb.Result, error) {
	go data.DeleteExpiredToken()
	log.Println("----DeleteToken----")
	log.Println("Receive AuthUser : " + in.GetUser() + ", AuthPass: " + in.GetPass())
	log.Println("Receive Token    : " + in.GetToken())

	if data.AdminUserCertification(in.GetUser(), in.GetPass(), in.GetToken()) == false {
		return &pb.Result{Status: false, Info: "Authentication failed!!"}, nil
	}
	data, result := db.GetDBToken(in.GetToken())
	if result == false {
		return &pb.Result{Status: false, Info: "DB Search failed!!"}, nil
	}
	info, result := db.RemoveDBToken(data.ID)
	if result {
		return &pb.Result{Status: true, Info: "OK!"}, nil
	} else {
		return &pb.Result{Status: false, Info: info}, nil
	}
}

func (s *server) GetAllToken(d *pb.Base, stream pb.Grpc_GetAllTokenServer) error {
	go data.DeleteExpiredToken()
	log.Println("----GetAllToken----")
	log.Println("Receive AuthUser : " + d.GetUser() + ", AuthPass: " + d.GetPass())
	log.Println("Receive Token    : " + d.GetToken())
	if data.AdminUserCertification(d.GetUser(), d.GetPass(), d.GetToken()) == false {
		fmt.Println("Administrator certification failed!!!")
		return nil
	}
	result := db.GetDBAllToken()
	fmt.Printf("DBstruct: ")
	fmt.Println(result)
	for _, a := range result {
		if err := stream.Send(&pb.TokenData{
			Id:        int64(a.ID),
			Token:     a.Token,
			Userid:    int32(a.Userid),
			Groupid:   a.Groupid,
			Begintime: int64(a.Begintime),
			Endtime:   int64(a.Endtime),
		}); err != nil {
			return err
		}
	}
	return nil
}
