package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/detal", func(c gin.Context) {
		c.JSON(200, gin.H{
			"name":   "JOHN",
			"role":   "Admin",
			"active": true,
		})
	})
	r.Run("8080")
}
