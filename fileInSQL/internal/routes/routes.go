package routes

import (
	"UploadFileInSql/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Bothroutes(r *gin.Engine) {

	r.POST("/upload", handlers.UploadCSV)

	r.GET("/pincode", handlers.PincodeDetails)
}
