package client

// Client gRPC Server

import (
	"context"
	"fmt"
	"github.com/yoneyan/vm_mgr/controller/data"
	"github.com/yoneyan/vm_mgr/controller/db"
	"github.com/yoneyan/vm_mgr/controller/node"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go/client"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

const port = ":50100"

type server struct {
	pb.UnimplementedVMServer
}

func (s *server) AddNode(ctx context.Context, in *pb.NodeData) (*pb.Result, error) {
	fmt.Println("----------StartVM-----")
	log.Printf("Receive NodeID    : %v", in.GetNodeID())
	log.Printf("Receive NodeHost  : %v", in.GetHostname())
	log.Printf("Receive NodeIP    : %v", in.GetIP())
	log.Printf("Receive NodePort  : %v", in.GetPort())
	log.Printf("Receive NodeEnable: %v", in.GetEnable())
	log.Printf("Receive NodeAuth  : %v", in.GetAuth())
	log.Printf("Receive NodeSpec  : %v", in.GetSepc())

	if data.AdminUserCertification(in.GetBase().User, in.GetBase().Pass, in.GetBase().Group) {
		if db.AddDBNode(db.Node{
			ID:       int(in.GetNodeID()),
			HostName: in.GetHostname(),
			IP:       in.GetIP(),
			Port:     int(in.GetPort()),
			Auth:     int(in.GetAuth()),
			MaxCPU:   int(in.GetSepc().Maxcpu),
			MaxMem:   int(in.GetSepc().Maxmem),
		}) {
			return &pb.Result{Status: true}, nil
		}
	}
	return &pb.Result{Status: false}, nil

}

func (s *server) DeleteNode(ctx context.Context, in *pb.NodeID) (*pb.Result, error) {
	fmt.Println("----------StopNode-----")
	log.Printf("Receive NodeID: %v", in.GetNodeID())
	if data.AdminUserCertification(in.GetBase().User, in.GetBase().Pass, in.GetBase().Group) {
		db.NodeDBStatusUpdate(int(in.GetNodeID()), 0)
		if db.DeleteDBNode(int(in.GetNodeID())) {
			d, r := db.GetDBNodeID(int(in.GetNodeID()))
			if r {
				node.NodeAllStop(data.GenerateAddress(d.IP, d.Port))
				fmt.Println("NodeID: " + strconv.Itoa(int(in.GetNodeID())) + "Stop send!! ")
				if db.DeleteDBNode(int(in.GetNodeID())) {
					return &pb.Result{Status: true}, nil
				}
			}
		}
	}
	return &pb.Result{Status: false}, nil
}

func (s *server) StartNode(ctx context.Context, in *pb.NodeID) (*pb.Result, error) {
	log.Println("----StartNode----")
	log.Printf("Receive NodeID: %v", in.GetNodeID())
	if data.AdminUserCertification(in.GetBase().User, in.GetBase().Pass, in.GetBase().Group) {
		if db.NodeDBStatusUpdate(int(in.GetNodeID()), 1) {
			return &pb.Result{Status: true}, nil
		}
	}
	return &pb.Result{Status: false}, nil
}

func (s *server) StopNode(ctx context.Context, in *pb.NodeID) (*pb.Result, error) {
	log.Println("----StopNode----")
	log.Printf("Receive NodeID: %v", in.GetNodeID())
	d, r := db.GetDBNodeID(int(in.GetNodeID()))
	if r {
		db.NodeDBStatusUpdate(int(in.GetNodeID()), 0)
		node.NodeAllStop(data.GenerateAddress(d.IP, d.Port))
		fmt.Println("NodeID: " + strconv.Itoa(int(in.GetNodeID())) + "Stop send!! ")
		return &pb.Result{Status: true}, nil
	} else {
		return &pb.Result{Status: false}, nil
	}
}

func ProcessClient() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterVMServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve: %v", err)
	}
}
