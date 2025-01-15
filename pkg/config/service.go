package config

// 服务配置
type ServiceConfig struct {
	ServiceID       int    `json:"id"`        //服务实例唯一id
	ServiceName     string `json:"name"`      // 服务名称
	ServiceDesc     string `json:"desc"`      // 服务描述
	ServiceAddress  string `json:"address"`   // 地址
}
