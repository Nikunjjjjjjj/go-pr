package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("me/:id", func(c *gin.Context) {
		var id = c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"userId": id,
		})

	})

	router.POST("post", func(c *gin.Context) {
		type MeRequest struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password"`
		}

		var meRequest MeRequest

		err := c.BindJSON(&meRequest)
		fmt.Println(err, 35)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"email":   meRequest.Email,
			"pasword": meRequest.Password,
		})
	})

	router.Run(":8080")
}
