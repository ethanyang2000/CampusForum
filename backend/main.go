package main

import (
	"forum/router"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	router.Register(engine)
	engine.Run("localhost:8080")
}
