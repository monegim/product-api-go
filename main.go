package main

import (
	"github.com/hashicorp-demoapp/go-hckit"
	"github.com/hashicorp/go-hclog"
	"github.com/nicholasjackson/env"
	"os"
)

type Config struct {
	DBConnection           string  `json:"db_connection"`
	BindAddress            string  `json:"bind_address"`
	MetricsAddress         string  `json:"metrics_address"`
	MaxRetries             int     `json:"max_retries"`
	BackoffExponentialBase float64 `json:"backoff_exponential_base"`
}

var logger hclog.Logger
var conf *Config

var configFile = env.String("CONFIG_FILE", false, "./conf.json", "Path to JSON encoded config file")
var dbConnection = env.String("DB_CONNECTION", false, "", "db connection string")
var bindAddress = env.String("BIND_ADDRESS", false, "", "Bind address")
var metricsAddress = env.String("METRICS_ADDRESS", false, "", "Metrics address")
var maxRetries = env.Int("MAX_RETRIES", false, 60, "Maximum number of connection retries")
var backoffExponentialBase = env.Float64("BACKOFF_EXPONENTIAL_BASE", false, 1, "Exponential base number to calculate the backoff")

func main() {
	logger = hclog.Default()
	err := env.Parse()
	if err != nil {
		logger.Error("Error parsing flags", "error", err)
		os.Exit(1)
	}

	closer, err := hckit.InitGlobalTracer("product-api")
	if err != nil {
		logger.Error("Unable to initialize Tracer", "error", err)
		os.Exit(1)
	}
	defer closer.Close()

	conf = &Config{
		DBConnection:           *dbConnection,
		BindAddress:            *bindAddress,
		MetricsAddress:         *metricsAddress,
		MaxRetries:             *maxRetries,
		BackoffExponentialBase: *backoffExponentialBase,
	}
	if conf.DBConnection == "" || conf.BindAddress == "" {
		//config.New
	}

}
