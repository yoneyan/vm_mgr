package data

import pb "github.com/yoneyan/vm_mgr/proto/proto-go"

func CreateVMCheck(d *pb.VMData) bool {
	if d.Vcpu < 0 || d.Vmem < 0 || d.Option.Vnc < 0 {
		return false
	}
	if d.Vmname == "" || d.Option.StoragePath == "" {
		return false
	}
	return true
}
