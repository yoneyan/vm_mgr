package etc

const (
	sockpath = "/kvm/socket"
)

func SocketConnectionPath(socketfile string) string {
	socket := "unix-connect:" + sockpath + "/" + socketfile + ".sock"
	return socket
}

func SocketPath(socketfile string) string {
	return sockpath + "/" + socketfile + ".sock"
}

func SocketGenerate(socketfile string) string {
	//unix:/tmp/monitor.sock,server,nowait
	return "unix:" + sockpath + "/" + socketfile + ".sock,server,nowait"
}

func GeneratePath(path, name string) string {
	return path + "/" + name
}
