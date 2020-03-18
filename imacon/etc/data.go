package etc

import (
	"fmt"
	"os"
)

func GeneratePath(imagetype int, filename string) string {
	var path string
	if imagetype == 0 {
		path = ConfigData.ISOPath + "/" + filename
	} else if imagetype == 1 {
		path = ConfigData.ImagePath + "/" + filename
	}
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
