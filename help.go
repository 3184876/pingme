package main

import "fmt"

func helper() {
	fmt.Println()
	fmt.Println("PingMe helper:")
	fmt.Println()
	fmt.Println("1. Basic usage: pingme <ip> <port>")
	fmt.Println("e.g. pingme google.com 443")
	fmt.Println()
	fmt.Println("2. Input ip only: pingme <ip>")
	fmt.Println("e.g. pingme facebook.com")
	fmt.Println("PingMe will scan 3 common ports: 22, 80, 443")
	fmt.Println()
	fmt.Println("3. Input URL: pingme <url>")
	fmt.Println("e.g. pingme https://twitter.com")
	fmt.Println("PingMe will scan 80 for http / 443 for https")
	fmt.Println()
	fmt.Println("4. Query IP info: pingme -i <ip>")
	fmt.Println("e.g. pingme -i https://wikipedia.org")
	fmt.Println("PingMe will query ip information from https://pingme.cc")
	fmt.Println()
}
