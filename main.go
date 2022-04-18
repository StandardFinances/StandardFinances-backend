package main

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	authyHandlers "main/authy"
	"main/ent"
)

func main() {
	router := gin.Default()

	authy := router.Group("/authy")
	{
		authy.POST("/sendcode", authyHandlers.SendCodeEndpoint)
		authy.POST("/confirmcode", authyHandlers.ConfirmCodeEndpoint)
	}

	router.Run()
}

func pingEndpoint(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func getDb() *ent.Client {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
