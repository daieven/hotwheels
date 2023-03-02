package user

import (
	"github.com/gin-gonic/gin"
	"main/service"
)

func Routers(e *gin.Engine) {
	e.GET("/login", service.Login)
	//e.GET("/comment", commentHandler)
}
