package router

import (
	"forum/model"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if cookie, err := ctx.Cookie("campus_forum"); err != nil {
			ctx.JSON(500, gin.H{
				"message": "unauthorized user!",
			})
		} else {
			if _, ok := model.CheckCookie(cookie); ok {
				return
			} else {
				ctx.JSON(500, gin.H{
					"message": "invalid cookie",
				})
			}
		}
	}
}
