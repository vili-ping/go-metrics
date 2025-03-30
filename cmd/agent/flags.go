package main

import "flag"

var (
	flagAddress        string
	flagReportInterval int
	flagPoolInterval   int
)

func parseArgs() {
	flag.StringVar(&flagAddress, "a", "localhost:8080", "address for server")
	flag.IntVar(&flagReportInterval, "r", 10, "report interval in seconds")
	flag.IntVar(&flagPoolInterval, "p", 2, "pool interval in second")
	flag.Parse()
}
