package client

import (
	"context"
	"fmt"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
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

func StartVMClient(token, vmid string) Result {
	id, err := strconv.Atoi(vmid)
	if err != nil {
		log.Println("id error!!")
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.StartVM(ctx, &pb.VMID{Base: &pb.Base{Token: token}, Id: int64(id)})
	if err != nil {
		log.Println("could not greet: %v", err)
	}
	return Result{
		Result: r.GetStatus(),
		Info:   r.GetInfo(),
	}
}

func ResetVMClient(token, vmid string) Result {
	id, err := strconv.Atoi(vmid)
	if err != nil {
		log.Println("id error!!")
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ResetVM(ctx, &pb.VMID{Base: &pb.Base{Token: token}, Id: int64(id)})
	if err != nil {
		log.Println("could not greet: %v", err)
	}
	return Result{
		Result: r.GetStatus(),
		Info:   r.GetInfo(),
	}
}

func StopVMClient(token, vmid string) Result {
	id, err := strconv.Atoi(vmid)
	if err != nil {
		log.Println("id error!!")
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.StopVM(ctx, &pb.VMID{Base: &pb.Base{Token: token}, Id: int64(id)})
	if err != nil {
		log.Println("could not greet: %v", err)
	}
	return Result{
		Result: r.GetStatus(),
		Info:   r.GetInfo(),
	}
}

func ShutdownVMClient(token, vmid string) Result {
	id, err := strconv.Atoi(vmid)
	if err != nil {
		log.Println("id error!!")
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ShutdownVM(ctx, &pb.VMID{Base: &pb.Base{Token: token}, Id: int64(id)})
	if err != nil {
		log.Println("could not greet: %v", err)
	}
	return Result{
		Result: r.GetStatus(),
		Info:   r.GetInfo(),
	}
}

func PauseVMClient(token, vmid string) Result {
	id, err := strconv.Atoi(vmid)
	if err != nil {
		log.Println("id error!!")
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.PauseVM(ctx, &pb.VMID{Base: &pb.Base{Token: token}, Id: int64(id)})
	if err != nil {
		log.Println("could not greet: %v", err)
	}
	return Result{
		Result: r.GetStatus(),
		Info:   r.GetInfo(),
	}
}

func ResumeVMClient(token, vmid string) Result {
	id, err := strconv.Atoi(vmid)
	if err != nil {
		log.Println("id error!!")
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ResumeVM(ctx, &pb.VMID{Base: &pb.Base{Token: token}, Id: int64(id)})
	if err != nil {
		log.Println("could not greet: %v", err)
	}
	return Result{
		Result: r.GetStatus(),
		Info:   r.GetInfo(),
	}
}

func GetUserVMClient(token string) []VMData {

	conn, err := grpc.Dial(GetgRPCServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Not connect; %v", err)
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

func GetVMClient(id, token string) VMData {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	vmid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("vmid error!!")
	}

	r, err := c.GetVM(ctx, &pb.VMID{Id: int64(vmid), Base: &pb.Base{Token: token}})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return VMData{
		NodeID:    int(r.Node),
		ID:        int(r.Option.Id),
		Name:      r.Vmname,
		CPU:       int(r.Vcpu),
		Mem:       int(r.Vmem),
		Net:       r.Vnet,
		AutoStart: r.Option.Autostart,
		Status:    int(r.Option.Status),
	}
}
