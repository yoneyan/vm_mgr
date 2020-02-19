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
	pb.UnimplementedGreaterServer
}

func (s *server) CreateVM(ctx context.Context, in *pb.VMData) (*pb.Result, error) {
	log.Printf("Receive ID: %v", in.GetId())
	log.Printf("Receive Storage: %v", in.GetStorage())
	log.Printf("Receive cpu: %v", in.GetVcpu())
	log.Printf("Receive mem: %v", in.GetVmem())
	log.Printf("Receive name: %v", in.GetVmname())
	log.Printf("Receive vnc: %v", in.GetVnc())
	log.Printf("Receive net: %v", in.GetVnet())
	return &pb.Result{Status: false}, nil
}

func web() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreaterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
