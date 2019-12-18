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

	// Start
	fmt.Println("Listening...")
	go r.Run(":5000")

	pingLoop()
}

func pingLoop() {
	if !isFlagPassed("i") {
		log.Fatal("Please provide target address with -i flag.")
	} else {
		// ICMP Ping
		for {
			dst, dur, err := ping.New(PingDst)
			logPing(dst, dur, err)
			keyStr := "ICMP"
			keyStr += ":" + PingDst
			keyStr += ":" + dst.String()
			keyStr += ":" + strconv.FormatInt(dur.Milliseconds(), 10)
			keyStr += ":" + strconv.FormatInt(time.Now().Unix(), 10)
			fmt.Println(keyStr)
			key := []byte(keyStr)
			err = db.Put(key, []byte(""), nil)
			//data, _ := db.Get([]byte("hostname"), nil)
		}
	}
}
