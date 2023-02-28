package main

import (
	"main/common"
	"main/config"
)

func start() {
	common.InitDB()
	common.Include()
	app := common.InitRouters()
	conf := config.GetConfigData()
	app.Run(conf.App.Host + ":" + conf.App.Port)
}

func main() {
	start()
}
