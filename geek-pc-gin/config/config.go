package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	DBUser    string
	DBPass    string
	DBName    string
	RedisAddr string
	JWTSecret string
}

var AppConfig *Config

func LoadConfig() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("读取文件失败:", err)
	}

	AppConfig = &Config{}
	if err := json.Unmarshal(data, AppConfig); err != nil {
		log.Fatal("解析Json数据失败:", err)
	}
}
