package middleware

import "github.com/gin-gonic/gin"

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
