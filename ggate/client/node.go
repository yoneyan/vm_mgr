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

type NodeDataStruct struct {
	ID       int    `json:"id"`
	HostName string `json:"hostname"`
	IsAdmin  bool   `json:"isadmin"`
}

func GetNode(token string) []NodeDataStruct {
	conn, err := grpc.Dial(GetgRPCServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := c.GetNode(ctx, &pb.Base{Token: token})
	if err != nil {
		log.Println(err)
	}

	var data []NodeDataStruct

	for {
		d, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		tmp := NodeDataStruct{ID: int(d.NodeID), HostName: d.Hostname, IsAdmin: d.OnlyAdmin}
		data = append(data, tmp)
	}
	fmt.Println(data)
	return data
}
