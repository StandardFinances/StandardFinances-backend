package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	authyHandlers "main/authy"
)

func main() {
	router := gin.Default()

	authy := router.Group("/authy")
	{
		authy.POST("/sendcode", authyHandlers.SendCodeEndpoint)
		authy.POST("/confirmcode", authyHandlers.ConfirmCodeEndpoint)
	}

	router.Run(":80")
}

func pingEndpoint(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
