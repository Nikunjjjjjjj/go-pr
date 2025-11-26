package routes

import (
	"eCommerce/internal/handlers"
	"eCommerce/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {

	prod1 := router.Group("/api/v2/product")
	prod1.Use(middleware.AuthMiddleware())
	prod1.Use(middleware.AdminMiddleware())
	router.MaxMultipartMemory = 8 << 20
	prod1.POST("/new", handlers.ADDNew)

	prod := router.Group("api/v1/product")
	prod.Use(middleware.AuthMiddleware())
	prod.GET("/all", handlers.GetAll)
}
