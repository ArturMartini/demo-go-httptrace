package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
)

var client http.Client

func init() {
	client = http.Client{}
}

func main() {
	req, _ := http.NewRequest(http.MethodGet, "https://example.com", nil)
	trace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", info)
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			fmt.Printf("Got DNS: %+v\n", info)
		},
		TLSHandshakeDone: func(state tls.ConnectionState, _ error) {
			fmt.Printf("TLS Handshake: %+v\n", state)
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
		log.Fatal(err)
	}

	client.Do(req)
	//The Transport in the net/http package supports tracing of both HTTP/1 and HTTP/2 requests.
}
