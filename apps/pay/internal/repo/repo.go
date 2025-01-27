package repo

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// 在mock数据时，在newServer的时候替换成mockImpl
type RepoInterface interface {
}

type Repo struct {
	db    *gorm.DB
	cache *redis.Client
}

func NewRepo(db *gorm.DB, cache *redis.Client) *Repo {
	return &Repo{
		db:    db,
		cache: cache,
	}
}
