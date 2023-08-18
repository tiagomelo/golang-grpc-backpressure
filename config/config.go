package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// Config holds all configuration needed by this app.
type Config struct {
	StockServiceMetricsServerPort       int    `envconfig:"STOCK_SERVICE_METRICS_SERVER_PORT" required:"true"`
	StockServiceClientMetricsServerPort int    `envconfig:"STOCK_SERVICE_CLIENT_METRICS_SERVER_PORT" required:"true"`
	PromTemplateFile                    string `envconfig:"PROM_TEMPLATE_FILE" required:"true"`
	PromOutputFile                      string `envconfig:"PROM_OUTPUT_FILE" required:"true"`
	PromTargetGrpcServerPort            int    `envconfig:"PROM_TARGET_GRPC_SERVER_PORT" required:"true"`
	PromTargetGrpcClientPort            int    `envconfig:"PROM_TARGET_GRPC_CLIENT_PORT" required:"true"`
	DsTemplateFile                      string `envconfig:"DS_TEMPLATE_FILE" required:"true"`
	DsOutputFile                        string `envconfig:"DS_OUTPUT_FILE" required:"true"`
	DsServerPort                        int    `envconfig:"DS_SERVER_PORT" required:"true"`
}

// For ease of unit testing.
var (
	godotenvLoad     = godotenv.Load
	envconfigProcess = envconfig.Process
)

// Read reads the environment variables from the given file and returns a Config.
func Read() (*Config, error) {
	if err := godotenvLoad(); err != nil {
		return nil, errors.Wrap(err, "loading env vars")
	}
	config := new(Config)
	if err := envconfigProcess("", config); err != nil {
		return nil, errors.Wrap(err, "processing env vars")
	}
	return config, nil
}
