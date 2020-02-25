package data

import (
	"context"
	"fmt"
	"github.com/yoneyan/vm_mgr/node/db"
	"github.com/yoneyan/vm_mgr/node/etc"
	"github.com/yoneyan/vm_mgr/node/manage"
	"github.com/yoneyan/vm_mgr/node/vm"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":50100"

type server struct {
	pb.UnimplementedVMServer
}

func (s *server) CreateVM(ctx context.Context, in *pb.VMData) (*pb.Result, error) {
	fmt.Println("----------CreateVM-----")
	log.Printf("Receive VMID: %v", in.GetId())
	log.Printf("Receive name: %v", in.GetVmname())
	log.Printf("Receive cpu: %v", in.GetVcpu())
	log.Printf("Receive mem: %v", in.GetVmem())
	log.Printf("Receive StoragePath: %v", in.GetStoragePath())
	log.Printf("Receive Storage: %v", in.GetStorage())
	log.Printf("Receive vnc: %v", in.GetVnc())
	log.Printf("Receive net: %v", in.GetVnet())
	log.Printf("Receive change: %v", in.GetAutostart())
	var r vm.CreateVMInformation

	r.ID = int(in.GetId())
	r.Name = in.GetVmname()
	r.CPU = int(in.GetVcpu())
	r.Mem = int(in.GetVmem())
	r.StoragePath = in.GetStoragePath()
	r.CDROM = in.GetCdromPath()
	r.Net = in.GetVnet()
	r.VNC = int(in.GetVnc())
	r.AutoStart = bool(in.GetAutostart())

	if etc.FileExists(in.GetStoragePath()+"/"+in.GetVmname()+".img") == false {
		fmt.Println("Not storage file exists")
		storage := manage.Storage{
			Path:   in.GetStoragePath(),
			Name:   in.GetVmname(),
			Format: "qcow2",
			Size:   int(in.GetStorage()),
		}
		manage.CreateStorage(&storage)
	}

	err := vm.CreateVMProcess(&r)
	if err != nil {
		return &pb.Result{Status: false}, nil
	}
	return &pb.Result{Status: true}, nil
}

func (s *server) DeleteVM(ctx context.Context, in *pb.VMID) (*pb.Result, error) {
	fmt.Println("----------DeleteVM-----")
	log.Printf("Receive VMID: %v", in.GetId())

	result := vm.DeleteVMProcess(int(in.GetId()))
	if result {
		fmt.Println("Delete success!!")
		return &pb.Result{Status: true}, nil
	} else {
		fmt.Println("Delete Failed....")
		return &pb.Result{Status: false}, nil
	}
}

func (s *server) StartVM(ctx context.Context, in *pb.VMID) (*pb.Result, error) {
	fmt.Println("----------StartVM-----")
	log.Printf("Receive VMID: %v", in.GetId())
	if vm.StartVMProcess(int(in.GetId())) {
		return &pb.Result{Status: true}, nil
	} else {
		return &pb.Result{Status: false}, nil
	}
}

func (s *server) StopVM(ctx context.Context, in *pb.VMID) (*pb.Result, error) {
	fmt.Println("----------StopVM-----")
	log.Printf("Receive VMID: %v", in.GetId())
	err := vm.VMStop(int(in.GetId()))
	if err != nil {
		fmt.Println("VMStop Error!!")
	}

	return &pb.Result{Status: true}, nil
}

func (s *server) GetVM(ctx context.Context, in *pb.VMID) (*pb.VMData, error) {
	fmt.Println("----------GetVMID-----")
	log.Printf("Receive VMID: %v", in.GetId())
	result, err := db.VMDBGetData(int(in.GetId()))
	if err != nil {
		fmt.Println("Error!!")
		return &pb.VMData{}, fmt.Errorf("Not Found!!")

	}
	return &pb.VMData{
		Id:          int64(result.ID),
		Vmname:      result.Name,
		Vcpu:        int64(result.CPU),
		Vmem:        int64(result.Mem),
		StoragePath: result.StoragePath,
		Vnet:        result.Net,
		Vnc:         int64(result.Vnc),
		Autostart:   bool(result.AutoStart),
	}, nil
}

func (s *server) GetVMName(ctx context.Context, in *pb.VMName) (*pb.VMData, error) {
	fmt.Println("----------GetVMName-----")
	log.Printf("Receive Name: %v", in.GetVmname())
	id, err := db.VMDBGetVMID(in.GetVmname())
	if err != nil {
		fmt.Println("NotFound VMID !!")
		return &pb.VMData{}, fmt.Errorf("Not Found VMID!!")
	}
	result, err := db.VMDBGetData(id)
	if err != nil {
		fmt.Println("Not Found!!")
		return &pb.VMData{}, fmt.Errorf("Not Found!!")

	}
	return &pb.VMData{
		Id:          int64(result.ID),
		Vmname:      result.Name,
		Vcpu:        int64(result.CPU),
		Vmem:        int64(result.Mem),
		StoragePath: result.StoragePath,
		Vnet:        result.Net,
		Vnc:         int64(result.Vnc),
		Autostart:   bool(result.AutoStart),
	}, nil
}

func (s *server) GetAllVM(ctx context.Context, in *pb.VMID) (*pb.Result, error) {
	log.Println("----GetAllVM----")
	log.Printf("Receive GetAllVM")
	fmt.Println(db.GetDBAll())
	return &pb.Result{Status: true}, nil
}

func (s *server) ShutdownVM(ctx context.Context, in *pb.VMID) (*pb.Result, error) {
	log.Println("----GetAllVM----")
	log.Printf("Receive GetAllVM")
	fmt.Println(db.GetDBAll())
	return &pb.Result{Status: true}, nil
}

func Processgrpc() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterVMServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
