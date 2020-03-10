package data

import (
	"context"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"log"
	"time"
)

func NodeStopVM(d *pb.NodeID, address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.StopNode(ctx, d)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("ID: ")
	log.Println(r.Status)
}
