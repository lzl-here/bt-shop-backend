package repo

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/lzl-here/bt-shop-backend/apps/order/internal/domain/model"
	oc "github.com/lzl-here/bt-shop-backend/kitex_gen/order/orderservice"
	pc "github.com/lzl-here/bt-shop-backend/kitex_gen/pay/payservice"
	uc "github.com/lzl-here/bt-shop-backend/kitex_gen/user/userservice"

	ggen "github.com/lzl-here/bt-shop-backend/kitex_gen/goods"
	gc "github.com/lzl-here/bt-shop-backend/kitex_gen/goods/goodsservice"
	pgen "github.com/lzl-here/bt-shop-backend/kitex_gen/pay"
)

// 在mock数据时，在newServer的时候替换成mockImpl
type RepoInterface interface {
	GetCache() *redis.Client
	GetDB() *gorm.DB

	GetGoodsList(ctx context.Context, req *ggen.GetGoodsListReq) (*ggen.GetGoodsListRsp, error) // 商品列表
	ReduceStock(ctx context.Context, req *ggen.StockReduceReq) (*ggen.StockReduceRsp, error)    // 扣减库存

	Pay(ctx context.Context, req *pgen.PayReq) (*pgen.PayRsp, error) //支付

	CreateTrade(ctx context.Context, tx *gorm.DB, create *model.Trade) error
	CreateOrders(ctx context.Context, tx *gorm.DB, create []*model.Order) error
	CreateOrderItems(ctx context.Context, tx *gorm.DB, create []*model.OrderItem) error
	CreateTradeInfo(ctx context.Context, tx *gorm.DB, create *model.TradeInfo) error

	UpdateTrade(ctx context.Context, tx *gorm.DB, where, update *model.Trade) error
	UpdateOrder(ctx context.Context, tx *gorm.DB, where, update *model.Order) error

	GetOrderItems(ctx context.Context, where *model.OrderItem) ([]*model.OrderItem, error)
	GetTradeInfo(ctx context.Context, where *model.TradeInfo) (*model.TradeInfo, error)
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

func (r *Repo) GetGoodsList(ctx context.Context, req *ggen.GetGoodsListReq) (*ggen.GetGoodsListRsp, error) {
	return r.GoodsClient.GetGoodsList(ctx, req)
}

func (r *Repo) CreateTrade(ctx context.Context, tx *gorm.DB, create *model.Trade) error {
	t := new(model.Trade)
	return tx.Model(t).Create(create).Error
}

func (r *Repo) CreateOrders(ctx context.Context, tx *gorm.DB, create []*model.Order) error {
	return tx.Model(&model.Order{}).Create(create).Error
}
func (r *Repo) CreateOrderItems(ctx context.Context, tx *gorm.DB, create []*model.OrderItem) error {
	return tx.Model(&model.OrderItem{}).Create(create).Error
}

func (r *Repo) CreateTradeInfo(ctx context.Context, tx *gorm.DB, create *model.TradeInfo) error {
	return tx.Model(&model.TradeInfo{}).Create(create).Error
}

func (r *Repo) Pay(ctx context.Context, req *pgen.PayReq) (*pgen.PayRsp, error) {
	return r.PayClient.Pay(ctx, req)
}

func (r *Repo) GetOrderItems(ctx context.Context, where *model.OrderItem) ([]*model.OrderItem, error) {
	var items []*model.OrderItem
	err := r.DB.Model(&model.OrderItem{}).Where(where).Find(&items).Error
	return items, err
}

func (r *Repo) UpdateTrade(ctx context.Context, tx *gorm.DB, where, update *model.Trade) error {
	return tx.Model(&model.Trade{}).Where(where).Updates(update).Error
}
func (r *Repo) UpdateOrder(ctx context.Context, tx *gorm.DB, where, update *model.Order) error {
	return tx.Model(&model.Order{}).Where(where).Updates(update).Error
}

func (r *Repo) ReduceStock(ctx context.Context, req *ggen.StockReduceReq) (*ggen.StockReduceRsp, error) {
	return r.GoodsClient.StockReduce(ctx, req)
}

func (r *Repo) GetTradeInfo(ctx context.Context, where *model.TradeInfo) (*model.TradeInfo, error) {
	var err error
	res := new(model.TradeInfo)
	err = r.DB.Model(&model.TradeInfo{}).Where(where).Find(res).Error
	return res, err
}
