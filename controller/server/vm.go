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
	"strings"
	"time"
)

type VMDataStruct struct {
	NodeID      int
	Group       string
	GroupID     int
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
	log.Println("Receive AuthUser: " + in.Base.GetUser() + ", AuthPass: " + in.Base.GetPass())
	log.Println("Receive Token     : " + in.Base.GetToken())

	if in.Base.GetGroup() == "" {
		return &pb.Result{Status: false, Info: "Group is not specified!!"}, nil
	}

	user := in.Base.GetUser()
	pass := in.Base.GetPass()
	group := in.Base.GetGroup()

	var image string
	var createtype int32

	if data.SuperUserCertification(&data.UserCertData{User: user, Pass: pass, Group: group, Token: in.Base.GetToken()}) == false {
		return &pb.Result{Status: false, Info: "Auth Failed!!"}, nil
	}
	isAdmin := false
	if data.AdminUserCertification(user, pass, in.Base.GetToken()) {
		isAdmin = true
	} else {
		if data.CheckMaxSpec(in) == false {
			return &pb.Result{Status: false, Info: "Spec has reached the upper limit...."}, nil
		}
	}

	if data.CheckOnlyAdmin(int(in.GetNode())) && isAdmin == false {
		return &pb.Result{Status: false, Info: "Node is only Admin...."}, nil
	}

	address, result := data.CheckNodeID(isAdmin, int(in.GetNode()))
	if result == false {
		return &pb.Result{Status: false, Info: "Node Not Found!!"}, nil
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if isAdmin && in.GetType() == 0 {
		createtype = 0
		fmt.Println("Admin Mode(manual)")
	} else if isAdmin {
		createtype = 1
		image = data.GetImagePath(in)
		if image == "" {
			fmt.Println("Image Path Error!!")
			return &pb.Result{Status: false, Info: "Image Path Error!!"}, nil
		}
		sdata := data.GenerateDiskPath(in)
		if sdata.Result == false {
			return &pb.Result{Status: false, Info: "GenerateDiskPath Error!!"}, nil
		}
		in.Option.StoragePath = sdata.Path
		in.Storage = sdata.Size
	} else if isAdmin == false {
		createtype = 11
		image = data.GetImagePath(in)
		if image == "" {
			fmt.Println("Image Path Error!!")
			return &pb.Result{Status: false, Info: "Image Path Error!!"}, nil
		}
		sdata := data.GenerateDiskPath(in)
		if sdata.Result == false {
			return &pb.Result{Status: false, Info: "GenerateDiskPath Error!!"}, nil
		}
		in.Option.StoragePath = sdata.Path
		in.Storage = sdata.Size
	} else {
		fmt.Println("StorageType Error!!")
		return &pb.Result{Status: false, Info: "Error: StorageType Error!!"}, nil

	}
	groupid, result := db.GetDBGroupID(group)
	if result == false {
		return &pb.Result{Status: false, Info: "GroupID Not Found!!"}, nil
	}
	name := strconv.Itoa(groupid) + "-" + strconv.Itoa(0) + "-" + in.GetVmname()
	fmt.Println("Name: " + name)

	r, err := c.CreateVM(ctx, &pb.VMData{
		Node:    in.GetNode(),
		Vmname:  name,
		Vcpu:    in.GetVcpu(),
		Vmem:    in.GetVmem(),
		Type:    createtype,
		Storage: in.GetStorage(),
		Cdrom:   in.GetCdrom(),
		Vnet:    in.GetVnet(),
		Option: &pb.Option{
			StoragePath: in.GetOption().GetStoragePath(),
			Vnc:         in.GetOption().GetVnc(),
			Autostart:   in.GetOption().GetAutostart(),
		},
		Image: &pb.Image{Path: image},
	})
	if err != nil {
		fmt.Printf("ERROR: ")
		fmt.Println(err)
	}
	return &pb.Result{Status: r.Status, Info: r.Info}, nil
}

func (s *server) DeleteVM(ctx context.Context, in *pb.VMID) (*pb.Result, error) {
	fmt.Println("----------DeleteVM-----")
	log.Printf("Receive VMID: %v", in.GetId())
	log.Println("Receive AuthUser: " + in.Base.GetUser() + ", AuthPass: " + in.Base.GetPass())
	log.Println("Receive Token     : " + in.Base.GetToken())

	nodeId := in.GetId() / 1000
	vmId := in.GetId() - (1000 * nodeId)

	user := in.Base.GetUser()
	pass := in.Base.GetPass()

	address, result := data.SuperUserVMCertification(&data.UserCertData{
		User:   user,
		Pass:   pass,
		Token:  in.Base.GetToken(),
		VMID:   int(vmId),
		NodeID: int(nodeId),
	})
	if result == false {
		return &pb.Result{Status: false, Info: "Auth Failed!!"}, nil
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.DeleteVM(ctx, &pb.VMID{Id: vmId})
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	fmt.Println(r.GetInfo())
	return &pb.Result{Status: r.GetStatus(), Info: r.GetInfo()}, nil
}

func (s *server) StartVM(ctx context.Context, in *pb.VMID) (*pb.Result, error) {
	fmt.Println("----------StartVM-----")
	log.Println("Receive VMID  : ", in.GetId())
	log.Println("Receive AuthUser  : " + in.Base.GetUser() + ", AuthPass: " + in.Base.GetPass())
	log.Println("Receive Token     : " + in.Base.GetToken())

	nodeId := in.GetId() / 1000
	vmId := in.GetId() - (1000 * nodeId)

	user := in.Base.GetUser()
	pass := in.Base.GetPass()

	address, result := data.StandardUserVMCertification(&data.UserCertData{
		User:   user,
		Pass:   pass,
		Token:  in.Base.GetToken(),
		VMID:   int(vmId),
		NodeID: int(nodeId),
	})
	if result == false {
		return &pb.Result{Status: false, Info: "Auth Failed!!"}, nil
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
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
	log.Println("Receive AuthUser  : " + in.Base.GetUser() + ", AuthPass: " + in.Base.GetPass())
	log.Println("Receive Token     : " + in.Base.GetToken())

	nodeId := in.GetId() / 1000
	vmId := in.GetId() - (1000 * nodeId)

	user := in.Base.GetUser()
	pass := in.Base.GetPass()

	address, result := data.StandardUserVMCertification(&data.UserCertData{
		User:   user,
		Pass:   pass,
		Token:  in.Base.GetToken(),
		VMID:   int(vmId),
		NodeID: int(nodeId),
	})
	if result == false {
		return &pb.Result{Status: false, Info: "Auth Failed!!"}, nil
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println(address)

	r, err := c.StopVM(ctx, &pb.VMID{Id: vmId})
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	fmt.Println(r.GetInfo())
	return &pb.Result{Status: r.GetStatus(), Info: r.GetInfo()}, nil
}

func (s *server) GetVM(ctx context.Context, in *pb.VMID) (*pb.VMData, error) {
	fmt.Println("----------GetVMID-----")
	log.Printf("Receive VMID: %v", in.GetId())
	log.Println("Receive AuthUser  : " + in.Base.GetUser() + ", AuthPass: " + in.Base.GetPass())
	log.Println("Receive Token     : " + in.Base.GetToken())

	nodeId := in.GetId() / 1000
	vmId := in.GetId() - (1000 * nodeId)

	user := in.Base.GetUser()
	pass := in.Base.GetPass()

	address, result := data.StandardUserVMCertification(&data.UserCertData{
		User:   user,
		Pass:   pass,
		Token:  in.Base.GetToken(),
		VMID:   int(vmId),
		NodeID: int(nodeId),
	})
	if result == false {
		fmt.Println(address)
		return &pb.VMData{}, nil
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println(address)

	r, err := c.GetVM(ctx, &pb.VMID{Id: vmId})
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	fmt.Println(r)

	return &pb.VMData{
		Node:    int32(nodeId),
		Vmname:  r.GetVmname(),
		Vcpu:    r.GetVcpu(),
		Vmem:    r.GetVmem(),
		Storage: r.GetStorage(),
		Vnet:    r.GetVnet(),
		Option: &pb.Option{
			Vnc:       r.Option.GetVnc(),
			Id:        in.GetId(),
			Autostart: r.Option.GetAutostart(),
			Status:    r.Option.GetStatus(),
		},
	}, nil
}

/*
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

func (s *server) GetUserVM(base *pb.Base, stream pb.Grpc_GetUserVMServer) error {
	token := base.GetToken()
	log.Println("----GetUserVM----")
	log.Println("Receive AuthUser  : " + base.GetUser() + ", AuthPass: " + base.GetPass())
	log.Println("Receive UserID    : " + strconv.Itoa(int(base.GetUserid())))
	log.Println("Receive Token     : " + token)

	//user := base.GetUser()
	//pass := base.GetPass()

	d1, result := db.GetDBToken(token)
	if result == false {
		fmt.Println("Error GetToken")
		return nil
	}
	//if d1.Userid != int(base.GetUserid()) {
	//	fmt.Println("Wrong UserID!!")
	//	return nil
	//}
	d2, result := db.GetDBUser(d1.Userid)
	if result == false {
		fmt.Println("Error GetDBUser")
		return nil
	}
	d3, result := data.SearchUserForAllGroup(d2.Name)
	if result == false {
		fmt.Println("Error GetAllGroup")
		return nil
	}
	fmt.Println(d3)

	var isAdmin bool

	if data.AdminUserCertification(base.GetUser(), base.GetPass(), base.GetToken()) {
		isAdmin = true
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
			fmt.Println(err)
		}

		for {
			article, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
			}
			s := strings.Split(article.Vmname, "-")
			for _, b := range d3 {
				if s[0] == strconv.Itoa(b) || isAdmin {
					d = append(d, VMDataStruct{
						GroupID:   b,
						NodeID:    a.ID,
						ID:        int(article.Option.Id) + (1000 * a.ID),
						Name:      article.Vmname,
						CPU:       int(article.Vcpu),
						Mem:       int(article.Vmem),
						Storage:   article.Storage,
						Net:       article.Vnet,
						AutoStart: article.Option.Autostart,
						Status:    int(article.Option.Status),
					})
					break
				}
			}
		}
	}

	for _, a := range d {
		if err := stream.Send(&pb.VMData{Base: &pb.Base{Groupid: int32(a.GroupID)}, Option: &pb.Option{Id: int64(a.ID), Autostart: a.AutoStart, Status: int32(a.Status)}, Vmname: a.Name, Node: int32(a.NodeID), Vcpu: int64(a.CPU), Vmem: int64(a.Mem), Vnet: a.Net}); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) GetGroupVM(base *pb.Base, stream pb.Grpc_GetGroupVMServer) error {
	log.Println("----GetGroupVM----")
	log.Println("Receive AuthUser  : " + base.GetUser() + ", AuthPass: " + base.GetPass() + ", Group: " + base.GetGroup())
	log.Println("Receive Token     : " + base.GetToken())

	user := base.GetUser()
	pass := base.GetPass()
	group := base.GetGroup()
	token := base.GetToken()

	if group == "" {
		fmt.Println("Group is not specified!!")
		return nil
	}

	if data.StandardUserCertification(&data.UserCertData{User: user, Pass: pass, Group: group, Token: token}) == false {
		fmt.Println("Auth Failed!!")
		return nil
	}

	var d []VMDataStruct

	groupid, result := db.GetDBGroupID(group)
	if result == false {
		fmt.Println("DB Not Found Group!!")
		return nil
	}

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
			fmt.Println(err)
		}
		for {
			article, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
			}
			s := strings.Split(article.Vmname, "-")
			if s[0] == strconv.Itoa(groupid) {
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
	}

	for _, a := range d {
		if err := stream.Send(&pb.VMData{Option: &pb.Option{Id: int64(a.ID), Autostart: a.AutoStart, Status: int32(a.Status)}, Vmname: a.Name, Node: int32(a.NodeID), Vcpu: int64(a.CPU), Vmem: int64(a.Mem), Vnet: a.Net}); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) GetAllVM(base *pb.Base, stream pb.Grpc_GetAllVMServer) error {
	log.Println("----GetAllVM----")
	log.Println("Receive AuthUser  : " + base.GetUser() + ", AuthPass: " + base.GetPass())
	log.Println("Receive Token     : " + base.GetToken())

	if data.AdminUserCertification(base.GetUser(), base.GetPass(), base.GetToken()) == false {
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
			fmt.Println(err)
		}
		for {
			article, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println(err)
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
	log.Println("Receive AuthUser  : " + in.Base.GetUser() + ", AuthPass: " + in.Base.GetPass())
	log.Println("Receive Token     : " + in.Base.GetToken())

	nodeId := in.GetId() / 1000
	vmId := in.GetId() - (1000 * nodeId)

	user := in.Base.GetUser()
	pass := in.Base.GetPass()

	address, result := data.StandardUserVMCertification(&data.UserCertData{
		User:   user,
		Pass:   pass,
		Token:  in.Base.GetToken(),
		VMID:   int(vmId),
		NodeID: int(nodeId),
	})
	if result == false {
		return &pb.Result{Status: false, Info: "Auth Failed!!"}, nil
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ShutdownVM(ctx, &pb.VMID{Id: vmId})
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	return &pb.Result{Status: r.GetStatus(), Info: r.GetInfo()}, nil
}

func (s *server) ResetVM(ctx context.Context, in *pb.VMID) (*pb.Result, error) {
	log.Println("----RebootVM----")
	log.Println("Receive VMID  : ", in.GetId())
	log.Println("Receive AuthUser  : " + in.Base.GetUser() + ", AuthPass: " + in.Base.GetPass())
	log.Println("Receive Token     : " + in.Base.GetToken())

	nodeId := in.GetId() / 1000
	vmId := in.GetId() - (1000 * nodeId)

	user := in.Base.GetUser()
	pass := in.Base.GetPass()

	address, result := data.StandardUserVMCertification(&data.UserCertData{
		User:   user,
		Pass:   pass,
		Token:  in.Base.GetToken(),
		VMID:   int(vmId),
		NodeID: int(nodeId),
	})
	if result == false {
		return &pb.Result{Status: false, Info: "Auth Failed!!"}, nil
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ResetVM(ctx, &pb.VMID{Id: vmId})
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	return &pb.Result{Status: r.GetStatus(), Info: r.GetInfo()}, nil
}

func (s *server) PauseVM(ctx context.Context, in *pb.VMID) (*pb.Result, error) {
	log.Println("----PauseVM----")
	log.Println("Receive VMID  : ", in.GetId())
	log.Println("Receive AuthUser  : " + in.Base.GetUser() + ", AuthPass: " + in.Base.GetPass())
	log.Println("Receive Token     : " + in.Base.GetToken())

	nodeId := in.GetId() / 1000
	vmId := in.GetId() - (1000 * nodeId)

	user := in.Base.GetUser()
	pass := in.Base.GetPass()

	address, result := data.StandardUserVMCertification(&data.UserCertData{
		User:   user,
		Pass:   pass,
		Token:  in.Base.GetToken(),
		VMID:   int(vmId),
		NodeID: int(nodeId),
	})
	if result == false {
		return &pb.Result{Status: false, Info: "Auth Failed!!"}, nil
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.PauseVM(ctx, &pb.VMID{Id: vmId})
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	return &pb.Result{Status: r.GetStatus(), Info: r.GetInfo()}, nil
}

func (s *server) ResumeVM(ctx context.Context, in *pb.VMID) (*pb.Result, error) {
	fmt.Println("-----ResumeVM-----")
	log.Println("Receive VMID  : ", in.GetId())
	log.Println("Receive AuthUser  : " + in.Base.GetUser() + ", AuthPass: " + in.Base.GetPass())
	log.Println("Receive Token     : " + in.Base.GetToken())

	nodeId := in.GetId() / 1000
	vmId := in.GetId() - (1000 * nodeId)

	user := in.Base.GetUser()
	pass := in.Base.GetPass()

	address, result := data.StandardUserVMCertification(&data.UserCertData{
		User:   user,
		Pass:   pass,
		Token:  in.Base.GetToken(),
		VMID:   int(vmId),
		NodeID: int(nodeId),
	})
	if result == false {
		return &pb.Result{Status: false, Info: "Auth Failed!!"}, nil
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ResumeVM(ctx, &pb.VMID{Id: vmId})
	if err != nil {
		fmt.Printf("Not connect; ")
		fmt.Println(err)
	}
	return &pb.Result{Status: r.GetStatus(), Info: r.GetInfo()}, nil
}
