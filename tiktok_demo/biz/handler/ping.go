package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping return pong to test network
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

