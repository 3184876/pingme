package main

import (
	"flag"

	"github.com/noobly314/pingme-cli/ping"
	"github.com/noobly314/pingme-cli/tcping"
)

func init() {
	init_log()
}

func main() {
	var address string

	// Parse Args
	flag.Parse()

	switch len(flag.Args()) {
	case 0:
		Log.Warn("Please specify target.")
		//helper()
	case 1:
		addr := flag.Args()[0]
		ip := lookupIP(addr)

		// ICMP Ping
		dst, dur, err := ping.New(addr)
		logPing(dst, dur, err)

		// TCP Ping
		address = ip + ":22"
		c := tcping.New(address)
		logTcping(c, address)
		address = ip + ":80"
		c = tcping.New(address)
		logTcping(c, address)
		address = ip + ":443"
		c = tcping.New(address)
		logTcping(c, address)
	case 2:
		addr := flag.Args()[0]
		port := flag.Args()[1]
		ip := lookupIP(addr)
		address = ip + ":" + port
		c := tcping.New(address)
		logTcping(c, address)
	default:
		Log.Warn("Too many arguments.")
	}
}
