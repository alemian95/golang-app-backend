package index_controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Database *gorm.DB

func RegisterDatabase(db *gorm.DB) {
	Database = db
}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", index)
}

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":  "Server Works",
		"database": Database.Name(),
	})
}
