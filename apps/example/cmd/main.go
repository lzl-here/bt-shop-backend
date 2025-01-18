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
	api "github.com/lzl-here/bt-shop-backend/apps/example/internal/api"
	"github.com/lzl-here/bt-shop-backend/apps/example/internal/api/middleware"
	"github.com/lzl-here/bt-shop-backend/apps/example/internal/repo"
	pgens "github.com/lzl-here/bt-shop-backend/kitex_gen/example/exampleservice"
	"github.com/lzl-here/bt-shop-backend/pkg/config"
	"gorm.io/gorm"
)

func main() {
	ctx := context.Background()
	var cfgFile string
	flag.StringVar(&cfgFile, "cfgFile", "", "配置文件路径")
	flag.Parse()
	// 读取配置文件
	appConfig := config.LoadService(cfgFile)

	// 初始化日志
	logFile, err := os.OpenFile(appConfig.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		klog.Error("无法打开或创建日志文件", "err", err)
		return
	}
	defer logFile.Close()

	klog.SetOutput(logFile)
	klog.SetLevel(klog.LevelInfo)

	// 连接中间件
	db, err := appConfig.ConnectDB(&gorm.Config{})
	if err != nil {
		panic(err)
	}

	cache, err := appConfig.ConnectCache(ctx)
	if err != nil {
		panic(err)
	}
	rep := repo.NewRepo(db, cache)
	// 启动rpc服务
	r, err := etcd.NewEtcdRegistry([]string{appConfig.RegisterAddress}, 
		etcd.WithAuthOpt(appConfig.RegisterUser, appConfig.RegisterPass),
	)
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", appConfig.ServiceAddress)
	if err != nil {
		panic(err)
	}

	srv := pgens.NewServer(api.NewExampleServer(rep),
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: appConfig.ServiceName}), // server name
		server.WithServiceAddr(addr),                                                               // address
		server.WithRegistry(r),                                                                     // registry
		server.WithMiddleware(endpoint.Chain(middleware.LogMiddleWare, middleware.ErrorMiddleWare)),
	)

	err = srv.Run()
	if err != nil {
		klog.Error("服务启动失败", "err", err)
	}

}
