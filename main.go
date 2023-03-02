package main

import (
	"main/api/user"
	"main/common"
	"main/config"
)

func start() {
	common.InitDB()
	common.Include(
		user.Routers,
	)
	app := common.InitRouters()
	conf := config.GetConfigData()
	app.Run(conf.App.Host + ":" + conf.App.Port)
}

func main() {
	start()
}
