package main

import (
	"log"
	"url-shortner/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	// router := gin.Default()
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// router.Run()

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	h := handlers.NewURLHandler(redisClient)

	r.POST("/shorten", h.ShortenUrl)
	r.GET("/:code", h.ResolveURL)

	log.Println("Server running on :8080")
	r.Run(":8080")
}
