package handlers

import "github.com/gin-gonic/gin"

// HealthCheck handles a ping for healthcheck and discovery purposes
func HealthCheck() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	}
}
