package authy

import (
	"github.com/gin-gonic/gin"
	"main/db"
	"net/http"
	"strconv"
)

type SendCodeParams struct {
	PhoneNumber string `json:"phone_number"`
}

func SendCodeEndpoint(c *gin.Context) {
	// Parsing the request body.
	var params SendCodeParams
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

	//
	// TODO: Check if the phone number is valid.
	//
	phone, _ := strconv.ParseInt(params.PhoneNumber, 10, 64)

	instance := db.GetInstance()
	defer instance.Close()
	user := db.GetUserByPhone(instance, phone)

	c.JSON(200, gin.H{
		"security_token": "security_token_will_be_here",
		"debug_userid":   &user.UID,
		"debug_phone":    &user.Phone,
	})
}
