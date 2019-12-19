package main

import (
	"fmt"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
)

func serve() {
	// Set router
	r := gin.Default()

	// CORS
	r.Use(cors.Default())

	// Handler
	r.GET("/", Hello)
	r.GET("/records", GetRecords)
	r.GET("/ws", func(c *gin.Context) {
		handler := websocket.Handler(Stream)
		handler.ServeHTTP(c.Writer, c.Request)
	})

	// Start
	fmt.Println("Listening...")
	r.Run(":" + strconv.Itoa(Config.ServePort))
}
