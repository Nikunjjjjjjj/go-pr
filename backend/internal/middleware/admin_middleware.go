package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, ok := c.Get("role")
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "role mot found",
			})
			c.Abort()
			return
		}
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "access denied",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
