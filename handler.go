package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
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

func Feed(ws *websocket.Conn) {
	var last_rec []byte
	for {
		now := time.Now().Unix()
		start := strconv.FormatInt(now-1, 10)
		end := strconv.FormatInt(now, 10)
		recMap, _ := getRecords(start, end)
		if len(recMap) != 0 {
			rec, _ := json.Marshal(recMap[0])
			if string(rec) != string(last_rec) {
				fmt.Println(string(rec))
				last_rec = rec
				_, _ = ws.Write(rec)
			}
		}
	}
}
