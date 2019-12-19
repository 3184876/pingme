package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/noobly314/pingme/ping"
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
		handler := websocket.Handler(Feed)
		handler.ServeHTTP(c.Writer, c.Request)
	})

	// Start
	fmt.Println("Listening...")
	r.Run(":" + strconv.Itoa(Config.ServePort))
}

func pingLoop() {
	for {
		dst, dur, err := ping.New(PingDst)
		logPing(dst, dur, err)
		key := "ICMP"
		key += ":" + strconv.FormatInt(time.Now().Unix(), 10)
		key += ":" + PingDst
		key += ":" + dst.String()
		key += ":" + strconv.FormatInt(dur.Milliseconds(), 10)
		err = db.Put([]byte(key), []byte(""), nil)
		if err != nil {
			log.Warn(err)
		}
	}
}
