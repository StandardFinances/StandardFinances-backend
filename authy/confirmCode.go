package authy

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ConfirmCodeEndpoint(c *gin.Context) {
	var params ConfirmCodeParams
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

type ConfirmCodeParams struct {
	SecurityToken    string `json:"security_token"`
	ConfirmationCode string `json:"confirmation_code"`
}
