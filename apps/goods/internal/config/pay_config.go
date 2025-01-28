package config

import (
	cfg "github.com/lzl-here/bt-shop-backend/pkg/config"
)

type GoodsConfig struct {
	cfg.AppServiceConfig
}

func LoadGoodsConfig(cfgFile string) *GoodsConfig {
	srvConfig := cfg.LoadService(cfgFile)
	p := &GoodsConfig{
		AppServiceConfig: *srvConfig,
	}
	return p
}
