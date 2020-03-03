package direct

import (
	"context"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"log"
	"time"
)

func AddGroup(a *AuthData, address, user, pass string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddGroup(ctx, &pb.GroupData{Base: &pb.Base{User: a.Name, Pass: a.Pass1}, User: user})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}
