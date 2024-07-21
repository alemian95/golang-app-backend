package index_controller

import (
	"golang-app/app/models/database"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", index)
}

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":  "Server Works",
		"database": database.Conn().Name(),
	})
}
