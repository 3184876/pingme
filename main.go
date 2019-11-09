package main

import (
	"flag"
	"fmt"
	"net"
	"regexp"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Global Variables
var fmtOK *color.Color
var fmtInfo *color.Color
var fmtWarn *color.Color
var fmtError *color.Color

func init() {
	// Set colorful fmt
	fmtOK = color.New(color.FgGreen).Add(color.Bold)
	fmtInfo = color.New(color.FgBlue).Add(color.Bold)
	fmtWarn = color.New(color.FgYellow).Add(color.Bold)
	fmtError = color.New(color.FgRed).Add(color.Bold)
}

func main() {
	// Parse Args
	flag.Parse()

	switch len(flag.Args()) {
	case 0:
		//fmtWarn.Println("Please specify target.")
		helper()
	case 1:
		addr := flag.Args()[0]
		ip := lookupIP(addr)
		tcpPing(ip, "22")
		tcpPing(ip, "80")
		tcpPing(ip, "443")
	case 2:
		addr := flag.Args()[0]
		port := flag.Args()[1]
		ip := lookupIP(addr)
		tcpPing(ip, port)
	default:
		fmtWarn.Println("Too many arguments.")
	}
}

func catDialErr(err error) string {
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

func tcpPing(ip string, port string) {
	tcpAddr := strings.Join([]string{ip, ":", port}, "")
	d := net.Dialer{Timeout: 3 * time.Second}
	_, err := d.Dial("tcp", tcpAddr)
	if err != nil {
		errType := catDialErr(err)
		if errType == "refused" {
			fmtInfo.Printf(tcpAddr)
			fmtWarn.Printf(" is reachable, but the port is closed.\n")
		}
		if errType == "timeout" {
			fmtInfo.Printf(tcpAddr)
			fmtError.Printf(" is unreachable.\n")
		}
	} else {
		fmtInfo.Printf(tcpAddr)
		fmtOK.Printf(" is reachable.\n")
	}
}

func helper() {
	fmtInfo.Println("PingMe helper:")
	fmt.Println()
	fmtInfo.Println("1. Basic usage: pingme <ip> <port>")
	fmt.Println("e.g. pingme google.com 443")
	fmt.Println()
	fmtInfo.Println("2. Input ip only: pingme <ip>")
	fmt.Println("e.g. pingme facebook.com")
	fmt.Println("PingMe will scan 3 common ports: 22, 80, 443")
	fmt.Println()
	fmtInfo.Println("3. Input URL: pingme <url>")
	fmt.Println("e.g. pingme https://twitter.com")
	fmt.Println("PingMe will scan 80 for http / 443 for https")
	fmt.Println()
	fmtInfo.Println("4. Add info flag: pingme -i <ip>")
	fmt.Println("e.g. pingme -i https://wikipedia.org")
	fmt.Println("PingMe will query detailed information about ip from https://pingme.cc")
	fmt.Println()
}
