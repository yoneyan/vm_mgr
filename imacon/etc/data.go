package etc

import (
	"fmt"
	"github.com/google/uuid"
	"os"
)

func GeneratePath(imagetype int, filename string) string {
	path := GetImagePath(imagetype) + "/" + filename
	fmt.Println("Path: " + path)
	return path
}

func FileSize(path string) int {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer func() {
		file.Close()
	}()
	fileinfo, staterr := file.Stat()
	if staterr != nil {
		fmt.Println(staterr)
		return 0
	}
	return int(fileinfo.Size())
}

func GenerateUUID() string {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return "0"
	}
	uu := u.String()
	return uu
}
