package main

import (
	"flag"
	"fmt"
	"os"
)

var CommonPorts []string = []string{"22", "80", "443"}

var Version bool
var IsServe bool
var IsDaemon bool
var PingDst string
var TCPingDst string
var HTTPingDst string
var MtrDst string
var Query string

var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

var Usage = func() {
	fmt.Fprintf(CommandLine.Output(), "Usage:\n")
	flag.PrintDefaults()
}

func init_flag() {
	flag.BoolVar(&Version, "v", false, "Version")
	flag.BoolVar(&Version, "s", false, "Serve mode")
	flag.BoolVar(&Version, "d", false, "Daemon mode")
	flag.StringVar(&PingDst, "i", "", "ICMP destination")
	flag.StringVar(&TCPingDst, "t", "", "TCP destination")
	flag.StringVar(&HTTPingDst, "h", "", "HTTP destination")
	flag.StringVar(&MtrDst, "m", "", "MTR destination")
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
