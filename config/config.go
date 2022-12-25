package config

import (
	"fmt"
	"github.com/caarlos0/env"
)

type Config struct {
	HTTP HTTPConfig
	GRPC GRPCConfig
}

type HTTPConfig struct {
	Address string `env:"HTTP_ADDRESS"`
}

func (tc HTTPConfig) String() string {
	return fmt.Sprintf("Addr: %s ", tc.Address)
}

type GRPCConfig struct {
	HOST string `env:"GRPC_HOST"`
	Port int    `env:"GRPC_PORT"`
}

func (tc GRPCConfig) String() string {
	return fmt.Sprintf("Port: %d Host: %s ", tc.Port, tc.HOST)
}
func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(&cfg.HTTP); err != nil {
		return nil, err
	}
	if err := env.Parse(&cfg.GRPC); err != nil {
		return nil, fmt.Errorf("fail when try load grpc configuration, err: %w", err)
	}
	return cfg, nil
}
