package controllers

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"os"
	"social-backend/internal/server/utils"
)

var supabaseUrl = os.Getenv("SUPABASE_URL")
var reversePath = supabaseUrl + "/auth/v1"

func Login(c *gin.Context) {
	utils.reverseProxy(c, reversePath+"/token", "POST")
}

func Register(c *gin.Context) {
	req := c.Request
	proxy, err := url.Parse(reversePath + "/signup")
	if err != nil {
		log.Printf("error in parse addr: %v", err)
		c.String(500, "error")
		return
	}
	req.URL.Scheme = proxy.Scheme
	req.URL.Host = proxy.Host
	req.URL.Path = proxy.Path
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
