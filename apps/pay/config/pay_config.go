package config

import (
	cfg "github.com/lzl-here/bt-shop-backend/pkg/config"
)

type PayConfig struct {
	cfg.AppServiceConfig
}

func LoadPayService(cfgFile string)(*PayConfig) {
	p := &PayConfig{}
	srvConfig := cfg.LoadService(cfgFile)
	p.AppServiceConfig = *srvConfig
	return p
}
