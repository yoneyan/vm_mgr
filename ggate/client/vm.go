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

type VMData struct {
	NodeID    int    `json:"nodeid"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CPU       int    `json:"cpu"`
	Mem       int    `json:"mem"`
	Net       string `json:"net"`
	AutoStart bool   `json:"autostart"`
	Status    int    `json:"status"`
}

func GetUserVMClient(token string) []VMData {

	conn, err := grpc.Dial(GetgRPCServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := c.GetUserVM(ctx, &pb.Base{Token: token})
	if err != nil {
		log.Fatal(err)
	}

	var data []VMData

	for {
		article, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		tmp := VMData{int(article.Node), int(article.Option.Id), article.Vmname,
			int(article.Vcpu), int(article.Vmem), article.Vnet,
			article.Option.Autostart, int(article.Option.Status)}
		data = append(data, tmp)
	}
	fmt.Println(data)
	return data
}
