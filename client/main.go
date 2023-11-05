package main

import (
	"flag"
)

var (
	server string
)

func main() {
	flag.StringVar(&server, "s", "", "主机域名/IP")
	flag.Parse()
}