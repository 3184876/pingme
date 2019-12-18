package main

import (
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	// Write response
	c.String(200, "OK")
}

func GetRecords(c *gin.Context) {
	var err error
	data := make(map[string]interface{})

	start := c.Query("start")
	end := c.Query("end")

	data["records"], err = getRecords(start, end)
	if err != nil {
		data["status"] = "fail"
	} else {
		data["status"] = "success"
	}

	// Write response
	c.JSON(200, data)
}
