package repo

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/lzl-here/bt-shop-backend/apps/order/internal/domain/model"
	oc "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/order/orderservice"
	pc "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/pay/payservice"
	uc "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/user/userservice"

	ggen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/goods"
	gc "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/goods/goodsservice"
	pgen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/pay"
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

	UpdateTrade(ctx context.Context, tx *gorm.DB, where, update *model.Trade) error
	UpdateOrder(ctx context.Context, tx *gorm.DB, where, update *model.Order) error

	GetTrade(ctx context.Context, where *model.Trade) (*model.Trade, error)
	CountTrade(ctx context.Context, where *model.Trade) (int64, error)
	PageTradeList(ctx context.Context, where *model.Trade, pageSize, pageNum int) ([]*model.Trade, error)
	GetOrder(ctx context.Context, where *model.Order) (*model.Order, error)
	GetOrderList(ctx context.Context, where *model.Order) ([]*model.Order, error)
	GetOrderListByTradeNo(ctx context.Context, tradeNo []string) ([]*model.Order, error)
	GetOrderItems(ctx context.Context, where *model.OrderItem) ([]*model.OrderItem, error)
	GetOrderItemsByOrderNo(ctx context.Context, orderNo []string) ([]*model.OrderItem, error)
	GetOrderItemsByTradeNo(ctx context.Context, tradeNo []string) ([]*model.OrderItem, error)
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

func (r *Repo) GetTrade(ctx context.Context, where *model.Trade) (*model.Trade, error) {
	res := &model.Trade{}
	err := r.DB.Model(&model.Trade{}).Where(where).First(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *Repo) GetOrder(ctx context.Context, where *model.Order) (*model.Order, error) {
	res := &model.Order{}
	err := r.DB.Model(&model.Order{}).Where(where).First(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Repo) GetOrderList(ctx context.Context, where *model.Order) ([]*model.Order, error) {
	res := []*model.Order{}
	err := r.DB.Model(&model.Order{}).Where(where).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Repo) PageTradeList(ctx context.Context, where *model.Trade, pageSize, pageNum int) ([]*model.Trade, error) {
	res := []*model.Trade{}
	err := r.DB.Model(&model.Trade{}).Where(where).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Repo) GetOrderListByTradeNo(ctx context.Context, tradeNo []string) ([]*model.Order, error) {
	res := []*model.Order{}
	err := r.DB.Model(&model.Order{}).Where("trade_no IN (?)", tradeNo).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Repo) GetOrderItemsByOrderNo(ctx context.Context, orderNo []string) ([]*model.OrderItem, error) {
	res := []*model.OrderItem{}
	err := r.DB.Model(&model.OrderItem{}).Where("order_no IN (?)", orderNo).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Repo) GetOrderItemsByTradeNo(ctx context.Context, tradeNo []string) ([]*model.OrderItem, error) {
	res := []*model.OrderItem{}
	err := r.DB.Model(&model.OrderItem{}).Where("trade_no IN (?)", tradeNo).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Repo) CountTrade(ctx context.Context, where *model.Trade) (int64, error) {
	var count int64
	err := r.DB.Model(&model.Trade{}).Where(where).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
