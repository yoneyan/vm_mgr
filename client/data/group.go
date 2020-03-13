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

func AddGroup(a *AuthData, address, group, net, maxvm, maxcpu, maxmem, maxstorage string) {
	vm, err := strconv.Atoi(maxvm)
	if err != nil {
		log.Fatal("MaxCPU Error")
	}
	cpu, err := strconv.Atoi(maxcpu)
	if err != nil {
		log.Fatal("MaxCPU Error")
	}
	mem, err := strconv.Atoi(maxmem)
	if err != nil {
		log.Fatal("MaxMem Error")
	}
	storage, err := strconv.Atoi(maxstorage)
	if err != nil {
		log.Fatal("MaxStorage Error")
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddGroup(ctx, &pb.GroupData{Base: &pb.Base{User: a.Name, Pass: a.Pass, Token: a.Token}, Name: group, Sepc: &pb.SpecData{Maxvm: int32(vm), Net: net, Maxcpu: int32(cpu), Maxmem: int32(mem), Maxstorage: int32(storage)}})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func RemoveGroup(a *AuthData, address, group string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.RemoveGroup(ctx, &pb.GroupData{Base: &pb.Base{User: a.Name, Pass: a.Pass, Token: a.Token}, Name: group})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func GetAllGroup(a *AuthData, address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.GetGroup(ctx, &pb.GroupData{Base: &pb.Base{User: a.Name, Pass: a.Pass, Token: a.Token}, Mode: 0})
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
		tmp := []string{strconv.Itoa(int(article.Id)), article.Name, article.Admin, article.User,
			strconv.Itoa(int(article.Sepc.Maxcpu)), strconv.Itoa(int(article.Sepc.Maxmem)), strconv.Itoa(int(article.Sepc.Maxstorage)),
			article.Sepc.Net}
		data = append(data, tmp)

	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Admin", "User", "MaxCPU", "MaxMem", "MaxStorage", "Net"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func GetSelectGroup(a *AuthData, address, name string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.GetGroup(ctx, &pb.GroupData{Base: &pb.Base{User: a.Name, Pass: a.Pass, Token: a.Token}, Name: name, Mode: 1})
	if err != nil {
		log.Fatal(err)
	}
	for {
		article, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: " + strconv.Itoa(int(article.Id)) + " Name: " + article.Name + " Admin: " + article.Admin + " User: " + article.User)
		fmt.Println(" MaxCPU: " + strconv.Itoa(int(article.Sepc.Maxcpu)) + " MaxMem: " + strconv.Itoa(int(article.Sepc.Maxmem)) + " MaxStorage: " + strconv.Itoa(int(article.Sepc.Maxstorage)))
	}
}

func GetMyGroup(a *AuthData, address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.GetGroup(ctx, &pb.GroupData{Base: &pb.Base{User: a.Name, Pass: a.Pass, Token: a.Token}, Mode: 2})
	if err != nil {
		log.Fatal(err)
	}
	for {
		article, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: " + strconv.Itoa(int(article.Id)) + " Name: " + article.Name + " Admin: " + article.Admin + " User: " + article.User)
		fmt.Println(" MaxCPU: " + strconv.Itoa(int(article.Sepc.Maxcpu)) + " MaxMem: " + strconv.Itoa(int(article.Sepc.Maxmem)) + " MaxStorage: " + strconv.Itoa(int(article.Sepc.Maxstorage)))
	}
}

func JoinAddGroup(a *AuthData, address, genre, group, user string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if genre == "Admin" || genre == "admin" {
		r, err := c.UserAddGroup(ctx, &pb.GroupData{Base: &pb.Base{User: a.Name, Pass: a.Pass, Token: a.Token}, Name: group, User: user, Mode: 0})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Status: ")
		log.Println(r.Status)

		log.Printf("Info: ")
		log.Println(r.Info)
	} else if genre == "User" || genre == "user" {
		r, err := c.UserAddGroup(ctx, &pb.GroupData{Base: &pb.Base{User: a.Name, Pass: a.Pass, Token: a.Token}, Name: group, User: user, Mode: 1})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Status: ")
		log.Println(r.Status)

		log.Printf("Info: ")
		log.Println(r.Info)
	}
}

func JoinRemoveGroup(a *AuthData, address, genre, group, user string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if genre == "Admin" || genre == "admin" {
		r, err := c.UserRemoveGroup(ctx, &pb.GroupData{Base: &pb.Base{User: a.Name, Pass: a.Pass, Token: a.Token}, Name: group, Admin: user, Mode: 0})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Status: ")
		log.Println(r.Status)

		log.Printf("Info: ")
		log.Println(r.Info)
	} else if genre == "User" || genre == "user" {
		r, err := c.UserRemoveGroup(ctx, &pb.GroupData{Base: &pb.Base{User: a.Name, Pass: a.Pass, Token: a.Token}, Name: group, User: user, Mode: 1})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Status: ")
		log.Println(r.Status)

		log.Printf("Info: ")
		log.Println(r.Info)
	}
}
