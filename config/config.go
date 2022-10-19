package config

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	AppName    string `json:"app_name"`
	Port       string `json:"port"`
	StaticPath string `json:"static_path"`
	Mode       string `json:"mode"`
}

func InitAppConfig() *AppConfig {
	//打开文件
	file, err := os.Open("./config.json")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	//读取数据
	//jsonData, err := ioutil.ReadAll()
	decoder := json.NewDecoder(file)
	con := AppConfig{}
	err = decoder.Decode(&con)
	if err != nil {
		panic(err.Error())
	}
	return &con
}
