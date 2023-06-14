package main

import (
	"main/api/user"
	"main/common"
	"main/config"
	"main/models"
)

func start() {

	dbEngine := common.InitDB()
	//tables := []interface{}{
	//	&models.User{},
	//}
	//common.InitTable(dbEngin, tables)
	dbEngine.AutoMigrate(&models.User{})
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
