package utils

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
)

func ReverseProxy(c *gin.Context, reversePath string, method string) {

	req := c.Request

	if method == "" {
		method = req.Method
	}

	proxy, err := url.Parse(reversePath)
	if err != nil {
		log.Printf("error in parse addr: %v", err)
		c.String(500, "error")
		return
	}
	//req.URL.Scheme = proxy.Scheme
	//req.URL.Host = proxy.Host
	req.URL = proxy
	req.Method = "POST"
	log.Printf("proxy to %s", req.URL.String())

	// step 2: use http.Transport to do request to real server.
	transport := http.DefaultTransport
	resp, err := transport.RoundTrip(req)
	if err != nil {
		log.Printf("error in roundtrip: %v", err)
		c.String(500, "error")
		return
	}

	// step 3: return real server response to upstream.
	for k, vv := range resp.Header {
		for _, v := range vv {
			c.Header(k, v)
		}
	}
	defer resp.Body.Close()
	bufio.NewReader(resp.Body).WriteTo(c.Writer)
	return
}
