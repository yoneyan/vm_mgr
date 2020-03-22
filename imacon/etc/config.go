package etc

import (
	"encoding/json"
	"io/ioutil"
)

var ConfigData struct {
	ISOPath   string
	ImagePath string
	KeyPath   string
}

type Config struct {
	Image ImageData `json:image`
	Node  NodeData  `json:"node"`
}

type ImageData struct {
	Path []DataPath `json:"path"`
}

type DataPath struct {
	Type int    `json:"type"`
	Path string `json:"path"`
}

type NodeData struct {
	Key string `json:"key"`
}

func GetISOImage(id int) string {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)
	for _, v := range config.Image.Path {
		if v.Type == 0 {
			return v.Path
		}
	}
	return ""
}

func GetDiskImage(id int) string {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)
	for _, v := range config.Image.Path {
		if v.Type == 1 {
			return v.Path
		}
	}
	return ""
}
