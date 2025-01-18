package config

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

