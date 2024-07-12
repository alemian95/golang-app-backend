package user_controller

import (
	"strconv"

	user_model "alessandromian.dev/golang-app/app/models/user_model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Database *gorm.DB

func RegisterDatabase(db *gorm.DB) {
	Database = db
}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/users", all)
	r.GET("/users/:id", read)
	r.POST("/users", create)
	r.PATCH("/users/:id", update)
	r.DELETE("/users/:id", delete)
}

func all(c *gin.Context) {
	var users = user_model.All(Database)
	c.JSON(200, users)
}

func read(c *gin.Context) {
	user_id, _ := strconv.ParseUint(c.Params[0].Value, 10, 8)
	var user, err = user_model.Find(Database, user_id)
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}
	c.JSON(200, user)
}

func create(c *gin.Context) {
	var user user_model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}
	if err := user.Create(Database); err != nil {
		c.JSON(500, gin.H{"error": "database error"})
		return
	}
	c.JSON(201, user)
}

func update(c *gin.Context) {
	var user user_model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}
	if err := user.Update(Database); err != nil {
		c.JSON(500, gin.H{"error": "database error"})
		return
	}
	c.JSON(201, user)
}

func delete(c *gin.Context) {
	var user user_model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}
	if err := user.Delete(Database); err != nil {
		c.JSON(500, gin.H{"error": "database error"})
		return
	}
	c.JSON(201, user)
}
