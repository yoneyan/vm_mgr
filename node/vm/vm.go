package vm

import (
	"fmt"
	"github.com/mattn/go-pipeline"
	"github.com/yoneyan/vm_mgr/node/db"
	"github.com/yoneyan/vm_mgr/node/etc"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
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
		fmt.Println(err)
		return err
	}
	fmt.Println(string(out))
	return nil
}
func RunQEMUCmd(command []string) error {
	fmt.Println("----CommandRun")
	//cmd = append(cmd,"-") //Intel VT-d support enable
	cmd := exec.Command("qemu-system-x86_64", command...)

	id, err := db.VMDBGetVMID(command[2])
	if err != nil {
		fmt.Println("DB Read Error")
	}
	db.VMDBStatusUpdate(id, 1)

	//go manage.VMLifeCheck(command[2], cmd)
	go func() {
		cmd.Start()
		fmt.Println("--------------------------------")
		fmt.Println("VMName: "+command[2]+" StartTime: ", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println("--------------------------------")
		cmd.Wait()
		db.VMDBStatusUpdate(id, 0)
		fmt.Println("--------------------------------")
		fmt.Println("VMName: "+command[2]+" EndTime: ", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println("--------------------------------")
	}()
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
	if len(c.Net) != 0 {
		netdata := GenerateNetworkCmd(c.Net)
		fmt.Println("safa")
		//add qemu network command
		for _, a := range netdata {
			cmd = append(cmd, a)
		}
	}

	fmt.Printf("GenerateCommand: ")
	fmt.Println(cmd)

	return cmd
}

//Generate Network Command
func GenerateNetworkCmd(net string) []string {
	data := strings.Split(net, ",")
	fmt.Println(data)
	mode, err := strconv.Atoi(data[0])
	if err != nil {
		log.Fatal(err)
	}
	var bridge []string
	var mac []string
	for i, a := range data {
		if i > 0 {
			if i%2 == 1 {
				bridge = append(bridge, a)
			} else {
				mac = append(mac, a)
			}
		}
	}

	//verify bridge and mac array length
	if len(bridge) != len(mac) {
		fmt.Println("Warning: NetworkCount Error")
	}

	var cmd []string
	///etc/qemu/bridge.conf <- allow br0
	//1 Network
	//-net nic,macaddr=52:54:01:11:22:33 -net bridge,br=br0
	//2 Network
	//-net nic,macaddr=52:54:01:11:22:33 -net bridge,br=br0 -net nic,macaddr=52:54:02:11:22:33 -net bridge,br=br0
	if mode == 0 {
		//default Network
		cmd = append(cmd, "-net")
		cmd = append(cmd, "nic,macaddr="+mac[0])
		cmd = append(cmd, "-net")
		cmd = append(cmd, "bridge,br="+bridge[0])
		for i, _ := range mac {
			if i > 0 {
				cmd = append(cmd, "-net")
				cmd = append(cmd, "nic,macaddr="+mac[i]+",vlan="+strconv.Itoa(i))
				cmd = append(cmd, "-net")
				cmd = append(cmd, "bridge,br="+bridge[i]+",vlan="+strconv.Itoa(i))
			}
		}
	} else if mode == 1 {
		//rtl8139

	}

	fmt.Printf("GenerateNetwork: ")
	fmt.Println(cmd)
	return cmd
}
