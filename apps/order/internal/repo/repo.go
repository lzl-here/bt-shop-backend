package repo

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	gc "github.com/lzl-here/bt-shop-backend/kitex_gen/goods/goodsservice"
	oc "github.com/lzl-here/bt-shop-backend/kitex_gen/order/orderservice"
	pc "github.com/lzl-here/bt-shop-backend/kitex_gen/pay/payservice"
	uc "github.com/lzl-here/bt-shop-backend/kitex_gen/user/userservice"
)

// 在mock数据时，在newServer的时候替换成mockImpl
type RepoInterface interface {
	GetCache() *redis.Client
	GetDB() *gorm.DB
}

// 数据访问层实现了RepoInterface
type Repo struct {
	DB          *gorm.DB
	Cache       *redis.Client
	PayClient   *pc.Client
	OrderClient *oc.Client
	GoodsClient *gc.Client
	UserClient  *uc.Client
}

func NewRepo(db *gorm.DB, cache *redis.Client, payClient *pc.Client, orderClient *oc.Client, goodsClient *gc.Client, userClient *uc.Client) *Repo {
	return &Repo{
		DB:          db,
		Cache:       cache,
		PayClient:   payClient,
		OrderClient: orderClient,
		GoodsClient: goodsClient,
		UserClient:  userClient,
	}
}

func (r *Repo) GetCache() *redis.Client {
	return r.Cache
}

func (r *Repo) GetDB() *gorm.DB {
	return r.DB
}
