package data

import (
	"context"
	"github.com/olekukonko/tablewriter"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	_ "os"
	"strconv"
	"time"
)

type AuthData struct {
	Name  string
	Pass  string
	Token string
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

	r, err := c.AddUser(ctx, &pb.UserData{Base: &pb.Base{User: a.Name, Pass: a.Pass, Token: a.Token}, User: user, Pass: pass})
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

	r, err := c.RemoveUser(ctx, &pb.UserData{Base: &pb.Base{User: a.Name, Pass: a.Pass, Token: a.Token}, User: user})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func GetAllUser(a *AuthData, address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := c.GetUser(ctx, &pb.UserData{Base: &pb.Base{User: a.Name, Pass: a.Pass, Token: a.Token}, Mode: 0})
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
		tmp := []string{strconv.Itoa(int(article.Id)), article.User}
		data = append(data, tmp)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"UserID", "User"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func UserNameChange(a *AuthData, address, user, pass string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ChangeUserName(ctx, &pb.UserData{Base: &pb.Base{User: a.Name, Pass: a.Pass, Token: a.Token}, User: user, Pass: pass})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func UserPassChange(a *AuthData, address, user, pass string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ChangeUserPass(ctx, &pb.UserData{Base: &pb.Base{User: a.Name, Pass: a.Pass, Token: a.Token}, User: user, Pass: pass})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}
