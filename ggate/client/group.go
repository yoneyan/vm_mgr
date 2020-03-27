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

type GroupDataResult struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Admin string `json:"admin"`
	User  string `json:"user"`
}

func GetGroup(token string) []GroupDataResult {
	conn, err := grpc.Dial(GetgRPCServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := c.GetGroup(ctx, &pb.GroupData{Base: &pb.Base{Token: token}})
	if err != nil {
		log.Println(err)
	}

	var data []GroupDataResult

	for {
		d, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}

		tmp := GroupDataResult{ID: int(d.Id), Name: d.Name, Admin: d.Admin, User: d.User}
		data = append(data, tmp)
	}
	fmt.Println(data)

	return data
}
