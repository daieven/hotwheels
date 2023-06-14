package main

import (
	"main/api/user"
	"main/common"
	"main/config"
	"main/models"
)

func start() {

	dbEngine := common.InitDB()
	tables := []interface{}{
		&models.User{},
	}
	common.InitTable(dbEngine, tables...)

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
