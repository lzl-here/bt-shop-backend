package main

import (
	"context"
	"flag"
	"net"
	"os"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	api "github.com/lzl-here/bt-shop-backend/apps/pay/internal/api"
	"github.com/lzl-here/bt-shop-backend/apps/pay/internal/config"
	repo "github.com/lzl-here/bt-shop-backend/apps/pay/internal/repo"
	"github.com/lzl-here/bt-shop-backend/pkg/middleware"

	pgens "github.com/lzl-here/bt-shop-backend/kitex_gen/pay/payservice"
	"gorm.io/gorm"

	"github.com/smartwalle/alipay/v3"
)

func main() {
	ctx := context.Background()
	var cfgFile string
	flag.StringVar(&cfgFile, "cfgFile", "", "配置文件路径")
	flag.Parse()
	// 读取配置文件
	config.AppConfig = config.LoadPayService(cfgFile)
	// 初始化日志
	logFile, err := os.OpenFile(config.AppConfig.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		klog.Error("无法打开或创建日志文件", "err", err)
		return
	}
	defer logFile.Close()

	klog.SetOutput(logFile)
	klog.SetLevel(klog.LevelInfo)

	// 注册中心
	r, err := etcd.NewEtcdRegistry([]string{config.AppConfig.RegisterAddress},
		etcd.WithAuthOpt(config.AppConfig.RegisterUser, config.AppConfig.RegisterPass),
	)
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", config.AppConfig.ServiceAddress)
	if err != nil {
		panic(err)
	}

	// 初始化数据访问层
	rep := newRepo(ctx)

	// 初始化server
	srv := pgens.NewServer(api.NewPayServer(rep),
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.AppConfig.ServiceName}), // server name
		server.WithServiceAddr(addr), // address
		server.WithRegistry(r),       // registry
		server.WithMiddleware(endpoint.Chain(middleware.LogMiddleWare, middleware.ErrorMiddleWare)),
	)

	err = srv.Run()
	if err != nil {
		klog.Error("服务启动失败", "err", err)
	}

}

func newRepo(ctx context.Context) *repo.Repo {

	// 连接db
	db, err := config.AppConfig.ConnectDB(&gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic(err)
	}

	// 连接缓存
	cache, err := config.AppConfig.ConnectCache(ctx)
	if err != nil {
		panic(err)
	}
	// 初始化支付宝client
	var alipayClient *alipay.Client
	if alipayClient, err = alipay.New(config.AppConfig.AlipayAppID, config.AppConfig.AlipayPrivateKey, false); err != nil {
		panic(err)
	}
	// 公私钥模式
	if err = alipayClient.LoadAliPayPublicKey(config.AppConfig.AlipayPublicKey); err != nil {
		panic(err)
	}
	//接口内容加密
	if err = alipayClient.SetEncryptKey(config.AppConfig.AlipayEncryptKey); err != nil {
		panic(err)
	}
	// grpc client
	payClient, goodsClient, orderClient, userClient := config.AppConfig.ConnectGrpcClient()

	return repo.NewRepo(db, cache, alipayClient, payClient, orderClient, goodsClient, userClient)
}
