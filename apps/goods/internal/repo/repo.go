package repo

import (
	"context"

	"github.com/lzl-here/bt-shop-backend/apps/goods/internal/domain/model"
	gc "github.com/lzl-here/bt-shop-backend/kitex_gen/goods/goodsservice"
	oc "github.com/lzl-here/bt-shop-backend/kitex_gen/order/orderservice"
	pc "github.com/lzl-here/bt-shop-backend/kitex_gen/pay/payservice"
	uc "github.com/lzl-here/bt-shop-backend/kitex_gen/user/userservice"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var _ (RepoInterface) = (*Repo)(nil)

// 在mock数据时，在newServer的时候替换成mockImpl
type RepoInterface interface {
	GetSpuListByIDs(ctx context.Context, IDs []uint64) ([]*model.GoodsSpu, error)
	GetSkuListBySpuIDs(ctx context.Context, SpuIDs []uint64) ([]*model.GoodsSku, error)
	GetSpecBySpuIDs(ctx context.Context, SpuIDs []uint64) ([]*model.Spec, error)
}

// 数据访问层实现了RepoInterface
type Repo struct {
	DB          *gorm.DB
	Cache       *redis.Client
	PayClient   pc.Client
	GoodsClient gc.Client
	OrderClient oc.Client
	UserClient  uc.Client
}

func NewRepo(db *gorm.DB, cache *redis.Client, payClient pc.Client, goodsClient gc.Client, orderClient oc.Client, userClient uc.Client) *Repo {
	return &Repo{
		DB:          db,
		Cache:       cache,
		PayClient:   payClient,
		GoodsClient: goodsClient,
		OrderClient: orderClient,
		UserClient:  userClient,
	}
}

func (r *Repo) GetSpuListByIDs(ctx context.Context, IDs []uint64) ([]*model.GoodsSpu, error) {
	var spus []*model.GoodsSpu
	err := r.DB.Model(&model.GoodsSpu{}).Where("id in ?", IDs).Find(spus).Error
	if err != nil {
		return nil, err
	}
	return spus, nil
}

func (r *Repo) GetSkuListBySpuIDs(ctx context.Context, SpuIDs []uint64) ([]*model.GoodsSku, error) {
	var skus []*model.GoodsSku
	err := r.DB.Model(&model.GoodsSku{}).Where("spu_id in ?", SpuIDs).Find(skus).Error
	if err != nil {
		return nil, err
	}
	return skus, nil
}

func (r *Repo) GetSpecBySpuIDs(ctx context.Context, SpuIDs []uint64) ([]*model.Spec, error) {
	var specs []*model.Spec
	err := r.DB.Model(&model.Spec{}).Where("spu_id in ?", SpuIDs).Find(specs).Error
	if err != nil {
		return nil, err
	}
	return specs, nil
}
