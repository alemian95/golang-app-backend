package router

import (
	"golang-app/app/controllers/auth_controller"
	"golang-app/app/controllers/index_controller"
	"golang-app/app/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	index_controller.RegisterRoutes(r)
	auth_controller.RegisterRoutes(r)
	user_controller.RegisterRoutes(r)
}
