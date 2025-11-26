package main

import (
	"eCommerce/internal/database"
	"eCommerce/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//cfg:=config.LoadEnv()
	database.ConnectDB()
	// Auto migrate kha karu?
	//database.DB.AutoMigrate()
	//add routes
	routes.AuthRoutes(r)
	routes.ProductRoutes(r)
	routes.CartRoutes(r)
	routes.OrderRoute(r)
	r.Run(":8080")
}
