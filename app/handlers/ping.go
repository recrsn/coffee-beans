package handlers

import "github.com/gin-gonic/gin"

// Ping handles a ping for healthcheck and discovery purposes
func Ping() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}
