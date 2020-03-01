package client

//Node gRPC Client

import (
	"context"
	"fmt"
	"github.com/yoneyan/vm_mgr/controller/data"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"log"
	_ "os"
	"time"
)

const (
	address = "localhost:50100"
)

//name string, vcpu, vmem, storage int64, storage_path string, cdrom string, vnet string, vnc int64, autostart bool
func CreateVM(address string, d *pb.VMData) bool {
	if data.CreateVMCheck(d) == false {
		fmt.Println("Valid value!!")
		return false
	}
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateVM(ctx, d)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetStatus())
	return r.GetStatus()
}

func DeleteVM(id int64) bool {
	if id < 1 {
		fmt.Println("Value False")
		fmt.Printf("Debug: value is ")
		fmt.Println(id)
		return false
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.DeleteVM(ctx, &pb.VMID{Id: id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetStatus())
	return r.GetStatus()
}

func StartVM(id int64) bool {
	if id < 1 {
		fmt.Println("Value False")
		fmt.Printf("Debug: value is ")
		fmt.Println(id)
		return false
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.StartVM(ctx, &pb.VMID{Id: id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetStatus())
	return r.GetStatus()
}

func StopVM(id int64) bool {
	if id < 1 {
		fmt.Println("Value False")
		fmt.Printf("Debug: value is ")
		fmt.Println(id)
		return false
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.StopVM(ctx, &pb.VMID{Id: id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetStatus())
	return r.GetStatus()
}

func ShutdownVM(id int64) bool {
	if id < 1 {
		fmt.Println("Value False")
		fmt.Printf("Debug: value is ")
		fmt.Println(id)
		return false
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ShutdownVM(ctx, &pb.VMID{Id: id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetStatus())
	return r.GetStatus()
}

func ResetVM(id int64) bool {
	if id < 1 {
		fmt.Println("Value False")
		fmt.Printf("Debug: value is ")
		fmt.Println(id)
		return false
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.StopVM(ctx, &pb.VMID{Id: id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetStatus())
	return r.GetStatus()
}

func PauseVM(id int64) bool {
	if id < 1 {
		fmt.Println("Value False")
		fmt.Printf("Debug: value is ")
		fmt.Println(id)
		return false
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.PauseVM(ctx, &pb.VMID{Id: id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetStatus())
	return r.GetStatus()
}

func ResumeVM(id int64) bool {
	if id < 1 {
		fmt.Println("Value False")
		fmt.Printf("Debug: value is ")
		fmt.Println(id)
		return false
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ResumeVM(ctx, &pb.VMID{Id: id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetStatus())
	return r.GetStatus()
}

func GetVM(id int64) {
	if id < 1 {
		fmt.Println("Value False")
		fmt.Printf("Debug: value is ")
		fmt.Println(id)
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetVM(ctx, &pb.VMID{Id: id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	if r.GetVmname() == "" {
		fmt.Println("None")
	}
	log.Printf("ID:        %d", r.Option.GetId())
	log.Printf("VMName:    %s", r.GetVmname())
	log.Printf("cpu:       %d", r.GetVcpu())
	log.Printf("memory:    %d", r.GetVmem())
	log.Printf("Storage:   %s", r.Option.GetStoragePath())
	log.Printf("VNC:       %d", r.Option.GetVnc())
	log.Printf("Net:       %s", r.GetVnet())
	log.Printf("AutoStart: %t", r.Option.GetAutostart())
}

func GetVMName(name string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetVMName(ctx, &pb.VMName{Vmname: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("ID:        %d", r.Option.GetId())
	log.Printf("VMName:    %s", r.GetVmname())
	log.Printf("cpu:       %d", r.GetVcpu())
	log.Printf("memory:    %d", r.GetVmem())
	log.Printf("Storage:   %s", r.Option.GetStoragePath())
	log.Printf("VNC:       %d", r.Option.GetVnc())
	log.Printf("Net:       %s", r.GetVnet())
	log.Printf("AutoStart: %t", r.Option.GetAutostart())

}

func GetAllVM(id int) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetAllVM(ctx, &pb.VMID{Id: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("ID: ")
	log.Println(r.Status)
}

func NodeAllStop(address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.StopNode(ctx, &pb.NodeID{}) //Timer do not work.
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("ID: ")
	log.Println(r.Status)
}
