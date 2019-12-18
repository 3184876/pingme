package main

import (
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	// Write response
	c.String(200, "OK")
}
