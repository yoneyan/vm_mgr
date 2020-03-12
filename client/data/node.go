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
	stream, err := c.GetNode(ctx, &pb.Base{User: a.Name, Pass: a.Pass, Token: a.Token})
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
		tmp := []string{strconv.Itoa(int(article.NodeID)), article.IP, article.Path, strconv.Itoa(int(article.Sepc.Maxcpu)),
			strconv.Itoa(int(article.Sepc.Maxmem)), article.Sepc.Net, strconv.FormatBool(article.OnlyAdmin), strconv.Itoa(int(article.Status))}
		data = append(data, tmp)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NodeID", "IP", "Path", "MaxCPU", "MaxMem", "Net", "OnlyAdmin", "Status"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
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
