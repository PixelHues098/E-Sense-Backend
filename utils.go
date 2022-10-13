package main

import (
	"github.com/gin-gonic/gin"
)

func makeGinResponse(c *gin.Context, statusCode int, value string) {
	c.JSON(statusCode, gin.H{
		"message":    value,
		"statusCode": statusCode,
	})
}
