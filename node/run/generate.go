package run

const (
	sockpath = "/kvm/socket"
)

func SocketConnectionPath(socketfile string) string {
	socket := "unix-connect:" + sockpath + "/" + socketfile + ".sock"
	return socket
}
