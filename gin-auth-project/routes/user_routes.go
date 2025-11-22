package routes

import (
	"gin-auth-project/controllers"
	"gin-auth-project/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {

	auth := r.Group("/users")
	auth.Use(middleware.AuthMiddleware())
	auth.GET("/details", controllers.GetUsers)
}
