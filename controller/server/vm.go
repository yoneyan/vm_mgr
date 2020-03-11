package server

import (
	"context"
	"fmt"
	"github.com/yoneyan/vm_mgr/controller/data"
	"github.com/yoneyan/vm_mgr/controller/db"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
	"time"
)

type VMDataStruct struct {
	NodeID      int
	ID          int
	Name        string
	CPU         int
	Mem         int
	Storage     string
	StoragePath string
	CDROM       string
	StorageType int
	Image       int
	Net         string
	VNC         int
	AutoStart   bool
	Status      int
}

func (s *server) CreateVM(ctx context.Context, in *pb.VMData) (*pb.Result, error) {
	fmt.Println("----------CreateVM-----")
	log.Printf("Receive VMID: %v", in.GetOption().GetId())
	log.Printf("Receive name: %v", in.GetVmname())
	log.Printf("Receive cpu: %v", in.GetVcpu())
	log.Printf("Receive mem: %v", in.GetVmem())
	log.Printf("Receive StoragePath: %v", in.GetOption().StoragePath)
	log.Printf("Receive Storage: %v", in.GetStorage())
	log.Printf("Receive CDROM: %v", in.GetCdrom())
	log.Printf("Receive vnc: %v", in.GetOption().Vnc)
	log.Printf("Receive net: %v", in.GetVnet())
	log.Printf("Receive change: %v", in.GetOption().Autostart)

	user := in.GetBase().GetUser()
	pass := in.GetBase().GetPass()
	group := in.GetBase().GetGroup()

	if data.AdminUserCertification(user, pass) == false {
		if data.GroupAdminCertification(user, pass, group) == false {
			return &pb.Result{Status: false, Info: "Auth Failed!!"}, nil
		}
	}
	d, result := db.GetDBNodeID(int(in.GetNode()))
	fmt.Println(d)
	if result == false {
		return &pb.Result{Status: false, Info: "Node Not Found!!"}, nil
	}
	conn, err := grpc.Dial(d.IP, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	groupid, result := db.GetDBGroupID(group)
	if result == false {
		return &pb.Result{Status: false, Info: "GroupID Not Found!!"}, nil
	}

	name := strconv.Itoa(groupid) + "-" + strconv.Itoa(0) + "-" + in.GetVmname()
	fmt.Println("Name: " + name)

	r, err := c.CreateVM(ctx, &pb.VMData{
		Node:        in.GetNode(),
		Vmname:      name,
		Vcpu:        in.GetVcpu(),
		Vmem:        in.GetVmem(),
		Storagetype: in.GetStoragetype(),
		Storage:     in.GetStorage(),
		Cdrom:       in.GetCdrom(),
		Vnet:        in.GetVnet(),
		Option: &pb.Option{
			StoragePath: in.GetOption().GetStoragePath(),
			Image:       in.GetOption().GetImage(),
			Vnc:         in.GetOption().GetVnc(),
			Autostart:   in.GetOption().GetAutostart(),
		},
	})
	if err != nil {
		fmt.Printf("ERROR: ")
		fmt.Println(err)
	}
	return &pb.Result{Status: r.GetStatus(), Info: r.GetInfo()}, nil
}

func (s *server) DeleteVM(ctx context.Context, in *pb.VMID) (*pb.Result, error) {
	fmt.Println("----------DeleteVM-----")
	log.Printf("Receive VMID: %v", in.GetId())
	log.Println("Receive AuthUser  : " + in.GetBase().GetUser() + ", AuthPass: " + in.GetBase().GetPass() + ", Group: " + in.GetBase().GetGroup())

	nodeid := in.GetId() / 1000
	vmid := in.GetId() - (1000 * nodeid)

	user := in.GetBase().GetUser()
	pass := in.GetBase().GetPass()
	group := in.GetBase().GetGroup()

	if data.AdminUserCertification(user, pass) == false {
		if data.GroupAdminCertification(user, pass, group) == false {
			return &pb.Result{Status: false, Info: "Auth Failed!!"}, nil
		}
	}

	d, result := db.GetDBNodeID(int(nodeid))
	if result == false {
		return &pb.Result{Status: false, Info: "Node Not Found!!"}, nil
	}
	info, result := data.VMCertification(int(vmid), d.ID, d.IP)
	if result == false {
		return &pb.Result{Status: false, Info: info}, nil
	}
	conn, err := grpc.Dial(d.IP, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.DeleteVM(ctx, &pb.VMID{Id: vmid})
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	return &pb.Result{Status: r.GetStatus(), Info: r.GetInfo()}, nil
}

func (s *server) StartVM(ctx context.Context, in *pb.VMID) (*pb.Result, error) {
	fmt.Println("----------StartVM-----")
	log.Println("Receive VMID  : ", in.GetId())
	log.Println("Receive AuthUser  : " + in.GetBase().GetUser() + ", AuthPass: " + in.GetBase().GetPass() + ", Group: " + in.GetBase().GetGroup())

	nodeId := in.GetId() / 1000
	vmId := in.GetId() - (1000 * nodeId)

	user := in.GetBase().GetUser()
	pass := in.GetBase().GetPass()
	group := in.GetBase().GetGroup()

	d, result := db.GetDBNodeID(int(nodeId))
	if result == false {
		return &pb.Result{Status: false, Info: "Node Not Found!!"}, nil
	}
	if data.AdminUserCertification(user, pass) == false {
		if data.GroupAdminCertification(user, pass, group) == false {
			if data.GroupUserCertification(user, pass, group) == false {
				return &pb.Result{Status: false, Info: "Auth Failed!!"}, nil
			}
			info, result := data.VMCertification(int(vmId), d.ID, d.IP)
			if result == false {
				return &pb.Result{Status: false, Info: info}, nil
			}
		}
	}
	conn, err := grpc.Dial(d.IP, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.StartVM(ctx, &pb.VMID{Id: vmId})
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	fmt.Println(r.GetInfo())
	return &pb.Result{Status: r.GetStatus(), Info: r.GetInfo()}, nil
}

func (s *server) StopVM(ctx context.Context, in *pb.VMID) (*pb.Result, error) {
	fmt.Println("----------StartVM-----")
	log.Println("Receive VMID  : ", in.GetId())
	log.Println("Receive AuthUser  : " + in.GetBase().GetUser() + ", AuthPass: " + in.GetBase().GetPass() + ", Group: " + in.GetBase().GetGroup())

	nodeid := in.GetId() / 1000
	vmid := in.GetId() - (1000 * nodeid)

	user := in.GetBase().GetUser()
	pass := in.GetBase().GetPass()
	group := in.GetBase().GetGroup()

	d, result := db.GetDBNodeID(int(nodeid))
	if result == false {
		fmt.Println("DBNodeID Error")
		return &pb.Result{Status: false, Info: "Node Not Found!!"}, nil
	}

	if data.AdminUserCertification(user, pass) == false {
		if data.GroupAdminCertification(user, pass, group) == false {
			if data.GroupUserCertification(user, pass, group) == false {
				return &pb.Result{Status: false, Info: "Auth Failed!!"}, nil
			}
			info, result := data.VMCertification(int(vmid), d.ID, d.IP)
			if result == false {
				return &pb.Result{Status: false, Info: info}, nil
			}
		}
	}

	conn, err := grpc.Dial(d.IP, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println(d.IP)

	r, err := c.StopVM(ctx, &pb.VMID{Id: vmid})
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	fmt.Println(r.GetInfo())
	return &pb.Result{Status: r.GetStatus(), Info: r.GetInfo()}, nil
}

/*
func (s *server) GetVM(ctx context.Context, in *pb.VMID) (*pb.VMData, error) {
	fmt.Println("----------GetVMID-----")
	log.Printf("Receive VMID: %v", in.GetId())
	result, err := db.VMDBGetData(int(in.GetId()))
	if err != nil {
		fmt.Println("Error!!")
		return &pb.VMData{}, fmt.Errorf("Not Found!!")

	}
	return &pb.VMData{
		Option: &pb.Option{
			StoragePath: result.StoragePath,
			Vnc:         int32(result.Vnc),
			Id:          int64(result.ID),
			Autostart:   result.AutoStart,
		},
		Vmname: result.Name,
		Vcpu:   int64(result.CPU),
		Vmem:   int64(result.Mem),
		Vnet:   result.Net,
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
		Option: &pb.Option{
			StoragePath: result.StoragePath,
			Vnc:         int32(result.Vnc),
			Id:          int64(result.ID),
			Autostart:   result.AutoStart,
		},
		Vmname: result.Name,
		Vcpu:   int64(result.CPU),
		Vmem:   int64(result.Mem),
		Vnet:   result.Net,
	}, nil
}
*/

func (s *server) GetAllVM(base *pb.Base, stream pb.Grpc_GetAllVMServer) error {
	log.Println("----GetAllVM----")
	log.Println("Receive AuthUser  : " + base.GetUser() + ", AuthPass: " + base.GetPass())

	if data.AdminUserCertification(base.GetUser(), base.GetPass()) == false {
		fmt.Println("Failed authentication")
		return nil
	}

	var d []VMDataStruct

	for _, a := range db.GetDBAllNode() {
		conn, err := grpc.Dial(a.IP, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			fmt.Printf("Not connect; ")
			fmt.Println(err)
		}
		defer conn.Close()

		c := pb.NewGrpcClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		stream, err := c.GetAllVM(ctx, base)
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
			d = append(d, VMDataStruct{
				NodeID:    a.ID,
				ID:        int(article.Option.Id) + (1000 * a.ID),
				Name:      article.Vmname,
				CPU:       int(article.Vcpu),
				Mem:       int(article.Vmem),
				Net:       article.Vnet,
				AutoStart: article.Option.Autostart,
				Status:    int(article.Option.Status),
			})
		}
	}

	for _, a := range d {
		if err := stream.Send(&pb.VMData{Option: &pb.Option{Id: int64(a.ID), Autostart: a.AutoStart, Status: int32(a.Status)}, Vmname: a.Name, Node: int32(a.NodeID), Vcpu: int64(a.CPU), Vmem: int64(a.Mem), Vnet: a.Net}); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) ShutdownVM(ctx context.Context, in *pb.VMID) (*pb.Result, error) {
	log.Println("----ShutdownVM----")
	log.Println("Receive VMID  : ", in.GetId())
	log.Println("Receive AuthUser  : " + in.GetBase().GetUser() + ", AuthPass: " + in.GetBase().GetPass() + ", Group: " + in.GetBase().GetGroup())

	nodeid := in.GetId() / 1000
	vmid := in.GetId() - (1000 * nodeid)

	user := in.GetBase().GetUser()
	pass := in.GetBase().GetPass()
	group := in.GetBase().GetGroup()

	d, result := db.GetDBNodeID(int(nodeid))
	if result == false {
		fmt.Println("DBNodeID Error")
		return &pb.Result{Status: false, Info: "Node Not Found!!"}, nil
	}

	if data.AdminUserCertification(user, pass) == false {
		if data.GroupAdminCertification(user, pass, group) == false {
			if data.GroupUserCertification(user, pass, group) == false {
				return &pb.Result{Status: false, Info: "Auth Failed!!"}, nil
			}
			info, result := data.VMCertification(int(vmid), d.ID, d.IP)
			if result == false {
				return &pb.Result{Status: false, Info: info}, nil
			}
		}
	}
	conn, err := grpc.Dial(d.IP, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ShutdownVM(ctx, &pb.VMID{Id: vmid})
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	return &pb.Result{Status: r.GetStatus(), Info: r.GetInfo()}, nil
}

func (s *server) ResetVM(ctx context.Context, in *pb.VMID) (*pb.Result, error) {
	log.Println("----RebootVM----")
	log.Println("Receive VMID  : ", in.GetId())
	log.Println("Receive AuthUser  : " + in.GetBase().GetUser() + ", AuthPass: " + in.GetBase().GetPass() + ", Group: " + in.GetBase().GetGroup())

	nodeid := in.GetId() / 1000
	vmid := in.GetId() - (1000 * nodeid)

	user := in.GetBase().GetUser()
	pass := in.GetBase().GetPass()
	group := in.GetBase().GetGroup()

	d, result := db.GetDBNodeID(int(nodeid))
	if result == false {
		fmt.Println("DBNodeID Error")
		return &pb.Result{Status: false, Info: "Node Not Found!!"}, nil
	}

	if data.AdminUserCertification(user, pass) == false {
		if data.GroupAdminCertification(user, pass, group) == false {
			if data.GroupUserCertification(user, pass, group) == false {
				return &pb.Result{Status: false, Info: "Auth Failed!!"}, nil
			}
			info, result := data.VMCertification(int(vmid), d.ID, d.IP)
			if result == false {
				return &pb.Result{Status: false, Info: info}, nil
			}
		}
	}
	conn, err := grpc.Dial(d.IP, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ResetVM(ctx, &pb.VMID{Id: vmid})
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	return &pb.Result{Status: r.GetStatus(), Info: r.GetInfo()}, nil
}

func (s *server) PauseVM(ctx context.Context, in *pb.VMID) (*pb.Result, error) {
	log.Println("----PauseVM----")
	log.Println("Receive VMID  : ", in.GetId())
	log.Println("Receive AuthUser  : " + in.GetBase().GetUser() + ", AuthPass: " + in.GetBase().GetPass() + ", Group: " + in.GetBase().GetGroup())

	nodeid := in.GetId() / 1000
	vmid := in.GetId() - (1000 * nodeid)

	user := in.GetBase().GetUser()
	pass := in.GetBase().GetPass()
	group := in.GetBase().GetGroup()

	d, result := db.GetDBNodeID(int(nodeid))
	if result == false {
		fmt.Println("DBNodeID Error")
		return &pb.Result{Status: false, Info: "Node Not Found!!"}, nil
	}

	if data.AdminUserCertification(user, pass) == false {
		if data.GroupAdminCertification(user, pass, group) == false {
			if data.GroupUserCertification(user, pass, group) == false {
				return &pb.Result{Status: false, Info: "Auth Failed!!"}, nil
			}
			info, result := data.VMCertification(int(vmid), d.ID, d.IP)
			if result == false {
				return &pb.Result{Status: false, Info: info}, nil
			}
		}
	}
	conn, err := grpc.Dial(d.IP, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.PauseVM(ctx, &pb.VMID{Id: vmid})
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	return &pb.Result{Status: r.GetStatus(), Info: r.GetInfo()}, nil
}

func (s *server) ResumeVM(ctx context.Context, in *pb.VMID) (*pb.Result, error) {
	fmt.Println("-----ResumeVM-----")
	log.Println("Receive VMID  : ", in.GetId())
	log.Println("Receive AuthUser  : " + in.GetBase().GetUser() + ", AuthPass: " + in.GetBase().GetPass() + ", Group: " + in.GetBase().GetGroup())

	nodeid := in.GetId() / 1000
	vmid := in.GetId() - (1000 * nodeid)

	user := in.GetBase().GetUser()
	pass := in.GetBase().GetPass()
	group := in.GetBase().GetGroup()

	d, result := db.GetDBNodeID(int(nodeid))
	if result == false {
		fmt.Println("DBNodeID Error")
		return &pb.Result{Status: false, Info: "Node Not Found!!"}, nil
	}

	if data.AdminUserCertification(user, pass) == false {
		if data.GroupAdminCertification(user, pass, group) == false {
			if data.GroupUserCertification(user, pass, group) == false {
				return &pb.Result{Status: false, Info: "Auth Failed!!"}, nil
			}
			info, result := data.VMCertification(int(vmid), d.ID, d.IP)
			if result == false {
				return &pb.Result{Status: false, Info: info}, nil
			}
		}
	}
	conn, err := grpc.Dial(d.IP, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ResumeVM(ctx, &pb.VMID{Id: vmid})
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	return &pb.Result{Status: r.GetStatus(), Info: r.GetInfo()}, nil
}
