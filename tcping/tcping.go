package tcping

import (
	"net"
	"regexp"
	"time"
)

func New(address string) int {
	d := net.Dialer{Timeout: 3 * time.Second}
	_, err := d.Dial("tcp", address)
	if err != nil {
		errType := getDialError(err)
		if errType == "refused" {
			// Closed
			return 1
		}
		if errType == "timeout" {
			// Error
			return 2
		}
	} else {
		// Open
		return 0
	}
	// Default
	return 2
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
