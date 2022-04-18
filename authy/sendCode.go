package authy

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	//db := getDb()

	//phoneInt, _ := strconv.ParseInt(params.PhoneNumber, 10, 64)
	//db.User.Create().SetUID(uuid.New().String()).SetPhone(phoneInt).SetRegtime(time.Now()).Save(context.Background())

	c.JSON(200, gin.H{
		"security_token": "example_security_token",
	})
}
