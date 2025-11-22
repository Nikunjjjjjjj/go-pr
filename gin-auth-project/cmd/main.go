package main

import (
	"gin-auth-project/config"
	"gin-auth-project/models"
	"gin-auth-project/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	// Connect DB
	config.ConnectDB()

	// Auto migrate
	config.DB.AutoMigrate(&models.User{})

	// Routes
	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.AdminRoutes(router)

	router.Run(":8080")
}

//task manager task,user api
