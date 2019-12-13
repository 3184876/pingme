package main

import (
	"flag"

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
		address = ip + ":22"
		logTcping(tcping.New(address), address)
		address = ip + ":80"
		logTcping(tcping.New(address), address)
		address = ip + ":443"
		logTcping(tcping.New(address), address)
	case 2:
		addr := flag.Args()[0]
		port := flag.Args()[1]
		ip := lookupIP(addr)
		address = ip + ":" + port
		logTcping(tcping.New(address), address)
	default:
		Log.Warn("Too many arguments.")
	}

	/*
		// ICMP Ping
		p := func(addr string) {
			dst, dur, err := Ping(addr)
			if err != nil {
				log.Printf("Ping %s (%s): %s\n", addr, dst, err)
				return
			}
			log.Printf("Ping %s (%s): %s\n", addr, dst, dur)
		}
		p("127.0.0.1")
		p("baidu.com")
		p("google.com")
	*/
}
