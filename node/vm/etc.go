package vm

import (
	"github.com/yoneyan/vm_mgr/node/etc"
	"log"
)

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
