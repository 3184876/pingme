package main

import (
	"flag"
	"fmt"
	"net"
	"regexp"
	"strings"
	"time"
)

func init() {
	init_log()
}

func main() {
	// Parse Args
	flag.Parse()

	switch len(flag.Args()) {
	case 0:
		Log.Warn("Please specify target.")
		helper()
	case 1:
		addr := flag.Args()[0]
		ip := lookupIP(addr)
		tcping(ip, "22")
		tcping(ip, "80")
		tcping(ip, "443")
	case 2:
		addr := flag.Args()[0]
		port := flag.Args()[1]
		ip := lookupIP(addr)
		tcping(ip, port)
	default:
		Log.Warn("Too many arguments.")
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

func lookupIP(addr string) string {
	ips, err := net.LookupIP(addr)
	if err != nil {
		fmt.Println(err)
	}
	return ips[0].String()
}

func tcping(ip string, port string) {
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
