package auth_controller

import (
	"net/http"

	"alessandromian.dev/golang-app/app/models/user_model"
	"alessandromian.dev/golang-app/app/utils/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	public := r.Group("/auth")
	{
		// public.GET("/csrf", getCsrf)
		public.GET("/check", check)
		public.POST("/login", login)
		public.POST("/logout", logout)
	}
}

// func getCsrf(c *gin.Context) {
// 	token := csrf.Token(c.Request)

// 	c.Header("X-CSRF-Token", token)

// 	c.SetCookie("X-CSRF-Token", token, 3600, "/", "localhost", false, true)

// 	c.JSON(200, gin.H{"csrf": token})
// }

func check(c *gin.Context) {

	token, err := c.Cookie("gosession")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	user := auth.GetUserBySession(token)

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	c.JSON(200, user)
}

func login(c *gin.Context) {
	var request auth.LoginRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if request.Email == "admin@example.com" && request.Password == "example" {
		user, err := user_model.FindByEmail(request.Email)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		token, err := auth.GenerateToken(user.ID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
			return
		}

		c.SetCookie("gosession", token, 7200, "/", "localhost", false, false)

		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
	}
}

func logout(c *gin.Context) {
	c.SetCookie("gosession", "", -1000, "/", "localhost", false, false)
	c.JSON(200, gin.H{"message": "Logged out"})
}
