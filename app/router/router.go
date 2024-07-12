package router

import (
	"alessandromian.dev/golang-app/app/controllers/user_controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	user_controller.RegisterRoutes(r)
}
