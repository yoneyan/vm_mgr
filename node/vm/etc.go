package vm

import (
	"github.com/yoneyan/vm_mgr/node/etc"
	"log"
)

func VMShutdown(name string) error {
	err := RunQEMUMonitor("system_powerdown", etc.SocketConnectionPath(name))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func VMRestart(name string) error {
	err := RunQEMUMonitor("system_reset", etc.SocketConnectionPath(name))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func VMPause(name string) error {
	err := RunQEMUMonitor("system_reset", etc.SocketConnectionPath(name))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func VMResume(name string) error {
	err := RunQEMUMonitor("system_reset", etc.SocketConnectionPath(name))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
