package vm

import (
	"fmt"
	"github.com/mattn/go-pipeline"
	"log"
)

func VMStop(name string) error {
	//ps axf | grep test|grep qemu  | grep -v grep | awk '{print "kill -9 " $1}' | sudo sh
	out, err := pipeline.Output(
		[]string{"ps", "axf"},
		[]string{"grep", name},
		[]string{"grep", "qemu"},
		[]string{"grep", "-v", "grep"},
		[]string{"awk", "{print \"kill -9 \" $1}"},
		[]string{"sudo", "sh"},
	)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println(string(out))
	return nil
}
