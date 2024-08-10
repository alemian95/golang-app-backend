package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang-app/app/utils/auth"
	"golang-app/app/utils/config"
	"golang-app/app/utils/helpers"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Set CORS headers
		c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("ALLOW_ORIGIN"))
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-XSRF-TOKEN, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")

		// Allow OPTIONS requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		// Pass to next middleware in chain
		c.Next()
	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(helpers.GetCurrentDateTimestamp(), " Received request:", c.Request.URL)
		c.Next()
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")

		// Check for Authorization header
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		// Check for valid JWT
		tokenParts := strings.Split(tokenString, " ")

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		tokenString = tokenParts[1]

		// Verify the token
		claims, err := auth.VerifyToken(tokenString)

		// Check for error
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", uint64(claims.Payload["user_id"].(float64)))

		newPayload := helpers.NewAssocArray()
		newPayload["user_id"] = uint64(claims.Payload["user_id"].(float64))

		// Regenerate session
		token, err := auth.GenerateToken(newPayload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
			return
		}
		c.SetCookie("gosession", token, 7200, "/", "localhost", false, false)

		c.Next()
	}
}

func CSRF() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" || c.Request.Method == "PATCH" || c.Request.Method == "DELETE" {
			tokenHeader := c.GetHeader(config.XSRF_header_name)
			tokenCookie, err := c.Cookie(config.XSRF_cookie_name)

			if err != nil {
				c.JSON(419, gin.H{"error": "CSRF token not valid"})
				c.Abort()
				return
			}

			if tokenHeader != tokenCookie {
				c.JSON(419, gin.H{"error": "CSRF token not valid"})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
