package client

import (
	"context"
	"fmt"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"log"
	"time"
)

func GenerateTokenClient(user, pass string) *AuthResult {
	fmt.Println(GetgRPCServerAddress())
	conn, err := grpc.Dial(GetgRPCServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GenerateToken(ctx, &pb.Base{User: user, Pass: pass})
	if err != nil {
		log.Println("could not greet: %v", err)
	}
	return &AuthResult{Result: r.Result, Token: r.Token, UserName: r.Name, UserID: int(r.Id)}
}

func CheckTokenClient(token string) Result {
	conn, err := grpc.Dial(GetgRPCServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.CheckToken(ctx, &pb.Base{Token: token})
	if err != nil {
		log.Println("could not greet: %v", err)
	}

	return Result{
		Result: r.Status,
		Info:   r.Info,
	}
}

func DeleteTokenClient(token string) Result {
	conn, err := grpc.Dial(GetgRPCServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.DeleteToken(ctx, &pb.Base{Token: token})
	if err != nil {
		log.Println("could not greet: %v", err)
	}
	return Result{
		Result: r.Status,
		Info:   r.Info,
	}
}

/*
func GetAllTokenClient(token string) bool {
	conn, err := grpc.Dial(GetgRPCServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//r, err := c.GetAllToken(ctx, &pb.Base{Token: token})
	//if err != nil {
	//	log.Println("could not greet: %v", err)
	//}
	//if r.Status {
	//	return true
	//} else {
	//	return false
	//}
}
*/
