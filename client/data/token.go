package data

import (
	"context"
	"fmt"
	"github.com/olekukonko/tablewriter"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"os"
	_ "os"
	"strconv"
	"time"
)

func GenerateToken(address, user, pass string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GenerateToken(ctx, &pb.Base{User: user, Pass: pass})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	if r.Result {
		log.Println("ok")
		log.Printf("User: " + user + "| Token: ")
		log.Println(r.Token)
	} else {
		log.Println("error")
		log.Printf("Info: ")
		log.Println(r.Token)
	}

}

func DeleteToken(a *AuthData, address, token string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.DeleteToken(ctx, &pb.Base{User: a.Name, Pass: a.Pass, Token: token})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func GetAllToken(a *AuthData, address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	md := metadata.New(map[string]string{"authuser": a.Name, "authpass": a.Pass, "authtoken": a.Token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	//var header metadata.MD

	stream, err := c.GetAllToken(ctx, &pb.Base{User: a.Name, Pass: a.Pass, Token: a.Token}, grpc.Header(&md))
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
		begintime := time.Unix(article.Begintime, 0).Format("2006/01/02 15:04:05")
		endtime := time.Unix(article.Endtime, 0).Format("2006/01/02 15:04:05")

		tmp := []string{strconv.Itoa(int(article.Id)), article.Token,
			strconv.Itoa(int(article.Userid)), article.Groupid,
			begintime, endtime}
		data = append(data, tmp)
	}
	fmt.Println(time.Now().Format("2006/01/02 15:04:05"))
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Token", "UserID", "GroupID", "BeginTime", "EndTime"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
