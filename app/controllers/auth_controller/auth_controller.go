package auth_controller

import (
	"net/http"
	"os"

	"golang-app/app/models/forgot_password_token_model.go"
	"golang-app/app/models/user_model"
	"golang-app/app/utils/auth"
	"golang-app/app/utils/config"
	"golang-app/app/utils/helpers"
	"golang-app/app/utils/mail"

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
		public.POST("/register", register)
		public.POST("/forgot-password", forgotPassword)
		public.GET("/reset-password/:token", getRequestedResetEmail)
		public.POST("/reset-password", resetPassword)
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

/**
 * Register a new user
 */
func register(c *gin.Context) {
	var request auth.RegisterRequest

	// validate request
	err := c.ShouldBindJSON(&request)

	// check if request is valid
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if request.Password != request.PasswordConfirm {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Password confirmation does not match"})
		return
	}

	user_exists := user_model.CheckIfEmailExists(request.Email)

	if user_exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already registered"})
		return
	}

	hash, _ := auth.HashPassword(request.Password)

	user := user_model.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hash,
	}

	user.Create()

	c.JSON(http.StatusNoContent, gin.H{})
}

func forgotPassword(c *gin.Context) {
	var request auth.ForgotPasswordRequest

	// validate request
	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	user, err := user_model.FindByEmail(request.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	str_token := auth.GenerateRandomToken()

	token := forgot_password_token_model.ForgotPasswordToken{
		UserId: user.ID,
		Token:  str_token,
	}

	token.Save()

	link := os.Getenv("ALLOW_ORIGIN") + "/reset-password/" + str_token

	mailBody := "Click the link to reset your password \n" + link

	dialer, message := mail.Prepare("admin@example.com", user.Email, "Password reset link", mailBody)

	mail.Send(dialer, message)

	c.JSON(http.StatusNoContent, gin.H{})
}

func getRequestedResetEmail(c *gin.Context) {

	token_str := c.Params[0].Value

	token, err := forgot_password_token_model.FindByToken(token_str)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Token not found"})
		return
	}

	user, err := user_model.Find(uint64(token.UserId))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": user.Email})
}

func resetPassword(c *gin.Context) {
	var request auth.ResetPasswordRequest

	// validate request
	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	user, err := user_model.FindByEmail(request.Email)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	token, err := forgot_password_token_model.FindByUser(user.ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No forgot password request found"})
		return
	}

	if token.Token != request.Token {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}

	if request.Password != request.PasswordConfirm {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Passwords do not match"})
		return
	}

	hash, _ := auth.HashPassword(request.Password)

	user.Password = hash
	user.Update()
	token.Delete()

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}
