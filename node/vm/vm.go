package vm

import (
	"fmt"
	"github.com/mattn/go-pipeline"
	"github.com/yoneyan/vm_mgr/node/etc"
	"log"
	"os/exec"
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
func RunQEMUCmd(cmd []string) {
	fmt.Println("----CommandRun")

	//cmd = append(cmd,"-") //Intel VT-d support enable

	out, _ := exec.Command("qemu-system-x86_64", cmd...).Output()
	fmt.Println(out)

}

func Start() {

}

func Stop() {

}

func Shutdown(sockname string) error {
	err := RunQEMUMonitor("system_powerdown", etc.SocketConnectionPath(sockname))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func Restart(sockname string) error {
	err := RunQEMUMonitor("system_reset", etc.SocketConnectionPath(sockname))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
