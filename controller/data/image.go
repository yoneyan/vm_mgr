package data

import (
	"context"
	"fmt"
	"github.com/yoneyan/vm_mgr/controller/db"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"time"
)

type ImageStruct struct {
	Address string
	Result  bool
}

func GetImaConHostIP() []string {
	d := db.GetDBAllImaCon()
	var hostip []string

	for _, a := range d {
		hostip = append(hostip, a.IP)
	}
	return hostip
}

func GetImage(name, tag string) ImageStruct {
	for _, a := range db.GetDBAllImaCon() {
		conn, err := grpc.Dial(a.IP, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			fmt.Printf("Not connect; ")
			fmt.Println(err)
		}
		defer conn.Close()

		c := pb.NewGrpcClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.ExistImage(ctx, &pb.ImageData{Name: name, Tag: tag})
		if err != nil {
			fmt.Println(err)
		}
		if r.Result {
			return ImageStruct{Address: a.IP, Result: true}
		}
	}
	return ImageStruct{Result: false}
}

func RegistImage(name, group string) (string, bool) {
	data, info, result := VerifyGroup(name, group)
	if result == false {
		return info, false
	}
	fmt.Println(data)
	d, result := ProcessStringToArray(data.User, name, 0)
	if result == false {
		fmt.Println("Error: User is exists this group")
		return "Error: User is exists this group", false
	}

	if db.ChangeDBGroupUser(data.ID, d) {
		return "ok", true
	}
	return "DBChangeError!!", false
}

func UnRegistImgae(name, group string) (string, bool) {
	data, info, result := VerifyGroup(name, group)
	if result == false {
		return info, false
	}
	fmt.Println(data)
	d, result := ProcessStringToArray(data.User, name, 0)
	if result == false {
		fmt.Println("Error: User is exists this group")
		return "Error: User is exists this group", false
	}

	if db.ChangeDBGroupUser(data.ID, d) {
		return "ok", true
	}
	return "DBChangeError!!", false
}
