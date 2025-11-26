package middleware

import (
	"eCommerce/internal/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "missing token",
			})
			c.Abort()
			return
		}

		// Check format
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}
		//fmt.Println(31)
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			fmt.Println("Token error:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "token problem"})
			c.Abort()
			return
		}
		// fmt.Println(37)
		// fmt.Println("token", token)
		// fmt.Println("User ID:", claims["user_id"])
		// fmt.Println("Role:", claims["role"])
		// fmt.Println("Email:", claims["email"])

		c.Set("user_id", claims["user_id"])
		c.Set("role", claims["role"])
		c.Set("email", claims["email"])
		c.Next()
	}
}
