# demo-go-httptrace

This project show how using http trace package of standard libraries

Often we need check metadata of request and response from calls http's. We usually implement a function to intercept request a response and do something. But GO provide a default package for do this, called http trace.

##### Example:
```go
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

```

##### For more details: https://blog.golang.org/http-tracing