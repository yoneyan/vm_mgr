package manage

import (
	"fmt"
	"github.com/yoneyan/vm_mgr/node/etc"
	"os/exec"
	"strconv"
)

type Storage struct {
	Path   string
	Name   string
	Format string
	Size   int
}

func RunStorageCmd(cmd []string) {
	out, _ := exec.Command("qemu-img", cmd...).Output()
	fmt.Println(out)
}

//path, name string, format, size int
func CreateStorage(s *Storage) error {
	fmt.Println("----storage create----")
	fmt.Println(s)
	if s.Size < 0 {
		return fmt.Errorf("Wrong storage size !!")
	}

	if s.Format != "qcow2" && s.Format != "raw" {
		return fmt.Errorf("Wrong storage format !!")
	}

	var cmd []string

	//qemu-img create [-f format] filename [size]

	size := strconv.Itoa(s.Size) + "M"
	path := etc.GeneratePath(s.Path, s.Name) + ".img"
	cmd = append(cmd, "create")
	cmd = append(cmd, "-f")
	cmd = append(cmd, s.Format)
	cmd = append(cmd, path)
	cmd = append(cmd, size)

	fmt.Println(cmd)

	RunStorageCmd(cmd)

	out, _ := exec.Command("qemu-img", cmd...).Output()
	fmt.Println(out)

	return nil
}

func DeleteStorage(s *Storage) error {
	var cmd []string

	filepath := etc.GeneratePath(s.Path, s.Name)
	if etc.FileExists(filepath) {
		cmd = append(cmd, "info")
		cmd = append(cmd, filepath+".img")

		return nil
	}
	RunStorageCmd(cmd)

	return fmt.Errorf("File not exits!!")
}

func ResizeStorage(s *Storage, soon bool) error {
	//qemu-img resize filename size

	var cmd []string

	cmd = append(cmd, "qemu-img")
	cmd = append(cmd, "resize")
	cmd = append(cmd, etc.GeneratePath(s.Path, s.Name)+".img")
	cmd = append(cmd, strconv.Itoa(s.Size)+"M")

	RunStorageCmd(cmd)

	return nil
}

func InformationStorage(s *Storage) error {
	//qemu-img info [-f format] filename
	var cmd []string

	cmd = append(cmd, "qemu-img")
	cmd = append(cmd, "info")
	cmd = append(cmd, etc.GeneratePath(s.Path, s.Name))

	RunStorageCmd(cmd)
	return nil
}
