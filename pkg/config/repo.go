package config

type RepoConfig struct {
	*DBRepoConfig
	*CacheRepoConfig
	*ESRepoConfig
}

type DBRepoConfig struct {
	DBHost string
	DBUser string
	DBPass string
	DBName string
}

type CacheRepoConfig struct {
	CacheHost         string
	CacheUser         string
	CachePass         string
	CacheReadTimeout  int
	CacheWriteTimeout int
	CacheMaxIdle      int // 100
	CacheMaxActive    int //12000
	CacheIdleTimeout  int // 180
}

type ESRepoConfig struct {
	ESHost string
}
