package controller

import (
	"forum/model"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	obj := struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		Pswd  string `json:"pswd"`
	}{}

	if err := c.BindJSON(&obj); err != nil {
		c.JSON(500, err.Error())
		return
	}
	if err := model.Register(obj.Email, obj.Name, obj.Pswd); err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"message": "registered succeed!",
	})

}

func LogIn(c *gin.Context) {
	if cookie, err := model.LogIn(c); err == nil {
		c.SetCookie("campus_forum", cookie, 0, "/", "127.0.0.1", true, true)
		c.JSON(200, gin.H{
			"message": "log in succeed",
		})
	} else {
		c.JSON(401, gin.H{
			"message": err.Error(),
		})
	}

}
