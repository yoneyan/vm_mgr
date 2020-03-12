package server

import (
	"context"
	"fmt"
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
	log.Println("Receive Token     : " + in.GetBase().GetToken())

	if data.AdminUserCertification(in.GetUser(), in.GetPass(), in.GetToken()) == false {
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
	log.Println("Receive AuthUser: " + in.GetBase().User + ", AuthPass: " + in.GetBase().Pass)
	log.Println("Receive Token     : " + in.GetBase().GetToken())

	if data.AdminUserCertification(in.GetUser(), in.GetPass(), in.GetToken()) == false {
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

func (s *server) GetUser(d *pb.UserData, stream pb.Grpc_GetUserServer) error {
	log.Println("----GetUser----")
	if d.Mode == 0 {
		log.Println("GetAllUser")
		if data.AdminUserCertification(d.GetUser(), d.GetPass(), d.GetToken()) == false {
			fmt.Println("Administrator certification failed!!!")
			return nil
		}
		result := db.GetDBAllUser()
		for _, a := range result {
			if err := stream.Send(&pb.UserData{
				Id:   int64(a.ID),
				User: a.Name,
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

/*
func (s *server) UserNameChange(ctx context.Context, in *pb.UserData) (*pb.Result, error) {
	log.Println("----UserNameChange----")
	afteruser := in.GetPass()
	log.Println("Receive Before UserName: " + in.GetUser())
	log.Println("Receive After UserName: " + afteruser)
	log.Println("Receive AuthUser: " + in.GetBase().User + ", AuthPass: " + in.GetBase().Pass)

	if data.AdminUserCertification(in.GetBase().User, in.GetBase().Pass) == false {
		return &pb.Result{Status: false, Info: "Authentication failed!!"}, nil
	}
	if data.ExistUserCheck(afteruser) == false {
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
*/

func (s *server) UserPassChange(ctx context.Context, in *pb.UserData) (*pb.Result, error) {
	log.Println("----UserPassChange----")
	log.Println("Receive UserName: " + in.GetUser())
	log.Println("Receive ChangeUserPass: " + in.GetPass())
	log.Println("Receive AuthUser: " + in.GetBase().User + ", AuthPass: " + in.GetBase().Pass)

	if data.UserCertification(in.GetBase().User, in.GetBase().Pass) == false {
		return &pb.Result{Status: false, Info: "Authentication failed!!"}, nil
	}
	id, err := db.GetDBUserID(in.GetUser())
	if err == false {
		return &pb.Result{Status: false, Info: "Error Search UserID!!"}, nil
	}
	if db.ChangeDBUserPassword(id, in.GetPass()) {
		return &pb.Result{Status: true, Info: "OK!"}, nil
	} else {
		return &pb.Result{Status: false, Info: "Change Pass Error!!!"}, nil
	}
}
