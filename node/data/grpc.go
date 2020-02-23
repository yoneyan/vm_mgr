package data

import (
	"context"
	"fmt"
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
	log.Println("----CreateVM----")
	log.Printf("Receive ID: %v", in.GetId())
	log.Printf("Receive name: %v", in.GetVmname())
	log.Printf("Receive cpu: %v", in.GetVcpu())
	log.Printf("Receive mem: %v", in.GetVmem())
	log.Printf("Receive StoragePath: %v", in.GetStoragePath())
	log.Printf("Receive Storage: %v", in.GetStorage())
	log.Printf("Receive vnc: %v", in.GetVnc())
	log.Printf("Receive net: %v", in.GetVnet())
	log.Printf("Receive change: %v", in.GetChange())

	var r vm.CreateVMInformation

	r.Name = in.GetVmname()
	r.CPU = int(in.GetVcpu())
	r.Mem = int(in.GetVmem())
	r.StoragePath = in.GetStoragePath()
	r.CDROM = in.GetCdromPath()
	r.Net = in.GetVnet()
	r.VNC = int(in.GetVnc())

	vm.CreateGenerateCmd(&r)

	if etc.FileExists(in.GetStoragePath()) == false {
		storage := manage.Storage{
			Path:   in.GetStoragePath(),
			Name:   in.GetVmname(),
			Format: "qcow2",
			Size:   int(in.GetStorage()),
		}
		manage.CreateStorage(&storage)
	}

	if manage.VMExistsName(in.GetVmname()) || manage.VMExistsID(int(in.GetId())) {
		id := int(in.GetId())
		vm.StartVMProcess(id)
	} else {
		fmt.Println("----VMNewCreate")
		vm.CreateVMDBProcess(&r)
		vm.RunQEMUCmd(vm.CreateGenerateCmd(&r))
	}
	return &pb.Result{Status: false}, nil
}

func (s *server) DeleteVM(ctx context.Context, in *pb.VMID) (*pb.Result, error) {
	log.Println("----DeleteVM----")
	log.Printf("Receive ID: %v", in.GetId())
	return &pb.Result{Status: false}, nil
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
