package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	endpoint := flag.String("h", "http://localhost:40080", "Endpoint")
	port := flag.String("p", ":8080", "Server port")
	flag.Parse()
	host, err := url.Parse(*endpoint)
	if err != nil {
		panic(err)
	}
	log.Fatal(http.ListenAndServe(*port, NewProxy(host)))
}

type Proxy struct {
	reverseProxy *httputil.ReverseProxy
}

func NewProxy(target *url.URL) *Proxy {
	return &Proxy{
		reverseProxy: httputil.NewSingleHostReverseProxy(target),
	}
}

func (proxy *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	orgHost := r.Host
	r.Host = "localhost"
	proxy.reverseProxy.ServeHTTP(w, r)
	r.Host = orgHost
}
