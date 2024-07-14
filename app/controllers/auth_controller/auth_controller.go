package auth_controller

import (
	"net/http"

	"alessandromian.dev/golang-app/app/models/user_model"
	"alessandromian.dev/golang-app/app/utils/auth"
	"github.com/gin-gonic/gin"
)

/**
 * Registering routes for auth controller
 *
 * @param c *gin.Engine
 */
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

/**
 * Check if user is logged in or not
 *
 * @param c *gin.Context
 */
func check(c *gin.Context) {

	// get token from cookie
	token, err := c.Cookie("gosession")

	// check if token exists
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	// get user from token
	user, err := auth.GetUserBySession(token)

	// check if user exists in db
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	// return user
	c.JSON(200, user)
}

/**
 * Login user and set session cookie
 *
 * @param c *gin.Context
 */
func login(c *gin.Context) {
	var request auth.LoginRequest

	// validate request
	err := c.ShouldBindJSON(&request)

	// check if request is valid
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// check if credentials are correct (now user and password are hardcoded)
	if request.Email == "admin@example.com" && request.Password == "example" {
		user, err := user_model.FindByEmail(request.Email)

		// check if user exists in db
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		// set session cookie with token
		token, err := auth.GenerateToken(user.ID)

		// check if token is generated correctly
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
			return
		}

		// set cookie with token
		c.SetCookie("gosession", token, 7200, "/", "localhost", false, false)

		// return token
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		// return error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
	}
}

/**
 * Logout user and delete session cookie
 *
 * @param c *gin.Context
 */
func logout(c *gin.Context) {
	c.SetCookie("gosession", "", -1000, "/", "localhost", false, false)
	c.JSON(200, gin.H{"message": "Logged out"})
}
