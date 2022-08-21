package main

import (
	"todolist/conf"
	"todolist/model"
)

func main() {
	conf.InitConfig()
	model.InitDB()
}
