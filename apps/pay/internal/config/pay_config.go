package config

import (
	cfg "github.com/lzl-here/bt-shop-backend/pkg/config"
)

type PayConfig struct {
	cfg.AppServiceConfig
	AlipayAppID        string // 支付宝appID
	AlipayPrivateKey   string // 支付宝应用私钥
	AlipayServerUrl string // 支付宝回调服务端地址
	AlipayFrontUrl  string // 支付宝前端跳转地址
	AlipayEncryptKey   string // 支付宝加密密钥
	AlipayPublicKey    string // 支付宝公钥
}

func LoadPayService(cfgFile string) *PayConfig {
	p := &PayConfig{}
	srvConfig := cfg.LoadService(cfgFile)
	p.AppServiceConfig = *srvConfig

	// 支付宝配置
	p.AlipayAppID = cfg.GetStrEnvOrDefault("PAY_ALIPAY_APPID", "")
	p.AlipayPrivateKey = cfg.GetStrEnvOrDefault("PAY_ALIPAY_PRIVATE_KEY", "")
	p.AlipayServerUrl = cfg.GetStrEnvOrDefault("PAY_SALIPAY_ERVER_URL", "")
	p.AlipayFrontUrl = cfg.GetStrEnvOrDefault("PAY_ALIPAY_FRONT_URL", "")
	p.AlipayEncryptKey = cfg.GetStrEnvOrDefault("PAY_ALIPAY_ENCRYPT_KEY", "")
	p.AlipayPublicKey = cfg.GetStrEnvOrDefault("PAY_ALIPAY_PUBLIC_KEY", "")
	return p
}
