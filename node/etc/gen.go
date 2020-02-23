package etc

const (
	sockpath = "/kvm/socket"
)

func SocketConnectionPath(socketfile string) string {
	socket := "unix-connect:" + sockpath + "/" + socketfile + ".sock"
	return socket
}

func GeneratePath(path, name string) string {
	return path + "/" + name
}
