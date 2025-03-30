package agentconfig

import (
	"flag"
	"os"
	"strconv"
)

type agentConfig struct {
	Address        string
	ReportInterval int
	PoolInterval   int
}

func parseConfAddr(c *agentConfig) {
	envAddr, envAddrExist := os.LookupEnv("ADDRESS")
	if envAddrExist {
		c.Address = envAddr
	}
}

func parseConfReportInterval(c *agentConfig) {
	envReportInterval, envReportIntervalExist := os.LookupEnv("REPORT_INTERVAL")
	if envReportIntervalExist {
		c.ReportInterval, _ = strconv.Atoi(envReportInterval)
	}
}

func parseConfPoolInterval(c *agentConfig) {
	envPollInterval, envPoolInterval := os.LookupEnv("POLL_INTERVAL")
	if envPoolInterval {
		c.PoolInterval, _ = strconv.Atoi(envPollInterval)
	}
}

func ParseConfig() agentConfig {
	var config agentConfig

	flag.StringVar(&config.Address, "a", "localhost:8080", "address for server")
	flag.IntVar(&config.ReportInterval, "r", 10, "report interval in seconds")
	flag.IntVar(&config.PoolInterval, "p", 2, "pool interval in second")
	flag.Parse()

	parseConfAddr(&config)
	parseConfPoolInterval(&config)
	parseConfReportInterval(&config)

	return config
}
