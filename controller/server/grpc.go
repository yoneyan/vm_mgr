package server

// Client gRPC Server

import (
	"fmt"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"net"
)

const port = ":50200"

type server struct {
	pb.UnimplementedGrpcServer
}

func Server() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("failed to listen: %v", err)
	}

	//opts := []grpc.ServerOption{grpc.UnaryInterceptor(authentication)}
	//s := grpc.NewServer(opts...)
	//
	//pb.RegisterGrpcServer(s, &server{})
	//s.Serve(lis)

	s := grpc.NewServer()
	pb.RegisterGrpcServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve: %v", err)
	}
}

//func authentication(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
//	_, ok := metadata.FromIncomingContext(ctx)
//	if !ok {
//		fmt.Println("Not")
//		return nil, status.Error(codes.Unauthenticated, "not found metadata")
//	}
//	//values := md["authorization"]
//	//if len(values) == 0 {
//	//	fmt.Println("Not")
//	//	return nil, status.Error(codes.Unauthenticated, "not found metadata")
//	//}
//	fmt.Println("OK!!!!!!!!!!!!!!")
//	return handler(ctx, req)
//}
