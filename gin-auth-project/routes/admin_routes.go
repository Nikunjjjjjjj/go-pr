package routes

import (
	"gin-auth-project/controllers"
	"gin-auth-project/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {
	admin := r.Group("/admin")

	// Apply both middlewares
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.AdminOnly())

	admin.GET("/users", controllers.GetUsers)
	//admin.DELETE("/user/:id", controllers.DeleteUser)
}
