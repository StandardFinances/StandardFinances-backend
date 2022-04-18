package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"main/ent"
	"net/http"
	"strconv"
	"time"
)

type sendCodeParams struct {
	PhoneNumber string `json:"phone_number"`
}

type confirmCodeParams struct {
	SecurityToken    string `json:"security_token"`
	ConfirmationCode string `json:"confirmation_code"`
}

func main() {
	router := gin.Default()

	router.GET("/ping", pingEndpoint)
	authy := router.Group("/authy")
	{
		authy.POST("/sendcode", sendCodeEndpoint)
		authy.POST("/confirmcode", confirmCodeEndpoint)
	}

	router.Run()
}

func pingEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func sendCodeEndpoint(c *gin.Context) {
	// Parsing the request body.
	var params sendCodeParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Checking request parameters.
	if params.PhoneNumber == "" {
		c.JSON(400, gin.H{
			"message": "phone_number is required",
		})
		return
	}

	db := getDb()

	phoneInt, _ := strconv.ParseInt(params.PhoneNumber, 10, 64)
	db.User.Create().SetUID(uuid.New().String()).SetPhone(phoneInt).SetRegtime(time.Now()).Save(context.Background())

	c.JSON(200, gin.H{
		"security_token": "example_security_token",
	})
}

func confirmCodeEndpoint(c *gin.Context) {
	var params confirmCodeParams
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if params.ConfirmationCode == "" {
		c.JSON(400, gin.H{
			"message": "confirmation_code is required",
		})
		return
	}
	if params.SecurityToken == "" {
		c.JSON(400, gin.H{
			"message": "security_token is required",
		})
		return
	}
	c.JSON(200, gin.H{
		"access_token": "example_access_token",
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
