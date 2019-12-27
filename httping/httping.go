package httping

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)

func New(address string) int {
	req, _ := http.NewRequest("GET", "http://google.com", nil)
	trace := &httptrace.ClientTrace{
		DNSStart: func(dnsInfo httptrace.DNSStartInfo) {
			fmt.Printf("DNS Start Info: %+v\n", dnsInfo)
			fmt.Println(time.Now().Unix())
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Done Info: %+v\n", dnsInfo)
			fmt.Println(time.Now().Unix())
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
			fmt.Println(time.Now().Unix())
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	_, err := http.DefaultTransport.RoundTrip(req)
	//fmt.Println(*c)
	if err != nil {
		log.Fatal(err)
	}
	return 1
}
