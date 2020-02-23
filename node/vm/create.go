package vm

import "C"
import (
	"fmt"
	"github.com/yoneyan/vm_mgr/node/db"
	"github.com/yoneyan/vm_mgr/node/etc"
	"strconv"
)

type CreateVMInformation struct {
	Name        string
	CPU         int
	Mem         int
	StoragePath string
	CDROM       string
	Net         string
	VNC         int
}

func CreateVMDBProcess(c *CreateVMInformation) {
	dbdata := db.NodeVM{
		Name:        c.Name,
		CPU:         c.CPU,
		Mem:         c.Mem,
		StoragePath: c.StoragePath,
		Net:         c.Net,
		Vnc:         c.VNC,
		Socket:      etc.SocketGenerate(c.Name),
	}
	db.AddDBVM(dbdata)
}

func StartVMProcess(id int) {

}

func CreateGenerateCmd(c *CreateVMInformation) []string {
	var cmd []string

	cmd = append(cmd, "-enable-kvm") //kvm enable
	cmd = append(cmd, "-smp")
	cmd = append(cmd, strconv.Itoa(c.CPU))
	cmd = append(cmd, "-m")
	cmd = append(cmd, strconv.Itoa(c.Mem))
	cmd = append(cmd, "-monitor")
	cmd = append(cmd, etc.SocketGenerate(c.Name))
	cmd = append(cmd, "-boot")
	if c.CDROM != "" {
		cmd = append(cmd, "order=d")
		cmd = append(cmd, "-cdrom")
		cmd = append(cmd, c.CDROM)
	} else {
		cmd = append(cmd, "order=c")
	}
	cmd = append(cmd, "-hda")
	cmd = append(cmd, c.StoragePath)
	cmd = append(cmd, "-vnc")
	cmd = append(cmd, ":"+strconv.Itoa(c.VNC))

	fmt.Println(cmd)

	return cmd
}
