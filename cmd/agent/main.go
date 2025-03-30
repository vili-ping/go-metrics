package main

import (
	"fmt"
	"time"

	"github.com/vili-ping/go-metrics/internal/agent/handlers"
	"github.com/vili-ping/go-metrics/internal/config/agentconfig"
)

var config = agentconfig.ParseConfig()

func main() {
	var metrics handlers.Metrics

	fmt.Printf("Agent start with address=%s;report-inteval=%d,pool-interval=%d\n", config.Address, config.ReportInterval, config.PoolInterval)

	for {
		metrics.CollectMetrics()
		time.Sleep(time.Duration(config.PoolInterval) * time.Second)

		metrics.SendMetrics(config.Address)
		time.Sleep(time.Duration(config.ReportInterval) * time.Second)
	}
}
