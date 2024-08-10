package auth_controller

import (
	"net/http"

	"golang-app/app/models/user_model"
	"golang-app/app/utils/auth"
	"golang-app/app/utils/config"
	"golang-app/app/utils/helpers"

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
		public.GET("/csrf", getCsrf)
		public.GET("/check", check)
		public.POST("/login", login)
		public.POST("/logout", logout)
	}
}

func getCsrf(c *gin.Context) {
	token := auth.GenerateRandomToken()
	c.SetCookie(config.XSRF_cookie_name, token, 60, "/", "localhost", false, false)
	c.JSON(http.StatusNoContent, gin.H{})
}

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
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	requested_user, err := user_model.FindByEmail(request.Email)

	if err != nil {
		// return invalid credentials, we don't want to give information about the user
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid credentials"})
		return
	}

	// check if credentials are correct (now user and password are hardcoded)
	if auth.CheckPasswordHash(request.Password, requested_user.Password) {

		payload := helpers.NewAssocArray()
		payload["user_id"] = requested_user.ID

		// set session cookie with token
		token, err := auth.GenerateToken(payload)

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
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid credentials"})
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
