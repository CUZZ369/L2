package controller

import (
	"RedRock_01/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(ctx *gin.Context)  {
	res:=service.Register(ctx)
	if res {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": "success",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "此用户名已存在",
		})
	}
}
func Login(ctx *gin.Context) {
	res := service.Login(ctx)
	if res {
		ctx.SetCookie("username", ctx.PostForm("username"), 240, "/", "127.0.0.1", false, true)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": "欢迎回来" + ctx.PostForm("username"),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "用户名或密码错误！",
		})
	}
}
