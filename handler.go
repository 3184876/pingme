package main

import (
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	// Write response
	c.String(200, "OK")
}

func GetAllData(c *gin.Context) {
	data := make(map[string]interface{})

	data["records"] = getAllData()
	data["status"] = "success"

	// Write response
	c.JSON(200, data)
}
