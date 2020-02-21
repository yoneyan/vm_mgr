package main

import (
	"context"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":50100"

type server struct {
	pb.UnimplementedVMServer
}

func (s *server) CreateVM(ctx context.Context, in *pb.VMData) (*pb.Result, error) {
	log.Println("----CreateVM----")
	log.Printf("Receive ID: %v", in.GetId())
	log.Printf("Receive Storage: %v", in.GetStorage())
	log.Printf("Receive cpu: %v", in.GetVcpu())
	log.Printf("Receive mem: %v", in.GetVmem())
	log.Printf("Receive name: %v", in.GetVmname())
	log.Printf("Receive vnc: %v", in.GetVnc())
	log.Printf("Receive net: %v", in.GetVnet())
	return &pb.Result{Status: false}, nil
}

func (s *server) DeleteVM(ctx context.Context, in *pb.VMID) (*pb.Result, error) {
	log.Println("----DeleteVM----")
	log.Printf("Receive ID: %v", in.GetId())
	return &pb.Result{Status: false}, nil
}

func grpc_test() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterVMServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
