package main

import (
	"fmt"
	"net"
)

func lookupIP(addr string) string {
	ips, err := net.LookupIP(addr)
	if err != nil {
		fmt.Println(err)
	}
	return ips[0].String()
}
