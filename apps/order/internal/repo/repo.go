package repo

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// 在mock数据时，在newServer的时候替换成mockImpl
type RepoInterface interface {
	GetCache() *redis.Client
	GetDB() *gorm.DB
}

// 数据访问层实现了RepoInterface
type Repo struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func NewRepo(db *gorm.DB, cache *redis.Client) *Repo {
	return &Repo{
		DB:    db,
		Cache: cache,
	}
}

func (r *Repo) GetCache() *redis.Client {
	return r.Cache
}

func (r *Repo) GetDB() *gorm.DB {
	return r.DB
}
