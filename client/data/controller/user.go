package direct

import (
	"context"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"log"
	_ "os"
	"time"
)

type AuthData struct {
	Name  string
	Pass1 string
	Pass2 string
}

func AddUser(a *AuthData, address, user, pass string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddUser(ctx, &pb.UserData{Base: &pb.Base{User: a.Name, Pass: a.Pass1}, User: user, Pass: pass})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func RemoveUser(a *AuthData, address, user string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.RemoveUser(ctx, &pb.UserData{Base: &pb.Base{User: a.Name, Pass: a.Pass1}, User: user})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}
