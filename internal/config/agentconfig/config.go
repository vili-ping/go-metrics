package agentconfig

import (
	"flag"

	"github.com/caarlos0/env"
	"github.com/vili-ping/go-metrics/internal/utils"
)

var cfg agentConfig

type agentConfig struct {
	Address        string `env:"ADDRESS" envDefault:"localhost:8080"`
	ReportInterval int    `env:"REPORT_INTERVAL" envDefault:"10"`
	PoolInterval   int    `env:"POOL_INTERVAL" envDefault:"2"`
}

func parseEnvs() {
	err := env.Parse(&cfg)
	if err != nil {
		panic("env not parse")
	}
}

func ParseConfig() agentConfig {
	parseEnvs()

	flagAddress := flag.String("a", "localhost:8080", "address for server")
	flagReportInterval := flag.Int("r", 10, "report interval in seconds")
	flagPoolInterval := flag.Int("p", 2, "pool interval in second")
	flag.Parse()

	if !utils.IsEnvSet("ADDRESS") && *flagAddress != "" {
		cfg.Address = *flagAddress
	}

	if !utils.IsEnvSet("REPORT_INTERVAL") && *flagReportInterval != 0 {
		cfg.ReportInterval = *flagReportInterval
	}

	if !utils.IsEnvSet("POOL_INTERVAL") && *flagPoolInterval != 0 {
		cfg.PoolInterval = *flagPoolInterval
	}

	return cfg
}
