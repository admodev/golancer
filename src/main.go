package main

import (
	"sync"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Backend struct {
	URL						*url.URL
	Alive					bool
	mux						sync.RWMutex
	ReverseProxy	*httputil.ReverseProxy
}

type ServerPool struct {
	backends	[]*Backend
	current		uint64
}

func main() {
	u, _ := url.Parse("http://localhost:8080")
	rp := httputil.NewSingleHostReverseProxy(u)
	
	http.HandlerFunc(rp.ServeHTTP)
}