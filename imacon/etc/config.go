package etc

import (
	"encoding/json"
	"io/ioutil"
)

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

func configGet() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)
	//for k, v := range config.Image.Path {
	//	fmt.Printf("Slave %d\n", k)
	//fmt.Printf("  weight is %d\n", v.type)
	//fmt.Printf("  ip is %s\n", v.Ip)
	//}
	//fmt.Printf("DB Username is :%s\n", config.Db.User)
	//fmt.Printf("DB Password is :%s\n", config.Db.Pass)
}
