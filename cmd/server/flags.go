package main

import "flag"

var flagAddress string

func parseArgs() {
	flag.StringVar(&flagAddress, "a", "localhost:8080", "address for server")
	flag.Parse()
}
