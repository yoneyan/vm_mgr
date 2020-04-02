package etc

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Controller ControllerData `json:controller`
	Web        WebData        `json:web`
}

type WebData struct {
	Https  bool   `json:"https"`
	Domain string `json:"Domain"`
}

type ControllerData struct {
	IP string `json:"ip"`
}

func GetControllerIP() string {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)
	return config.Controller.IP
}

func GetDomain() (bool, string) {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)
	return config.Web.Https, config.Web.Domain
}
