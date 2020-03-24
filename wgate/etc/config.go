package etc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Server ServerData `json:server`
	Path   PathData   `json:path`
}

type ServerData struct {
	IP string `json:"ip"`
}
type PathData struct {
	noVnc     string `json:"novnc"`
	DashBoard string `json:"dashboard"`
}

func GetServerIP() string {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)
	return config.Server.IP
}

func GetnovncPath() string {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)
	fmt.Println(config.Path.noVnc)
	return config.Path.noVnc
}

func GetDashBoardPath() string {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)
	return config.Path.DashBoard
}
