package conf

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode     string
	HttpPort    string
	RedisAddr   string
	RedisPwd    string
	RedisDbName string
	Db          string
	DbHost      string
	DbPort      string
	DbUser      string
	DbName      string
	DbPasswd    string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径！")
	}

	LoadServer(file)
}

func LoadServer(file *.ini.File) {
}