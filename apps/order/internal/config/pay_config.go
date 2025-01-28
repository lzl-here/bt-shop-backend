package config

import (
	cfg "github.com/lzl-here/bt-shop-backend/pkg/config"
)

type OrderConfig struct {
	cfg.AppServiceConfig
}

func LoadOrderConfig(cfgFile string) *OrderConfig {
	srvConfig := cfg.LoadService(cfgFile)
	p := &OrderConfig{
		AppServiceConfig: *srvConfig,
	}
	return p
}
