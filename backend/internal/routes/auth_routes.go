package routes

import (
	"eCommerce/internal/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {

	reg := router.Group("/api")

	reg.POST("/signup", handlers.SignUP)
	reg.POST("/login", handlers.Login)
}
