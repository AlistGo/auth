package main

import (
	"auth-api/app"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})
	app.Setup(r.Group("/api"))
	r.Run(":8080")
}
