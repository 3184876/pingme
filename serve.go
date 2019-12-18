package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/noobly314/pingme/ping"
)

func serve() {
	// Set router
	r := gin.Default()

	// Handler
	r.GET("/", Hello)
	r.GET("/all", GetAllData)

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
