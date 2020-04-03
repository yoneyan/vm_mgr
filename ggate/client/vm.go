package client

import (
	"context"
	"fmt"
	"github.com/yoneyan/vm_mgr/ggate/etc"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
	"time"
)

type VMDataResult struct {
	NodeID    int    `json:"nodeid"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CPU       int    `json:"cpu"`
	Mem       int    `json:"mem"`
	Net       string `json:"net"`
	Storage   string `json:"storage"`
	VNCUrl    string `json:"vncurl"`
	AutoStart bool   `json:"autostart"`
	Status    int    `json:"status"`
}

type CreateVMData struct {
	NodeID      int    `json:"nodeid"`
	VMName      string `json:"vmname"`
	Group       string `json:"group"`
	CPU         int    `json:"cpu"`
	Mem         int    `json:"mem"`
	Storage     int    `json:"storage"`
	StorageType int    `json:"storagetype"`
	AutoStart   bool   `json:"autostart"`
	ImageName   string `json:"imagename"`
	ImageType   string `json:"imagetype"`
}

func CreateVM(vm CreateVMData, token string) Result {

	conn, err := grpc.Dial(GetgRPCServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.CreateVM(ctx, &pb.VMData{
		Base:    &pb.Base{Token: token, Group: vm.Group},
		Node:    int32(vm.NodeID),
		Vmname:  vm.VMName,
		Vcpu:    int64(vm.CPU),
		Vmem:    int64(vm.Mem),
		Type:    1,
		Storage: strconv.Itoa(vm.Storage),
		Option:  &pb.Option{StoragePath: strconv.Itoa(vm.StorageType), Autostart: true},
		Image:   &pb.Image{Name: vm.ImageName, Tag: vm.ImageType},
	})
	if err != nil {
		log.Println("could not greet: %v", err)
	}
	return Result{
		Result: r.GetStatus(),
		Info:   r.GetInfo(),
	}
}

func DeleteVM(id, token string) Result {

	conn, err := grpc.Dial(GetgRPCServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	vmid, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("vmid error!!")
	}

	r, err := c.DeleteVM(ctx, &pb.VMID{Id: int64(vmid), Base: &pb.Base{Token: token}})
	if err != nil {
		log.Println("could not greet: %v", err)
	}
	return Result{
		Result: r.GetStatus(),
		Info:   r.GetInfo(),
	}
}

func StartVMClient(token, vmid string) Result {
	id, err := strconv.Atoi(vmid)
	if err != nil {
		log.Println("id error!!")
	}

	conn, err := grpc.Dial(GetgRPCServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
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

	conn, err := grpc.Dial(GetgRPCServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
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

	conn, err := grpc.Dial(GetgRPCServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
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

	conn, err := grpc.Dial(GetgRPCServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
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

	conn, err := grpc.Dial(GetgRPCServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
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

	conn, err := grpc.Dial(GetgRPCServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
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

func GetUserVMClient(token string) []VMDataResult {

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
		log.Println(err)
	}

	var data []VMDataResult

	for {
		r, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
		tmp := VMDataResult{
			NodeID:    int(r.GetNode()),
			ID:        int(r.Option.GetId()),
			Name:      r.Vmname,
			CPU:       int(r.Vcpu),
			Mem:       int(r.Vmem),
			Net:       r.GetVnet(),
			Storage:   r.GetStorage(),
			VNCUrl:    r.Option.GetVncurl(),
			AutoStart: r.Option.GetAutostart(),
			Status:    int(r.Option.GetStatus()),
		}
		data = append(data, tmp)
	}
	fmt.Println(data)
	return data
}

func GetVMClient(id, token string) VMDataResult {

	conn, err := grpc.Dial(GetgRPCServerAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Not connect; %v", err)
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
		log.Println("could not greet: %v", err)
	}

	isHttps, domain := etc.GetDomain()

	if isHttps {
		domain = "https://" + domain
	} else {
		domain = "http://" + domain
	}

	return VMDataResult{
		NodeID:    int(r.Node),
		ID:        int(r.Option.Id),
		Name:      r.Vmname,
		CPU:       int(r.Vcpu),
		Mem:       int(r.Vmem),
		Net:       r.Vnet,
		Storage:   r.Option.StoragePath,
		VNCUrl:    domain + r.Option.Vncurl,
		AutoStart: r.Option.Autostart,
		Status:    int(r.Option.Status),
	}
}
