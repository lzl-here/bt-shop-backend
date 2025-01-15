package config

import (
	"context"
	"fmt"
	"log/slog"

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
	slog.Info("连接数据库", "dsn", dsn)
	db, err := gorm.Open(mysql.Open(dsn), gormCfg)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (a *AppServiceConfig) ConnectCache(ctx context.Context) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     a.CacheHost,
		Password: a.CacheHost,
		MaxIdleConns: a.CacheMaxIdle,
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	slog.Info("连接redis成功", "pong", pong)
	return rdb, nil
}
