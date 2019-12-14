package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
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
		Log.Info("    TCP     OPEN      ", address)
	} else if code == 1 {
		Log.Warn("    TCP     CLOSED    ", address)
	} else if code == 2 {
		Log.Warn("    TCP     ERROR     ", address)
	}
}

func logPing(dst *net.IPAddr, dur time.Duration, err error) {
	if err != nil {
		match, _ := regexp.MatchString("operation not permitted", err.Error())
		if match {
			Log.Warn(fmt.Sprintf("    ICMP    ERROR     Root permission is required."))
		} else {
			Log.Warn(fmt.Sprintf("    ICMP    ERROR     %s", dst.String()))
		}
		return
	}
	Log.Info(fmt.Sprintf("    ICMP    OPEN      %s    %s ms", dst.String(), strconv.FormatInt(dur.Milliseconds(), 10)))
}

func logMtr(hops []string, address string) {
	Log.Info(hops)
}
