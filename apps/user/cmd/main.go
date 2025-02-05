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
	api "github.com/lzl-here/bt-shop-backend/apps/user/internal/api"
	"github.com/lzl-here/bt-shop-backend/apps/user/internal/config"
	repo "github.com/lzl-here/bt-shop-backend/apps/user/internal/repo"
	"github.com/lzl-here/bt-shop-backend/pkg/middleware"

	ugens "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/user/userservice"
	"gorm.io/gorm"
)

func main() {
	ctx := context.Background()
	var cfgFile string
	flag.StringVar(&cfgFile, "cfgFile", "", "配置文件路径")
	flag.Parse()
	// 读取配置文件
	config.AppConfig = config.LoadUserConfig(cfgFile)
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
	srv := ugens.NewServer(api.NewUserServer(rep),
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

	// grpc client
	payClient, goodsClient, orderClient, userClient := config.AppConfig.ConnectGrpcClient()

	return repo.NewRepo(db, cache, payClient, orderClient, goodsClient, userClient)
}
