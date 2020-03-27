package client

import (
	"context"
	"fmt"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

type ImageDataResult struct {
	ID       int    `json:"id"`
	FileName string `json:"filename"`
	Name     string `json:"name"`
	Tag      string `json:"tag"`
	MinMem   int    `json:"minmem"`
}

func GetAllImage(token string) []ImageDataResult {
	conn, err := grpc.Dial(GetgRPCServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := c.GetAllImage(ctx, &pb.Base{Token: token})
	if err != nil {
		log.Println(err)
	}

	var data []ImageDataResult

	for {
		d, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}

		tmp := ImageDataResult{ID: int(d.Id), FileName: d.Filename, Name: d.Name, Tag: d.Tag, MinMem: int(d.Minmem)}
		data = append(data, tmp)
	}
	fmt.Println(data)
	return data
}
