package main

import (
	"flag"
	"log"

	"github.com/noobly314/pingme/tcping"
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
		//helper()
	case 1:
		addr := flag.Args()[0]
		ip := lookupIP(addr)
		tcping.New(ip, "22")
		tcping.New(ip, "80")
		tcping.New(ip, "443")
	case 2:
		addr := flag.Args()[0]
		port := flag.Args()[1]
		ip := lookupIP(addr)
		tcping.New(ip, port)
	default:
		Log.Warn("Too many arguments.")
	}

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
}
