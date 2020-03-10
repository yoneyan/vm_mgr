package server

import (
	"context"
	"fmt"
	"github.com/yoneyan/vm_mgr/controller/data"
	"github.com/yoneyan/vm_mgr/controller/db"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"log"
	"strconv"
)

func (s *server) AddNode(ctx context.Context, in *pb.NodeData) (*pb.Result, error) {
	log.Println("----AddNode----")
	log.Println("Receive NodeID    : " + strconv.Itoa(int(in.GetNodeID())))
	log.Println("Receive HostName  : " + in.GetHostname())
	log.Println("Receive IP        : " + in.GetIP())
	log.Println("Receive OnlyAdmin : " + strconv.FormatBool(in.GetOnlyAdmin()))
	log.Println("Receive Storage   : " + in.GetPath())
	log.Printf("Receive Spec      : ")
	log.Println(in.GetSepc())
	log.Println("Receive AuthUser  : " + in.GetBase().User + ", AuthPass: " + in.GetBase().Pass)

	if data.AdminUserCertification(in.GetBase().User, in.GetBase().Pass) == false {
		return &pb.Result{Status: false, Info: "Authentication failed!!"}, nil
	}
	info, result := data.ExistNodeCheck(in.GetHostname(), in.GetIP())
	if result {
		return &pb.Result{Status: false, Info: info}, nil
	}
	var admin int
	if in.GetOnlyAdmin() {
		admin = 0
	} else {
		admin = 1
	}
	if db.AddDBNode(db.Node{
		ID:        int(in.GetNodeID()),
		HostName:  in.GetHostname(),
		IP:        in.GetIP(),
		Path:      in.GetPath(),
		OnlyAdmin: admin,
		MaxCPU:    int(in.GetSepc().GetMaxcpu()),
		MaxMem:    int(in.GetSepc().GetMaxmem()),
	}) {
		return &pb.Result{Status: true, Info: "OK!"}, nil
	} else {
		return &pb.Result{Status: false, Info: "DB Error!!"}, nil
	}
}

func (s *server) RemoveNode(ctx context.Context, in *pb.NodeID) (*pb.Result, error) {
	log.Println("----RemoveNode----")
	log.Println("Receive ID       : " + strconv.Itoa(int(in.GetNodeID())))
	log.Println("Receive AuthUser : " + in.GetBase().User + ", AuthPass: " + in.GetBase().Pass)

	if data.AdminUserCertification(in.GetBase().User, in.GetBase().Pass) == false {
		return &pb.Result{Status: false, Info: "Authentication failed!!"}, nil
	}
	if db.RemoveDBNode(int(in.GetNodeID())) {
		return &pb.Result{Status: true, Info: "OK!"}, nil
	} else {
		return &pb.Result{Status: false, Info: "DB Error!!"}, nil
	}
}

func (s *server) GetNode(d *pb.Base, stream pb.Grpc_GetNodeServer) error {
	log.Println("----GetNode----")
	log.Println("Receive AuthUser : " + d.GetUser() + ", AuthPass: " + d.GetPass())
	if data.AdminUserCertification(d.GetUser(), d.GetPass()) == false {
		fmt.Println("Administrator certification failed!!!")
		return nil
	}
	result := db.GetDBAllNode()
	fmt.Println(result)
	var OnlyAdmin bool
	for _, a := range result {
		if a.OnlyAdmin == 0 {
			OnlyAdmin = true
		} else {
			OnlyAdmin = false
		}
		if err := stream.Send(&pb.NodeData{
			NodeID:    int32(a.ID),
			Hostname:  a.HostName,
			IP:        a.IP,
			Path:      a.Path,
			OnlyAdmin: OnlyAdmin,
			Status:    int32(a.Status),
			Sepc: &pb.SpecData{
				Maxcpu: int32(a.MaxCPU),
				Maxmem: int32(a.MaxMem),
			},
		}); err != nil {
			return err
		}
	}
	return nil
}
