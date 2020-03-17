package sftp

//type 0:ISO 1: DiskImage
type FileData struct {
	Type       int
	Name       string
	Path       string
	LocalPath  string
	RemotePath string
	IsLock     bool
}

type SSHInfo struct {
	IP      string
	Port    string
	User    string
	KeyPath string
}
