package main

import (
	"UploadFileInSql/internal/database"
	"UploadFileInSql/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDB()

	routes.Bothroutes(r)
	r.Run(":8080")
}
