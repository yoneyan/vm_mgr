package run

import (
	"fmt"
	"github.com/mattn/go-pipeline"
	"log"
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
