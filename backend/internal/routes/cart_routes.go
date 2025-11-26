package routes

import (
	"eCommerce/internal/handlers"
	"eCommerce/internal/middleware"

	"github.com/gin-gonic/gin"
)

var cartHandler = handlers.CartHandler{}

func CartRoutes(router *gin.Engine) {
	cart := router.Group("/api/v1/cart")
	cart.Use(middleware.AuthMiddleware())

	cart.POST("/add", cartHandler.AddToCart)
}
