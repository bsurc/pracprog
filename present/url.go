package main

import (
	"fmt"
	"log"
	"net"
	"net/url"
)

func main() {
	// START_OMIT
	u, err := url.Parse("https://boisestate.edu:443/magic/place?name=kyle&date=2018-01-01")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Scheme: %s\n", u.Scheme)
	h, p, _ := net.SplitHostPort(u.Host)
	fmt.Printf("Host: %s\n", h)
	fmt.Printf("Port: %s\n", p)
	fmt.Printf("Path: %s\n", u.Path)
	fmt.Printf("Query String: %s\n", u.RawQuery)
	for k, v := range u.Query() {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}
	// END_OMIT
}
