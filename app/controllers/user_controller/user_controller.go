package user_controller

import (
	"strconv"

	"alessandromian.dev/golang-app/app/models/database"
	"alessandromian.dev/golang-app/app/models/user_model"
	"alessandromian.dev/golang-app/app/router/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	public := r.Group("/users")
	{
		public.GET("", all)
		public.GET(":id", read)
		public.POST("", create)
		public.PATCH(":id", update)
		public.DELETE(":id", delete)
	}

	protected := r.Group("/users")
	protected.Use(middlewares.Auth())
	{
		// protected.GET("", all)
	}
}

func all(c *gin.Context) {
	var users = user_model.All(database.Conn())
	c.JSON(200, users)
}

func read(c *gin.Context) {
	user_id, _ := strconv.ParseUint(c.Params[0].Value, 10, 8)
	var user, err = user_model.Find(database.Conn(), user_id)
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
	if err := user.Create(database.Conn()); err != nil {
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
	if err := user.Update(database.Conn()); err != nil {
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
	if err := user.Delete(database.Conn()); err != nil {
		c.JSON(500, gin.H{"error": "database error"})
		return
	}
	c.JSON(201, user)
}
