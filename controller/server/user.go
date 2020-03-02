package server

import (
	"context"
	"github.com/yoneyan/vm_mgr/controller/data"
	"github.com/yoneyan/vm_mgr/controller/db"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"log"
)

func (s *server) AddUser(ctx context.Context, in *pb.UserData) (*pb.Result, error) {
	log.Println("----AddUser----")
	log.Println("Receive UserName: " + in.GetUser())
	log.Println("Receive Pass: " + in.GetPass())
	log.Println("Receive AuthUser: " + in.GetBase().User + ", AuthPass: " + in.GetBase().Pass)

	if data.AdminUserCertification(in.GetBase().User, in.GetBase().Pass) == false {
		return &pb.Result{Status: false, Info: "Authentication failed!!"}, nil
	}
	if data.ExistUserCheck(in.GetUser()) {
		return &pb.Result{Status: false, Info: "Exists User!!"}, nil
	}
	if data.GroupAllUserCheck(in.GetUser()) {
		return &pb.Result{Status: false, Info: "Exists GroupUser!!"}, nil
	}
	db.AddDBUser(db.User{Name: in.GetUser(), Pass: in.GetPass()})
	{
		return &pb.Result{Status: true, Info: "OK!"}, nil
	}
	return &pb.Result{Status: false, Info: "DB Error!!"}, nil
}

func (s *server) RemoveUser(ctx context.Context, in *pb.UserData) (*pb.Result, error) {
	log.Println("----RemoveUser----")
	log.Println("Receive UserName: " + in.GetUser())

	if data.AdminUserCertification(in.GetBase().User, in.GetBase().Pass) == false {
		return &pb.Result{Status: false, Info: "Authentication failed!!"}, nil
	}
	if data.ExistUserCheck(in.GetUser()) == false {
		return &pb.Result{Status: false, Info: "Not exists User!!"}, nil
	}
	if data.GroupAllUserCheck(in.GetUser()) == false {
		return &pb.Result{Status: false, Info: "Exists GroupUser!!"}, nil
	}
	db.RemoveDBUser(in.GetUser())
	{
		return &pb.Result{Status: true, Info: "OK!"}, nil
	}
	return &pb.Result{Status: false, Info: "DB Error!!"}, nil
}
