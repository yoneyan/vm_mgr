package etc

import (
	"encoding/json"
	"io/ioutil"
)

var ConfigData struct {
	DiskPath  []DiskPath
	ImagePath []ImagePath
	KeyPath   string
}

type Config struct {
	Disk  DiskData  `json:"disk"`
	Image ImageData `json:image`
}

type DiskData struct {
	Path []DiskPath `json:"path"`
}

type DiskPath struct {
	Type   int    `json:"type"`
	Path   string `json:"path"`
	Status bool   `json:"status"`
}

type ImageData struct {
	Path []ImagePath `json:"path"`
}

type ImagePath struct {
	ID   int    `json:"id"`
	Path string `json:"path"`
}

func GetDiskPath(id int) string {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)
	for _, v := range config.Disk.Path {
		if v.Type == id && v.Status {
			return v.Path
		}
	}
	return ""
}

func GetImagePath(id int) string {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)
	for _, v := range config.Image.Path {
		if v.ID == id {
			return v.Path
		}
	}
	return ""
}
