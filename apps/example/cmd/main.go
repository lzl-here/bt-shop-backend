package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	pgens "github.com/lzl-here/bt-shop-backend/kitex_gen/example/exampleservice"
	"github.com/lzl-here/bt-shop-backend/pkg/config"
	api "github.com/lzl-here/bt-shop-backend/services/example/internal/api"
	"github.com/lzl-here/bt-shop-backend/services/example/internal/repo"
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
		slog.Error("无法打开或创建日志文件", "err", err)
		return
	}
	defer logFile.Close()

	logger := slog.New(slog.NewJSONHandler(logFile, nil))
	slog.SetDefault(logger)

	slog.Info(fmt.Sprintf("%s 启动", appConfig.ServiceName), "key", "value")

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
	r, err := etcd.NewEtcdRegistry([]string{appConfig.RegisterAddress})
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", appConfig.ServiceAddress)
	if err != nil {
		panic(err)
	}

	srv := pgens.NewServer(api.NewExampleServer(rep),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: appConfig.ServiceName}), // server name
		server.WithServiceAddr(addr), // address
		server.WithRegistry(r),       // registry
	)

	err = srv.Run()
	if err != nil {
		slog.Error("服务启动失败", "err", err)
	}

}
