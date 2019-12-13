package main

import (
	"flag"
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
		helper()
	case 1:
		addr := flag.Args()[0]
		ip := lookupIP(addr)
		tcping(ip, "22")
		tcping(ip, "80")
		tcping(ip, "443")
	case 2:
		addr := flag.Args()[0]
		port := flag.Args()[1]
		ip := lookupIP(addr)
		tcping(ip, port)
	default:
		Log.Warn("Too many arguments.")
	}
}
