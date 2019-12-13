package main

import (
	"os"

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
