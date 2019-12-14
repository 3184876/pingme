package main

import (
	"flag"
	"fmt"
	"os"
)

var Query string
var IDst string
var TDst string

var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

var Usage = func() {
	fmt.Fprintf(CommandLine.Output(), "Usage:\n")
	flag.PrintDefaults()
}

func init_flag() {
	flag.StringVar(&Query, "q", "", "Query address")
	flag.StringVar(&IDst, "i", "", "ICMP destination")
	flag.StringVar(&TDst, "t", "", "TCP destination")
	flag.Parse()
}

func isFlagPassed() bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		found = true
	})
	return found
}
