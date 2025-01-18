package repo

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/lzl-here/bt-shop-backend/apps/example/internal/domain/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// 在mock数据时，在newServer的时候替换成mockImpl
type RepoInterface interface {
	GetPersonByID(id uint64)
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

func (r *Repo) GetPersonByID(ctx context.Context, id uint64) (*model.Person, error) {
	p := new(model.Person)
	var str string
	var err error
	str, err = r.cache.Get(ctx, "person:"+strconv.Itoa(int(id))).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	// 缓存为空
	if err == redis.Nil || str == "" {
		err = r.db.Model(p).Where("id = ?", id).Find(p).Error
		if err != nil {
			return nil, err
		}

		// 写入缓存失败不应该影响返回值
		if err := r.cache.Set(ctx, "person:"+strconv.Itoa(int(id)), p, 0).Err(); err != nil {
			klog.Error("redis set error: %v", err)
		}
		return p, nil
	}

	// 缓存不为空
	if err = json.Unmarshal([]byte(str), p); err != nil {
		return nil, err
	}
	return p, nil
}
