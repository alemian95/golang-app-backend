package auth_controller

import (
	"net/http"

	"alessandromian.dev/golang-app/app/models/database"
	"alessandromian.dev/golang-app/app/models/user_model"
	"alessandromian.dev/golang-app/app/utils/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/login", login)
}

func login(c *gin.Context) {
	var request auth.LoginRequest

	err := c.ShouldBindJSON(&request)

	print(request.Email)
	print(request.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if request.Email == "admin@example.com" && request.Password == "example" {
		user, err := user_model.FindByEmail(database.Conn(), request.Email)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		token, err := auth.GenerateToken(user.ID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
	}
}
