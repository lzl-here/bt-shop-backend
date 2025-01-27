package config

import (
	baseConfig "github.com/lzl-here/bt-shop-backend/pkg/config"
	cfg "github.com/lzl-here/bt-shop-backend/pkg/config"
)

type ExampleConfig struct {
	cfg.AppServiceConfig
	Something string // 自定义配置
}

// 加载通用配置和自定义配置
func LoadExampleService(cfgFile string) *ExampleConfig {
	p := &ExampleConfig{}
	srvConfig := cfg.LoadService(cfgFile)
	p.AppServiceConfig = *srvConfig
	p.loadExtraConfig()
	return p
}

// 加载服务的自定义配置
func (cfg *ExampleConfig) loadExtraConfig() {
	cfg.Something = baseConfig.GetStrEnvOrDefault("SOMETHING", "lalala")
}
