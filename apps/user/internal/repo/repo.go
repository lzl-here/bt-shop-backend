package repo

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/lzl-here/bt-shop-backend/apps/user/internal/domain/model"
	oc "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/order/orderservice"
	pc "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/pay/payservice"
	uc "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/user/userservice"

	gc "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/goods/goodsservice"
)

// 在mock数据时，在newServer的时候替换成mockImpl
type RepoInterface interface {
	GetCache() *redis.Client
	GetDB() *gorm.DB
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)

	GetUser(ctx context.Context, where *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
}

var _ RepoInterface = (*Repo)(nil)

// 数据访问层实现了RepoInterface
type Repo struct {
	DB          *gorm.DB
	Cache       *redis.Client
	PayClient   pc.Client
	OrderClient oc.Client
	GoodsClient gc.Client
	UserClient  uc.Client
}

func NewRepo(db *gorm.DB, cache *redis.Client, payClient pc.Client, orderClient oc.Client, goodsClient gc.Client, userClient uc.Client) *Repo {
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

func (r *Repo) GetUser(ctx context.Context, where *model.User) (*model.User, error) {
	var user model.User
	if err := r.DB.Where(where).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repo) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	err := r.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repo) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	err := r.DB.Model(&model.User{}).Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
