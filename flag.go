package main

import (
	"flag"
	"fmt"
	"os"
)

var PingDst string
var TCPingDst string
var Query string

var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

var Usage = func() {
	fmt.Fprintf(CommandLine.Output(), "Usage:\n")
	flag.PrintDefaults()
}

func init_flag() {
	flag.StringVar(&PingDst, "i", "", "ICMP destination")
	flag.StringVar(&TCPingDst, "t", "", "TCP destination")
	flag.StringVar(&Query, "q", "", "Query address")
	flag.Parse()
}

func hasFlag() bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		found = true
	})
	return found
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
