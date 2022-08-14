package conf

import "gopkg.in/ini.v1"

// const (
// 	AppMode string
// 	HttpPort string
// 	RedisAddr string
// 	RedisPwd string
// 	RedisDbName string
// 	Db string
// 	DbHost string
// 	DbPort string
// 	DbUser string
// 	DbName string
// 	DbPasswd string
// )

func Init() {
	file, err := ini.Load("./config.ini")
}
