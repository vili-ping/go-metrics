package main

import (
	"time"

	"github.com/vili-ping/go-metrics/internal/agent/handlers"
)

func main() {
	parseArgs()

	var metrics handlers.Metrics

	for {
		metrics.CollectMetrics()
		time.Sleep(time.Duration(flagPoolInterval) * time.Second)

		metrics.SendMetrics(flagAddress)
		time.Sleep(time.Duration(flagReportInterval) * time.Second)
	}
}
