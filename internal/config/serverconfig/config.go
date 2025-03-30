package serverconfig

import (
	"flag"
	"os"
)

type serverConfig struct {
	Address string
}

func parseConfAddr(c *serverConfig) {
	envAddr, envAddrExist := os.LookupEnv("ADDRESS")
	if envAddrExist {
		c.Address = envAddr
		return
	}
}

func ParseConfig() serverConfig {
	var config serverConfig

	flag.StringVar(&config.Address, "a", "localhost:8080", "address for server")
	flag.Parse()

	parseConfAddr(&config)

	return config
}
