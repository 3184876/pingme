package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"time"

	"github.com/noobly314/pingme/httping"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init_log() {
	formatter := &logrus.TextFormatter{
		DisableTimestamp: true,
	}
	log.SetFormatter(formatter)
	log.Out = os.Stdout
}

func logHttping(stats httping.Stats, err error, address string) {
	if err == nil {
		if stats.Scheme == "http" {
			fmt.Printf("Scheme    :    %s\n", stats.Scheme)
			fmt.Printf("Host      :    %s\n", parseInput(address))
			fmt.Printf("DNS Lookup:    %.2f ms\n", float64(stats.DNS)/1e6)
			fmt.Printf("TCP       :    %.2f ms\n", float64(stats.TCP)/1e6)
			fmt.Printf("Process   :    %.2f ms\n", float64(stats.Process)/1e6)
			fmt.Printf("Transfer  :    %.2f ms\n", float64(stats.Transfer)/1e6)
			fmt.Printf("Total     :    %.2f ms\n", float64(stats.Total)/1e6)
		} else if stats.Scheme == "https" {
			fmt.Printf("Scheme    :    %s\n", stats.Scheme)
			fmt.Printf("Host      :    %s\n", parseInput(address))
			fmt.Printf("DNS Lookup:    %.2f ms\n", float64(stats.DNS)/1e6)
			fmt.Printf("TCP       :    %.2f ms\n", float64(stats.TCP)/1e6)
			fmt.Printf("TLS       :    %.2f ms\n", float64(stats.TLS)/1e6)
			fmt.Printf("Process   :    %.2f ms\n", float64(stats.Process)/1e6)
			fmt.Printf("Transfer  :    %.2f ms\n", float64(stats.Transfer)/1e6)
			fmt.Printf("Total     :    %.2f ms\n", float64(stats.Total)/1e6)
		}
	}
}

func logTcping(code int, address string) {
	if code == 0 {
		fmt.Printf("TCP     OPEN      %s\n", address)
	} else if code == 1 {
		fmt.Printf("TCP     CLOSED    %s\n", address)
	} else if code == 2 {
		fmt.Printf("TCP     ERROR     %s\n", address)
	}
}

func logPing(dst *net.IPAddr, dur time.Duration, err error) {
	if err != nil {
		match, _ := regexp.MatchString("operation not permitted", err.Error())
		if match {
			fmt.Printf("ICMP    ERROR     No privileges\n")
		} else {
			fmt.Printf("ICMP    ERROR     %s\n", dst.String())
		}
		return
	}
	fmt.Printf("ICMP    OPEN      %s    %s ms\n", dst.String(), fmt.Sprintf("%.1f", float64(dur.Microseconds())/1000))
}

func logMtr(hops []string, address string) {
	for _, h := range hops {
		fmt.Printf("MTR     %s\n", h)
	}
}
