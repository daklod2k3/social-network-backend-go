package utils

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReverseProxy(c *gin.Context, reversePath string, method string) {

	req := c.Request

	if method == "" {
		method = req.Method
	}

	target, err := url.Parse(reversePath)
	if err != nil {
		log.Printf("error in parse addr: %v", err)
		c.String(500, "error")
		return
	}
	//req.URL.Scheme = proxy.Scheme
	//req.URL.Host = proxy.Host
	req.URL = target
	req.Method = method
	log.Printf("proxy to %s", req.URL.String())

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	proxy.ServeHTTP(c.Writer, c.Request)
	//// step 2: use http.Transport to do request to real server.
	//transport := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	//resp, err := transport.RoundTrip(req)
	//if err != nil {
	//	log.Printf("error in roundtrip: %v", err)
	//	c.String(500, "error")
	//	return
	//}
	//
	//// step 3: return real server response to upstream.
	//for k, vv := range resp.Header {
	//	for _, v := range vv {
	//		c.Header(k, v)
	//	}
	//}
	//defer resp.Body.Close()
	//bufio.NewReader(resp.Body).WriteTo(c.Writer)
	return
}

type ProxyHandler struct {
	p *httputil.ReverseProxy
}

func (ph *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	ph.p.ServeHTTP(w, r)
}
