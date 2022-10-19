package config

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	AppName    string   `json:"app_name"`
	Port       string   `json:"port"`
	StaticPath string   `json:"static_path"`
	Mode       string   `json:"mode"`
	Database   Database `json:"data_base"`
	Redis      Redis    `json:"redis"`
}

// "drive":"mysql",
//
//	"host": "127.0.0.1",
//	"port": "3306",
//	"user": "root",
//	"pwd": "root",
//	"database": "cmsdata"
type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Drive    string `json:"drive"`
}

//	"redis": {
//	   "network": "tcp",
//	   "port": "6379",
//	   "addr": "127.0.0.1",
//	   "password": "",
//	   "prefix": "cms_"
//	 }
type Redis struct {
	NetWork  string `json:"netWork"`
	Port     string `json:"port"`
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Prefix   string `json:"prefix"`
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
