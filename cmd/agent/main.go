package main

import (
	"time"

	"github.com/vili-ping/go-metrics/internal/agent/handlers"
)

func main() {
	var metrics handlers.Metrics

	for {
		metrics.CollectMetrics()
		time.Sleep(2 * time.Second)

		metrics.SendMetrics()
		time.Sleep(10 * time.Second)
	}
}
