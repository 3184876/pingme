package main

import (
	"flag"
	"fmt"
	"net"

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
	if isFlagPassed("c") {
		init_config()
		init_db()
	}
}

func main() {
	if !hasFlag() {
		switch len(flag.Args()) {
		case 0:
			log.Warn("Please specify target.")
		case 1:
			addr := flag.Args()[0]
			ip := lookupIP(addr)

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
		} else if isFlagPassed("s") {
			// Serve mode
			if !isFlagPassed("c") {
				log.Fatal("Please provide config file with -c flag.")
			} else {
				if !isFlagPassed("i") {
					serve()
				} else {
					go serve()
					pingLoop()
				}
			}
		} else if isFlagPassed("d") {
			// Daemon mode
		} else if isFlagPassed("i") {
			// ICMP Ping
			dst, dur, err := ping.New(PingDst)
			logPing(dst, dur, err)
		} else if isFlagPassed("t") {
			// TCP Ping
			c := tcping.New(TCPingDst)
			logTcping(c, TCPingDst)
		} else if isFlagPassed("m") {
			// MTR
			hops, err := mtr.New(MtrDst)
			if err != nil {
				log.Fatal(err)
			}
			logMtr(hops, MtrDst)
		} else if isFlagPassed("q") {
			ip := lookupIP(Query)
			queryInfo(ip)
		}
	}
}
