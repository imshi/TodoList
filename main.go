package main

import (
	"todolist/conf"
	"todolist/model"
	"todolist/router"

	"github.com/spf13/viper"
)

func main() {
	conf.InitConfig()
	model.InitDB()

	r := router.NewRouter()
	_ = r.Run(viper.GetString("service.port"))
}

func 