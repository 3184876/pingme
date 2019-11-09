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
var fmtWarn *color.Color
var fmtError *color.Color

func init() {
	// Set colorful fmt
	fmtOK = color.New(color.FgGreen).Add(color.Bold)
	fmtWarn = color.New(color.FgYellow).Add(color.Bold)
	fmtError = color.New(color.FgRed).Add(color.Bold)
}

func main() {
	// Parse Args
	flag.Parse()
	l := len(flag.Args())
	if l == 0 {
		fmtWarn.Println("Please specify target.")
	} else if l == 1 {
		// ICMP Ping
	} else if l == 2 {
		// TCP Ping
		addr := flag.Args()[0]
		port := flag.Args()[1]
		err := tcpPing(addr, port)
		if err != nil {
			//fmtError.Println("Error:", err.Error())
			errType := catDialErr(err)
			//fmt.Println(errType)
			if errType == "refused" {
				fmtWarn.Println("TCP is reachable, but the port is closed.")
				tcpAddr := strings.Join([]string{addr, ":", port}, "")
				fmt.Println(tcpAddr)
			}
			if errType == "timeout" {
				fmtError.Println("TCP is unreachable.")
				tcpAddr := strings.Join([]string{addr, ":", port}, "")
				fmt.Println(tcpAddr)
			}
		} else {
			fmtOK.Println("TCP is OK")
			tcpAddr := strings.Join([]string{addr, ":", port}, "")
			fmt.Println(tcpAddr)
		}
	} else {
		fmtWarn.Println("Too many arguments.")
	}
}

func tcpPing(addr string, port string) error {
	tcpAddr := strings.Join([]string{addr, ":", port}, "")
	d := net.Dialer{Timeout: 3 * time.Second}
	_, err := d.Dial("tcp", tcpAddr)
	return err
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
