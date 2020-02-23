package data

import (
	"context"
	"fmt"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"log"
	_ "os"
	"time"
)

const (
	address     = "localhost:50100"
	defaultName = "world"
)

func grpc_client_test() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewVMClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateVM(ctx, &pb.VMData{Id: 1, Vmname: "test", Vcpu: 1, Vmem: 1024, Vnet: "br100", Vnc: 10000})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetStatus())
}

func CreateVM(name string, vcpu, vmem, storage int64, storage_path string, cdrom string, vnet string, vnc int64, change bool) bool {
	//value verification
	i := []int64{vcpu, vmem, storage, vnc}
	s := []string{name, storage_path, vnet}
	for _, a := range i {
		if a == 0 {
			fmt.Println("Value False!!")
			fmt.Printf("Debug: ")
			fmt.Println(a)
			return false
		}
	}
	for _, s := range s {
		if s == "none" {
			fmt.Println("Value False!!")
			fmt.Println("Debug: " + s)
			return false
		}
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewVMClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateVM(ctx, &pb.VMData{Vmname: name, Vcpu: vcpu, Vmem: vmem, Vnet: vnet, Vnc: vnc, Storage: storage, StoragePath: storage_path, CdromPath: cdrom, Change: change})
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
	c := pb.NewVMClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.DeleteVM(ctx, &pb.VMID{Id: id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetStatus())
	return r.GetStatus()
}
