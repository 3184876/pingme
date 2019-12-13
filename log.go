package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init_log() {
	formatter := &logrus.TextFormatter{
		DisableTimestamp: true,
	}
	Log.SetFormatter(formatter)
	Log.Out = os.Stdout
}

func logTcping(code int, address string) {
	if code == 0 {
		Log.Info("    OPEN      ", address)
	} else if code == 1 {
		Log.Warn("    CLOSED    ", address)
	} else if code == 2 {
		Log.Warn("    ERROR     ", address)
	}
}

func logPing(dst *net.IPAddr, dur time.Duration, err error, address string) {
	if err != nil {
		Log.Warn("Ping "+address+" ("+dst.String()+"): "+err.Error()+"\n", address, dst, err)
		return
	}
	Log.Info("Ping "+address+" ("+dst.String()+") \n", address, dst)
	fmt.Println(dur)
}
