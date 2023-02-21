package main

import (
	"github.com/gin-gonic/gin"
	"main/config"
)

func main() {

	app := gin.Default()
	app.GET("/ping", func(c *gin.Context) {
		//输出json结果给调用方
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	conf := config.GetConfigData()
	app.Run(conf.App.Host + ":" + conf.App.Port)
}
