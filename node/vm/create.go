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
		Status:      0,
	}
	db.AddDBVM(dbdata)
}

func StartVMProcess(id int) {

}

func CreateGenerateCmd(c *CreateVMInformation) []string {
	var cmd []string

	begin := []string{"-enable-kvm", "-name", c.Name, "-smp", strconv.Itoa(c.CPU), "-m", strconv.Itoa(c.Mem), "-monitor", etc.SocketGenerate(c.Name), "-hda", c.StoragePath + "/" + c.Name + ".img", "-vnc", ":" + strconv.Itoa(c.VNC)}
	cmdarray := []string{"-boot"}

	cmd = append(cmd, begin...)
	if c.CDROM != "" {
		cmd = append(cmd, cmdarray[0])
		cmd = append(cmd, "order=d")
		cmd = append(cmd, "-cdrom")
		cmd = append(cmd, c.CDROM)
	} else {
		cmd = append(cmd, "order=c")
	}

	fmt.Println(cmd)

	return cmd
}
