package data

import (
	"context"
	"fmt"
	"github.com/yoneyan/vm_mgr/controller/db"
	pb "github.com/yoneyan/vm_mgr/proto/proto-go"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"strings"
	"time"
)

//User Certification Tool is testing now !!
func AdminUserCertification(name, pass string) bool {
	if db.PassAuthDBUser(name, pass) {
		if SearchGroupUser(name, "admin", 0) {
			fmt.Println("Certification OK!! (Administrator)")
			return true
		}
	}
	fmt.Println("Certification NG!! (Administrator)")
	return false
}

func UserCertification(name, pass string) bool {
	if db.PassAuthDBUser(name, pass) {
		fmt.Println("Certification OK!! (User)")
		return true
	}
	fmt.Println("Certification NG!! (User)")
	return false
}

func GroupAdminCertification(name, pass, group string) bool {
	if db.PassAuthDBUser(name, pass) && SearchGroupUser(name, group, 0) {
		fmt.Println("Certification OK!! (GroupAdmin)")
		return true
	}
	fmt.Println("Certification NG!! (GroupAdmin)")
	return false
}

func GroupUserCertification(name, pass, group string) bool {
	if db.PassAuthDBUser(name, pass) && SearchGroupUser(name, group, 1) {
		fmt.Println("Certification OK!! (GroupUser)")
		return true
	}
	fmt.Println("Certification NG!! (GroupUser)")
	return false
}

//VM Certification
func VMCertification(vmid, groupid int, address string) (string, bool) {
	if vmid < 1 {
		fmt.Println("Value False")
		fmt.Printf("Debug: value is ")
		return "Error VMID", false
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Not connect; %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetVM(ctx, &pb.VMID{Id: int64(vmid)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	if r.GetVmname() == "" {
		return "None", false
	}
	d := strings.Split(r.GetVmname(), "-")
	if d[0] == strconv.Itoa(groupid) {
		fmt.Println("Certification OK!!")
		return "ok", true
	}
	return "None", false
}
