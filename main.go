package main

import (
	"flag"

	"github.com/noobly314/pingme-cli/mtr"
	"github.com/noobly314/pingme-cli/ping"
	"github.com/noobly314/pingme-cli/tcping"
)

func init() {
	init_log()
	init_flag()
}

func main() {
	if !hasFlag() {
		switch len(flag.Args()) {
		case 0:
			Log.Warn("Please specify target.")
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
			Log.Warn("Too many arguments.")
		}
	} else {
		if isFlagPassed("i") {
			// ICMP Ping
			dst, dur, err := ping.New(PingDst)
			logPing(dst, dur, err)
		}
		if isFlagPassed("t") {
			// TCP Ping
			address := TCPingDst
			c := tcping.New(address)
			logTcping(c, address)
		}
		if isFlagPassed("q") {
			//fmt.Println(Query)
		}
	}

	mtr.New("adf")
}
