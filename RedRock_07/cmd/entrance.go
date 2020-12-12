package cmd

import (
	"RedRock_01/controller"
	"github.com/gin-gonic/gin"
)

func Entrance()  {
	router:=gin.Default()
	router.POST("/register",controller.Register)
	router.GET("/login",controller.Login)
	router.Run(":8080")
}
