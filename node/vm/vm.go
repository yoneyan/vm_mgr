package vm

import (
	"fmt"
	"github.com/mattn/go-pipeline"
	"github.com/yoneyan/vm_mgr/node/etc"
	"log"
	"os/exec"
	"strconv"
)

func RunQEMUMonitor(command, socket string) error {
	//Example:
	//echo "system_powerdown" | socat - unix-connect:/var/run/someapp/vm.sock
	//

	out, err := pipeline.Output(
		[]string{"echo", command},
		[]string{"sudo", "socat", "-", socket},
	)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println(string(out))
	return nil
}
func RunQEMUCmd(cmd []string) error {
	fmt.Println("----CommandRun")

	//cmd = append(cmd,"-") //Intel VT-d support enable

	err := exec.Command("qemu-system-x86_64", cmd...).Start()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Command Error!")
		return err
	}
	return nil
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
	}

	fmt.Println(cmd)

	return cmd
}
