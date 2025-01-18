package config

// 网关配置=服务+网关
type AppGatewayConfig struct {
	*AppServiceConfig
	*GatewayConfig
}

// 服务配置=数据+服务
type AppServiceConfig struct {
	*RepoConfig
	*ServiceConfig
	*CosConfig
	*LogConfig
	*RegisterConfig
}

// 网关配置
type GatewayConfig struct {
	GatewayHttpPort int `json:"http_port"` // 网关对外暴露的http端口
}

// 对象存储配置
type CosConfig struct {
	CosBucket    string `json:"cos_bucket"`
	CosHost      string `json:"cos_host"`
	CosSecretID  string `json:"cos_secret_id"`
	CosSecretKey string `json:"cos_secret_key"`
}

type LogConfig struct {
	LogLevel string `json:"log_level"`
	LogPath  string `json:"log_path"`
}

type RegisterConfig struct {
	RegisterAddress string `json:"register_address"`
	RegisterUser    string `json:"register_user"`
	RegisterPass    string `json:"register_pass"`
}

// 服务配置
type ServiceConfig struct {
	ServiceID       int    `json:"id"`        //服务实例唯一id
	ServiceName     string `json:"name"`      // 服务名称
	ServiceDesc     string `json:"desc"`      // 服务描述
	ServiceAddress  string `json:"address"`   // 地址
}
