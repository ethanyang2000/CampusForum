package router

import (
	"forum/controller"

	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	engine.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userGroup := engine.Group("user")
	rootGroup := engine.Group("api")

	rootGroup.POST("/login", controller.LogIn)
	rootGroup.POST("/register", controller.Register)

	userGroup.Use(Auth())
	userGroup.POST("/changepswd", controller.ChangePswd)
	userGroup.POST("/changename", controller.ChangeName)
}
