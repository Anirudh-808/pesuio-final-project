package auth

import (
	"github.com/anuragrao04/pesuio-final-project/database"
	"github.com/anuragrao04/pesuio-final-project/models"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {

	var request models.SignUpRequest
	err := c.BindJSON(&request)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid input",
		})
	}

	// implement
	value, _ := database.CreateUser("credentialsDatabase", request.Username, request.Password)

	c.JSON(200, gin.H{
		"message": value,
	})
}
