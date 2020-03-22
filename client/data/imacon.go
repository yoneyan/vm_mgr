package data

import (
	"context"
	"fmt"
	"github.com/olekukonko/tablewriter"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func AddImacon(d *pb.ImaconData, address string) {
	fmt.Println(address)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddImacon(ctx, d)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("Result: ")
	fmt.Println(r.Status)
	fmt.Printf("Info: ")
	fmt.Println(r.Info)
}

func RemoveImacon(d *pb.ImaconData, address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.RemoveImacon(ctx, d)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("Result: ")
	fmt.Println(r.Status)
	fmt.Printf("Info: ")
	fmt.Println(r.Info)
}

func GetAllImacon(a *pb.Base, address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.GetNode(ctx, a)
	if err != nil {
		log.Fatal(err)
	}
	var data [][]string
	for {
		article, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		tmp := []string{strconv.Itoa(int(article.NodeID)), article.IP, article.Hostname, strconv.Itoa(int(article.Status))}
		data = append(data, tmp)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ImaconID", "IP", "HostName", "Status"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
