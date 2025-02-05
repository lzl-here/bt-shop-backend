package config

import (
	cfg "github.com/lzl-here/bt-shop-backend/pkg/config"
)

type UserConfig struct {
	cfg.AppServiceConfig
}

func LoadUserConfig(cfgFile string) *UserConfig {
	srvConfig := cfg.LoadService(cfgFile)
	p := &UserConfig{
		AppServiceConfig: *srvConfig,
	}
	return p
}
