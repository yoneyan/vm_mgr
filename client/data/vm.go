package data

import (
	"context"
	"fmt"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"io"
	"log"
	_ "os"
	"strconv"
	"time"
)

//name string, vcpu, vmem, storage int64, storage_path string, cdrom string, vnet string, vnc int64, autostart bool
func CreateVM(d *pb.VMData, address string) {
	if CreateVMCheck(d) == false {
		log.Fatal("Valid value!!")
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
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func DeleteVM(d *pb.VMID, address string) {
	if d.Id < 1 {
		fmt.Println("Value False")
		fmt.Printf("Debug: value is ")
		fmt.Println(d.Id)
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.DeleteVM(ctx, d)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetStatus())
	log.Printf("Status: ")
	log.Println(r.Status)

	log.Printf("Info: ")
	log.Println(r.Info)
}

func StartVM(d *pb.VMID, address string) bool {
	if d.Id < 1 {
		fmt.Println("Value False")
		fmt.Printf("Debug: value is ")
		fmt.Println(d.Id)
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
	r, err := c.StartVM(ctx, d)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetStatus())
	return r.GetStatus()
}

func StopVM(d *pb.VMID, address string) bool {
	if d.Id < 1 {
		fmt.Println("Value False")
		fmt.Printf("Debug: value is ")
		fmt.Println(d.Id)
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
	r, err := c.StopVM(ctx, d)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetStatus())
	return r.GetStatus()
}

func ShutdownVM(d *pb.VMID, address string) bool {
	if d.Id < 1 {
		fmt.Println("Value False")
		fmt.Printf("Debug: value is ")
		fmt.Println(d.Id)
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
	r, err := c.ShutdownVM(ctx, d)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetStatus())
	return r.GetStatus()
}

func ResetVM(d *pb.VMID, address string) bool {
	if d.Id < 1 {
		fmt.Println("Value False")
		fmt.Printf("Debug: value is ")
		fmt.Println(d.Id)
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
	r, err := c.StopVM(ctx, d)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetStatus())
	return r.GetStatus()
}

func PauseVM(d *pb.VMID, address string) bool {
	if d.Id < 1 {
		fmt.Println("Value False")
		fmt.Printf("Debug: value is ")
		fmt.Println(d.Id)
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
	r, err := c.PauseVM(ctx, d)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetStatus())
	return r.GetStatus()
}

func ResumeVM(d *pb.VMID, address string) bool {
	if d.Id < 1 {
		fmt.Println("Value False")
		fmt.Printf("Debug: value is ")
		fmt.Println(d.Id)
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
	r, err := c.ResumeVM(ctx, d)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetStatus())
	return r.GetStatus()
}

func GetVM(d *pb.VMID, address string) {
	if d.Id < 1 {
		fmt.Println("Value False")
		fmt.Printf("Debug: value is ")
		fmt.Println(d.Id)
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetVM(ctx, d)
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

func GetVMName(d *pb.VMName, address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetVMName(ctx, d)
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

func GetAllVM(d *pb.Base, address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()

	c := pb.NewGrpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.GetAllVM(ctx, d)
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
		fmt.Printf("ID: " + strconv.Itoa(int(article.Option.Id)) + " Name: " + article.Vmname + " CPU: " + strconv.Itoa(int(article.Vcpu)) + " Mem: " + strconv.Itoa(int(article.Vmem)))
		fmt.Println(" Net: " + article.Vnet + " AutoStart: " + strconv.FormatBool(article.Option.Autostart) + " status: " + strconv.Itoa(int(article.Option.Status)))
	}
}
