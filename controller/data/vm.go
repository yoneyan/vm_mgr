package data

import (
	"context"
	"fmt"
	"github.com/yoneyan/vm_mgr/controller/db"
	"github.com/yoneyan/vm_mgr/controller/server"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"io"
	"strconv"
	"strings"
	"time"
)

type SpecData struct {
	cpu     int
	mem     int
	storage int
}

type StorageData struct {
	Path   string
	Size   string
	Result bool
}

//mode 0: Administrator 1: SuperUser
func CheckNodeID(isadmin bool, nodeid int) (string, bool) {
	d, result := db.GetDBNodeID(nodeid)
	fmt.Println(d)
	if result == false {
		return "", false
	}
	if isadmin == false && d.OnlyAdmin == 1 {
		return "", false
	}

	return d.IP, true
}

func CheckMaxSpec(d *pb.VMData) bool {
	data := TotalSpec(d.Base.GetGroup())
	id, r := db.GetDBGroupID(d.Base.GetGroup())
	if r == false {
		fmt.Println("NodeDB Error!!")
		return false
	}
	group, r := db.GetDBGroup(id)
	if r == false {
		fmt.Println("NodeDB Error!!")
		return false
	}
	if group.MaxMem >= data.mem && group.MaxStorage >= data.storage {
		return false
	}
	return true
}

func TotalSpec(group string) SpecData {
	var d []server.VMDataStruct

	groupid, result := db.GetDBGroupID(group)
	if result == false {
		fmt.Println("DB Not Found Group!!")
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
		stream, err := c.GetAllVM(ctx, &pb.Base{})
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

				d = append(d, server.VMDataStruct{
					NodeID:    a.ID,
					ID:        int(article.Option.Id) + (1000 * a.ID),
					Name:      article.Vmname,
					CPU:       int(article.Vcpu),
					Mem:       int(article.Vmem),
					Net:       article.Vnet,
					Storage:   article.Storage,
					AutoStart: article.Option.Autostart,
					Status:    int(article.Option.Status),
				})
			}
		}
	}
	var cpu, mem, storage int

	for _, a := range d {
		s := strings.Split(a.Storage, ",")
		for _, b := range s {
			tmp, err := strconv.Atoi(b)
			if err != nil {
				fmt.Println("Error!! string to int")
			}
			storage = storage + tmp
		}
		cpu = cpu + a.CPU
		mem = mem + a.Mem
	}
	return SpecData{cpu: cpu, mem: mem, storage: storage}
}

func GenerateDiskPath(d *pb.VMData) StorageData {
	var size, path string
	var tmpsize, tmppath []string

	data := strings.Split(d.Storage, ",")
	if len(data) == 1 {
		path = GetNodePath(int(d.Node), 1)
		size = d.Storage
	} else {
		for i, a := range data {
			path, err := strconv.Atoi(a)
			if err != nil {
				fmt.Println("string to int false...")
				return StorageData{Result: false}
			}
			if i%2 == 0 {
				tmppath = append(tmppath, "1")
				tmppath = append(tmppath, GetNodePath(int(d.Node), path))
			} else {
				tmpsize = append(tmpsize, a)
			}
		}
		path = strings.Join(tmppath, ",")
		size = strings.Join(tmpsize, ",")
	}
	return StorageData{
		Path: path,
		Size: size,
	}
}

func GetNodePath(node, path int) string {
	var j int

	d, result := db.GetDBNodeID(node)
	if result == false {
		fmt.Println("GetNodeDB Error!!!")
	}
	data := strings.Split(d.Path, ",")
	for i, a := range data {
		if a == strconv.Itoa(path) {
			j = i + 1
		}
		if i == j {
			return a
		}
	}
	return ""
}
