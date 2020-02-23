package vm

import (
	"github.com/yoneyan/vm_mgr/node/run"
	"log"
)

func Start() {

}

func Stop() {

}

func Shutdown(sockname string) error {
	err := run.RunQEMUMonitor("system_powerdown", run.SocketConnectionPath(sockname))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func Restart(sockname string) error {
	err := run.RunQEMUMonitor("system_reset", run.SocketConnectionPath(sockname))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
