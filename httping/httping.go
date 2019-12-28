package httping

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"strings"
	"time"
)

type Stats struct {
	Scheme   string
	DNS      int64
	TCP      int64
	TLS      int64
	Process  int64
	Transfer int64
	Total    int64
}

func New(address string) (Stats, error) {
	var err error
	var stats Stats
	var t0, t1, t2, t3, t4, t5, t6, t7 int64

	if strings.HasPrefix(address, "http://") {
		stats.Scheme = "http"
	} else if strings.HasPrefix(address, "https://") {
		stats.Scheme = "https"
	} else {
		address = "http://" + address
		stats.Scheme = "http"
	}
	fmt.Println("address: ", address)

	req, _ := http.NewRequest("GET", address, nil)
	trace := &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			//fmt.Printf("DNS Start Info: %+v\n", info)
			t0 = time.Now().UnixNano()
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			//fmt.Printf("DNS Done Info: %+v\n", info)
			t1 = time.Now().UnixNano()
			if info.Err != nil {
				err = info.Err
				log.Fatal(info.Err)
			}
		},
		ConnectStart: func(net, addr string) {
		},
		ConnectDone: func(net, addr string, err error) {
			if err != nil {
				log.Fatalf("unable to connect to host %v: %v", addr, err)
			}
			t2 = time.Now().UnixNano()
		},
		GotConn: func(info httptrace.GotConnInfo) {
			//fmt.Printf("Got Conn: %+v\n", info)
			t3 = time.Now().UnixNano()
		},
		GotFirstResponseByte: func() {
			t4 = time.Now().UnixNano()
		},
		TLSHandshakeStart: func() {
			t5 = time.Now().UnixNano()
		},
		TLSHandshakeDone: func(_ tls.ConnectionState, _ error) {
			t6 = time.Now().UnixNano()
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	_, err = http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}

	t7 = time.Now().UnixNano()

	stats.DNS = t1 - t0
	stats.TCP = t3 - t1
	stats.Process = t4 - t3
	stats.Transfer = t7 - t4
	stats.TLS = t6 - t5
	stats.Total = t7 - t0

	return stats, err
}
