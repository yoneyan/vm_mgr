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

func AddImage(d *pb.ImageData, address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddImage(ctx, d)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("Result: ")
	fmt.Println(r.Result)
	fmt.Printf("Info: ")
	fmt.Println(r.Info)
}

func DeleteImage(d *pb.ImageData, address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.DeleteImage(ctx, d)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("Result: ")
	fmt.Println(r.Result)
	fmt.Printf("Info: ")
	fmt.Println(r.Info)
}

func GetAllImage(d *pb.Base, address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.GetAllImage(ctx, d)
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
		tmp := []string{strconv.Itoa(int(article.Imaconid)), strconv.Itoa(int(article.Id)),
			strconv.Itoa(int(article.Type)), article.Filename, article.Name, article.Tag,
			strconv.Itoa(int(article.Minmem)), strconv.Itoa(int(article.Authority))}
		data = append(data, tmp)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ImaconID", "ID", "ID", "FileName", "Name", "Tag", "minmem", "authority"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func ImageTagChange(d *pb.ImageData, address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ChangeTagImage(ctx, d)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("Result: ")
	fmt.Println(r.Result)
	fmt.Printf("Info: ")
	fmt.Println(r.Info)
}

func ImageNameChange(d *pb.ImageData, address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ChangeNameImage(ctx, d)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("Result: ")
	fmt.Println(r.Result)
	fmt.Printf("Info: ")
	fmt.Println(r.Info)
}
