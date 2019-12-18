package main

import (
	"flag"
	"fmt"

	"github.com/noobly314/pingme/mtr"
	"github.com/noobly314/pingme/ping"
	"github.com/noobly314/pingme/tcping"
)

const (
	VersionString string = "pingme v0.1.1"
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
				address := ip + ":" + port
				c := tcping.New(address)
				logTcping(c, address)
			}
		case 2:
			addr := flag.Args()[0]
			port := flag.Args()[1]
			ip := lookupIP(addr)
			address := ip + ":" + port
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
					log.Fatal("Please provide target address with -i flag.")
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
			address := TCPingDst
			c := tcping.New(address)
			logTcping(c, address)
		} else if isFlagPassed("m") {
			// MTR
			address := MtrDst
			hops, err := mtr.New(address)
			if err != nil {
				log.Fatal(err)
			}
			logMtr(hops, address)
		} else if isFlagPassed("q") {
			//fmt.Println(Query)
		}
	}
}
