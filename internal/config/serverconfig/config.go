package serverconfig

import (
	"flag"

	"github.com/caarlos0/env"
	"github.com/vili-ping/go-metrics/internal/utils"
)

type serverConfig struct {
	Address  string `env:"ADDRESS" envDefault:"localhost:8080"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}

var cfg serverConfig

func parseEnvs() {
	err := env.Parse(&cfg)
	if err != nil {
		panic("env not parse")
	}
}

func ParseConfig() serverConfig {
	parseEnvs()

	flagAddress := flag.String("a", "localhost:8080", "address for server")
	flagLogLevel := flag.String("l", "info", "log level")
	flag.Parse()

	if !utils.IsEnvSet("ADDRESS") && *flagAddress != "" {
		cfg.Address = *flagAddress
	}

	if !utils.IsEnvSet("LOG_LEVEL") && *flagLogLevel != "" {
		cfg.LogLevel = *flagLogLevel
	}

	return cfg
}
