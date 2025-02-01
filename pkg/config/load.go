package config

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/transport"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	gc "github.com/lzl-here/bt-shop-backend/kitex_gen/goods/goodsservice"
	oc "github.com/lzl-here/bt-shop-backend/kitex_gen/order/orderservice"
	pc "github.com/lzl-here/bt-shop-backend/kitex_gen/pay/payservice"
	uc "github.com/lzl-here/bt-shop-backend/kitex_gen/user/userservice"
)

func (a *AppServiceConfig) ConnectDB(gormCfg *gorm.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		a.DBUser,
		a.DBPass,
		a.DBHost,
		a.DBName,
	)
	klog.Infof("连接数据库, dsn: %s", dsn)
	db, err := gorm.Open(mysql.Open(dsn), gormCfg)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (a *AppServiceConfig) ConnectCache(ctx context.Context) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         a.CacheHost,
		Password:     a.CacheHost,
		MaxIdleConns: a.CacheMaxIdle,
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	klog.Infof("连接到redis成功, pong: %s", pong)
	return rdb, nil
}

func (a *AppServiceConfig) ConnectGrpcClient() (*pc.Client, *gc.Client, *oc.Client, *uc.Client) {
	// grpc client
	r, err := etcd.NewEtcdResolver(
		[]string{a.RegisterAddress},
		etcd.WithAuthOpt(a.RegisterUser, a.RegisterPass),
	)
	if err != nil {
		panic(err)
	}
	payClient, err := pc.NewClient("pay", client.WithResolver(r),
		client.WithTransportProtocol(transport.GRPC),
	)
	if err != nil {
		panic(err)
	}
	goodsClient, err := gc.NewClient("goods", client.WithResolver(r),
		client.WithTransportProtocol(transport.GRPC),
	)
	if err != nil {
		panic(err)
	}

	orderClient, err := oc.NewClient("order", client.WithResolver(r),
		client.WithTransportProtocol(transport.GRPC),
	)
	if err != nil {
		panic(err)
	}

	userClient, err := uc.NewClient("user", client.WithResolver(r),
		client.WithTransportProtocol(transport.GRPC),
	)
	if err != nil {
		panic(err)
	}
	return &payClient, &goodsClient, &orderClient, &userClient
}
