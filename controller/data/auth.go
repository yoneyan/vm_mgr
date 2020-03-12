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

type UserCertData struct {
	User    string
	Pass    string
	Group   string
	Token   string
	VMID    int
	GroupID int
	NodeID  int
}

//Token Authentication
func TokenCertification(token string) (int, []int, bool) {
	go DeleteExpiredToken()
	data, result := db.GetDBToken(token)
	if result == false {
		return 0, nil, false
	}
	if VerifyTime(data.Endtime) == false {
		return 1, nil, false
	}
	return data.Userid, nil, true
}

//User Certification Tool is testing now !!
func AdminUserCertification(name, pass, token string) bool {
	if token != "" {
		tokendata, result := db.GetDBToken(token)
		if result == false {
			fmt.Println("Error: Token Get Error")
			fmt.Println("Certification NG!! (Administrator)")
			return false
		}

		userdata, result := db.GetDBUser(tokendata.Userid)
		if result == false {
			fmt.Println("Error: Token Get Error")
			fmt.Println("Certification NG!! (Administrator)")
			return false
		}
		name = userdata.Name
	} else {
		if db.PassAuthDBUser(name, pass) == false {
			fmt.Println("Certification NG!! (Administrator)")
			return false
		}
	}
	if SearchGroupUser(name, "admin", 0) {
		fmt.Println("Certification OK!! (Administrator)")
		return true
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

func GroupAdminCertification(name, pass, token, group string) bool {
	if token != "" {
		tokendata, result := db.GetDBToken(token)
		if result == false {
			fmt.Println("Error: Token Get Error 1")
			fmt.Println("Certification NG!! (GroupAdmin)")
			return false
		}

		userdata, result := db.GetDBUser(tokendata.Userid)
		if result == false {
			fmt.Println("Error: Token Get Error 2")
			fmt.Println("Certification NG!! (GroupAdmin)")
			return false
		}
		name = userdata.Name

	} else {
		if db.PassAuthDBUser(name, pass) == false {
			fmt.Println("Certification NG!! (GroupAdmin)")
			return false
		}
	}
	if SearchGroupUser(name, group, 0) {
		fmt.Println("Certification OK!! (GroupAdmin)")
		return true
	}

	fmt.Println("Certification NG!! (GroupAdmin)")
	return false
}

func GroupUserCertification(name, pass, token, group string) bool {
	if token != "" {
		tokendata, result := db.GetDBToken(token)
		if result == false {
			fmt.Println("Error: Token Get Error 1")
			fmt.Println("Certification NG!! (GroupUser)")
			return false
		}

		userdata, result := db.GetDBUser(tokendata.Userid)
		if result == false {
			fmt.Println("Error: Token Get Error 2")
			fmt.Println("Certification NG!! (GroupUser)")
			return false
		}
		name = userdata.Name

	} else {
		if db.PassAuthDBUser(name, pass) == false {
			fmt.Println("Certification NG!! (GroupUser)")
			return false
		}
	}
	if SearchGroupUser(name, group, 1) {
		fmt.Println("Certification OK!! (GroupUser)")
		return true
	}

	fmt.Println("Certification NG!! (GroupUser)")
	return false
}

//VM Certification
//
// 10-0-testVM
// VM verify group
func VMCertification(vmid, groupid int, address string) (string, bool) {
	if vmid < 1 {
		fmt.Println("Value False" + strconv.Itoa(vmid))
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

func SuperUserCertification(d *UserCertData) bool {
	if AdminUserCertification(d.User, d.Pass, d.Token) == false {
		if GroupAdminCertification(d.User, d.Pass, d.Token, d.Group) == false {
			return false
		}
	}
	return true
}

func StandardUserCertification(d *UserCertData) bool {
	if AdminUserCertification(d.User, d.Pass, d.Token) == false {
		if GroupAdminCertification(d.User, d.Pass, d.Token, d.Group) == false {
			if GroupUserCertification(d.User, d.Pass, d.Token, d.Group) == false {
				return false
			}
		}
	}
	return true
}

//Only VM Certification

func SuperUserVMCertification(d *UserCertData) (string, bool) {
	g, result := db.GetDBGroupID(d.Group)
	if result == false {
		fmt.Println("Error: Get NodeData Error")
		return "", false
	}
	n, result := db.GetDBNodeID(d.NodeID)
	if result == false {
		fmt.Println("Error: Get NodeData Error")
		return "", false
	}
	if AdminUserCertification(d.User, d.Pass, d.Token) == false {
		if GroupAdminCertification(d.User, d.Pass, d.Token, d.Group) == false {
			return "", false
		} else {
			info, result := VMCertification(d.VMID, g, n.IP)
			if result == false {
				fmt.Println("Error: " + info)
				return "", false
			}
		}
	}
	return n.IP, true
}

func StandardUserVMCertification(d *UserCertData) (string, bool) {
	g, result := db.GetDBGroupID(d.Group)
	if result == false {
		fmt.Println("Error: Get NodeData Error")
		return "", false
	}
	n, result := db.GetDBNodeID(d.NodeID)
	if result == false {
		fmt.Println("Error: Get NodeData Error")
		return "", false
	}
	if AdminUserCertification(d.User, d.Pass, d.Token) == false {
		if GroupAdminCertification(d.User, d.Pass, d.Token, d.Group) == false {
			if GroupUserCertification(d.User, d.Pass, d.Token, d.Group) == false {
				return "", false
			} else {
				info, result := VMCertification(d.VMID, g, n.IP)
				if result == false {
					fmt.Println("Error: " + info)
					return "", false
				}
			}
		}
	}
	return n.IP, true
}
