package tcping

import (
	"net"
	"regexp"
	"strings"
	"time"
)

func New(ip string, port string) {
	address := strings.Join([]string{ip, ":", port}, "")
	d := net.Dialer{Timeout: 3 * time.Second}
	_, err := d.Dial("tcp", address)
	if err != nil {
		errType := getDialError(err)
		if errType == "refused" {
			Log.Warn("CLOSED    ", address)
		}
		if errType == "timeout" {
			Log.Warn("ERROR     ", address)
		}
	} else {
		Log.Info("OPEN      ", address)
	}
}

func getDialError(err error) string {
	match, _ := regexp.MatchString("timeout", err.Error())
	if match {
		return "timeout"
	}
	match, _ = regexp.MatchString("refused", err.Error())
	if match {
		return "refused"
	}
	return ""
}
