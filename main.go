package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/noobly314/pingme/httping"
	"github.com/noobly314/pingme/mtr"
	"github.com/noobly314/pingme/ping"
	"github.com/noobly314/pingme/tcping"
)

var (
	VersionString string
)

func init() {
	init_log()
	init_flag()
}

func main() {
	if !hasFlag() {
		switch len(flag.Args()) {
		case 0:
			log.Warn("Please specify target.")
		case 1:
			addr := flag.Args()[0]
			ip := lookupIP(addr)

			// Query
			address := parseInput(addr)
			queryInfo(address)
			fmt.Println()

			// ICMP Ping
			dst, dur, err := ping.New(ip)
			logPing(dst, dur, err)

			// TCP Ping
			for _, port := range CommonPorts {
				address := net.JoinHostPort(ip, port)
				c := tcping.New(address)
				logTcping(c, address)
			}
		case 2:
			addr := flag.Args()[0]
			port := flag.Args()[1]
			ip := lookupIP(addr)
			address := net.JoinHostPort(ip, port)
			c := tcping.New(address)
			logTcping(c, address)
		default:
			log.Warn("Too many arguments.")
		}
	} else {
		if isFlagPassed("v") {
			// Version
			fmt.Println(VersionString)
		} else if isFlagPassed("i") {
			// ICMP Ping
			dst, dur, err := ping.New(PingDst)
			logPing(dst, dur, err)
		} else if isFlagPassed("t") {
			// TCP Ping
			c := tcping.New(TCPingDst)
			logTcping(c, TCPingDst)
		} else if isFlagPassed("h") {
			// HTTP Ping
			_ = httping.New(HTTPingDst)
			//c := httping.New(HTTPingDst)
			//logHttping(c, HTTPingDst)
		} else if isFlagPassed("m") {
			// MTR
			hops, err := mtr.New(MtrDst)
			if err != nil {
				log.Fatal(err)
			}
			logMtr(hops, MtrDst)
		} else if isFlagPassed("q") {
			address := parseInput(Query)
			queryInfo(address)
		}
	}
}
