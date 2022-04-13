package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/authy.sendcode", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"security_token": "example_security_token",
		})
	})

	r.GET("/authy.confirmcode", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"confirmation_code": "example_confirmation_code",
			"security_token":    "example_security_token",
		})
	})
	r.Run()
}
