// Code generated by hertz generator.

package main

import (
	"flag"

	"github.com/cloudwego/hertz/pkg/app/server"

	"github.com/lzl-here/bt-shop-backend/apps/gateway/global"
	"github.com/lzl-here/bt-shop-backend/pkg/config"
)

func main() {

	// config
	var cfgFile string
	flag.StringVar(&cfgFile, "cfgFile", "", "配置文件路径")
	flag.Parse()
	// 读取配置文件
	global.AppConfig = config.LoadGateway(cfgFile)

	// 连接到注册中心
	// pay client
	pc, gc, oc, uc := global.AppConfig.ConnectGrpcClient()
	global.PayClient = pc
	global.GoodsClient = gc
	global.OrderClient = oc
	global.UserClient = uc

	h := server.Default(
		server.WithHostPorts(global.AppConfig.GatewayHttpPort),
	)

	register(h)
	h.Spin()

}
