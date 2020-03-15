package client

import (
	"context"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"log"
	"time"
	//a "github.com/yoneyan/vm_mgr/ggate"
)

func GenerateTokenClient(user, pass string) (string, bool) {
	conn, err := grpc.Dial(GetgRPCServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GenerateToken(ctx, &pb.Base{User: user, Pass: pass})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	if r.Result {
		return r.Token, true
	} else {
		return "", false
	}
}

func CheckTokenClient(token string) bool {
	conn, err := grpc.Dial(GetgRPCServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.CheckToken(ctx, &pb.Base{Token: token})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	if r.Status {
		return true
	} else {
		return false
	}
}

func DeleteTokenClient(token string) bool {
	conn, err := grpc.Dial(GetgRPCServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.DeleteToken(ctx, &pb.Base{Token: token})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	if r.Status {
		return true
	} else {
		return false
	}
}
