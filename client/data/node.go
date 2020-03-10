package data

import (
	"context"
	"fmt"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
	"time"
)

func NodeAdd(d *pb.NodeData, address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddNode(ctx, d)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("Result: ")
	fmt.Println(r.Status)
	fmt.Printf("Info: ")
	fmt.Println(r.Info)
}
func NodeRemove(d *pb.NodeID, address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RemoveNode(ctx, d)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("Result: ")
	fmt.Println(r.Status)
	fmt.Printf("Info: ")
	fmt.Println(r.Info)
}

func GetNode(a *AuthData, address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.GetNode(ctx, &pb.Base{User: a.Name, Pass: a.Pass1})
	if err != nil {
		log.Fatal(err)
	}
	for {
		article, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: " + strconv.Itoa(int(article.NodeID)) + " IP: " + article.IP + " Path: " + article.Path)
		fmt.Printf(" MaxCPU: " + strconv.Itoa(int(article.Sepc.Maxcpu)) + " MaxMem: " + strconv.Itoa(int(article.Sepc.Maxmem)))
		fmt.Println(" Status: " + strconv.Itoa(int(article.Status)))
	}
}

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
	fmt.Printf("Result: ")
	fmt.Println(r.Status)
	fmt.Printf("Info: ")
	fmt.Println(r.Info)
}
