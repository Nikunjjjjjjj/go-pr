package routes

import (
	"eCommerce/internal/handlers"
	"eCommerce/internal/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoute(r *gin.Engine) {
	order := r.Group("api/v1/order")
	order.Use(middleware.AuthMiddleware())
	order.POST("/new", handlers.CreateOrder)
	order.GET("/", handlers.GetAllOrder)
}
