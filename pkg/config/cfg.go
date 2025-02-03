package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadService(cfgFile string) *AppServiceConfig {
	if cfgFile == "" {
		panic("配置文件路径不能为空")
	}
	if err := loadFromFile(cfgFile); err != nil {
		panic(err)
	}

	return &AppServiceConfig{
		&RepoConfig{
			&DBRepoConfig{
				DBHost: GetStrEnvOrDefault("DB_HOST", "localhost:3306"),
				DBUser: GetStrEnvOrDefault("DB_USER", "root"),
				DBPass: GetStrEnvOrDefault("DB_PASS", "root"),
				DBName: GetStrEnvOrDefault("DB_TABLE_NAME", "test_db"),
			},
			&CacheRepoConfig{
				CacheHost:         GetStrEnvOrDefault("CACHE_HOST", "localhost"),
				CacheUser:         GetStrEnvOrDefault("CACHE_USER", ""),
				CachePass:         GetStrEnvOrDefault("CACHE_PASS", ""),
				CacheReadTimeout:  GetIntEnvOrDefault("CACHE_READ_TIMEOUT", 5),
				CacheWriteTimeout: GetIntEnvOrDefault("CACHE_WRITE_TIMEOUT", 5),
				CacheMaxIdle:      GetIntEnvOrDefault("CACHE_MAX_IDLE", 100),
				CacheMaxActive:    GetIntEnvOrDefault("CACHE_MAX_ACTIVE", 12000),
				CacheIdleTimeout:  GetIntEnvOrDefault("CACHE_IDLE_TIMEOUT", 180),
			},
			&ESRepoConfig{
				ESHost: GetStrEnvOrDefault("ES_HOST", "http://localhost:9200"),
			},
		},
		&ServiceConfig{
			ServiceAddress: GetStrEnvOrDefault("SERVICE_ADDRESS", "localhost:9090"),
			ServiceDesc:    GetStrEnvOrDefault("SERVICE_DESC", "暂无描述"),
			ServiceID:      GetIntEnvOrDefault("SERVICE_ID", -1),
			ServiceName:    GetStrEnvOrDefault("SERVICE_NAME", "暂无名称"),
		},
		&CosConfig{
			CosBucket:    GetStrEnvOrDefault("COS_BUCKET", ""),
			CosHost:      GetStrEnvOrDefault("COS_HOST", ""),
			CosSecretID:  GetStrEnvOrDefault("COS_SECRET_ID", ""),
			CosSecretKey: GetStrEnvOrDefault("COS_SECRET_KEY", ""),
		},
		&LogConfig{
			LogLevel: GetStrEnvOrDefault("LOG_LEVEL", "debug"),
			LogPath:  GetStrEnvOrDefault("LOG_PATH", "./logs"),
		},
		&RegisterConfig{
			RegisterAddress: GetStrEnvOrDefault("REGISTER_ADDRESS", "127.0.0.1:2379"),
			RegisterUser:    GetStrEnvOrDefault("REGISTER_USER", "root"),
			RegisterPass:    GetStrEnvOrDefault("REGISTER_PASS", "root"),
		},
	}
}

func LoadGateway(cfgFile string) *AppGatewayConfig {
	asc := LoadService(cfgFile)
	agc := &AppGatewayConfig{
		AppServiceConfig: asc,
		GatewayConfig: &GatewayConfig{
			GatewayHttpPort: GetStrEnvOrDefault("GATEWAY_HTTP_PORT", "8080"),
		},
	}
	return agc
}

func loadFromFile(cfg string) error {
	if err := godotenv.Load(cfg); err != nil {
		return err
	}
	return nil
}

func GetStrEnvOrDefault(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func GetIntEnvOrDefault(key string, defaultValue int) int {
	if value, ok := os.LookupEnv(key); ok {
		if v, err := strconv.Atoi(value); err == nil {
			return v
		}
	}
	return defaultValue
}

func GetFloatEnvOrDefault(key string, defaultValue float64) float64 {
	if value, ok := os.LookupEnv(key); ok {
		if v, err := strconv.ParseFloat(value, 64); err == nil {
			return v
		}
	}
	return defaultValue
}

func GetBoolEnvOrDefault(key string, defaultValue bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if v, err := strconv.ParseBool(value); err == nil {
			return v
		}
	}
	return defaultValue
}
