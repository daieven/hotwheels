package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")

	// 处理登录信息
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "登录成功",
		"code":    200,
		"data":    "username:" + username + "password" + password,
	})
}
