package util

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

var (
	instance *Config
	conf     = flag.String("conf", "../etc/config.json", "这是描述")

	//Config *Config
)

type Config struct {
	BasePath string `json:"base_path"`
	DataPath string `json:"date_path"`
}

func init() {
	flag.Parse()
	//fmt.Println()
	NewUConfigWithFile("/Users/rookie/Desktop/project/go/src/system/etc/config.json")
}

func NewUConfigWithFile(name string)  {
	if instance == nil {
		c := &Config{}

		bytes,err := ioutil.ReadFile(name)
		if err != nil {
			fmt.Println("读取配置文件失败：", name)
			return
		}
		//fmt.Println(string(bytes),bytes, c)
		err = json.Unmarshal(bytes, c)
		if err!= nil {
			fmt.Println("解析json文件失败")
			return
		}

		instance =  c
	}
}

func  GetConfig() *Config {
	return instance
}