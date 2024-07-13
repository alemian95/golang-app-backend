package router

import (
	"alessandromian.dev/golang-app/app/controllers/index_controller"
	"alessandromian.dev/golang-app/app/controllers/user_controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	index_controller.RegisterRoutes(r)
	user_controller.RegisterRoutes(r)
}
