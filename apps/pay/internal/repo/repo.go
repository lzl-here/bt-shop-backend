package repo

import (
	"context"
	"net/url"

	"github.com/lzl-here/bt-shop-backend/apps/pay/internal/config"
	"github.com/lzl-here/bt-shop-backend/apps/pay/internal/domain/model"
	"github.com/redis/go-redis/v9"
	"github.com/smartwalle/alipay/v3"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	gc "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/goods/goodsservice"
	oc "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/order/orderservice"
	pc "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/pay/payservice"
	uc "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/user/userservice"

	ogen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/order"
)

// 在mock数据时，在newServer的时候替换成mockImpl
type RepoInterface interface {
	GetDB() *gorm.DB
	GetCache() *redis.Client

	GetOrderItems(ctx context.Context, req *ogen.GetOrderItemsReq) (*ogen.GetOrderItemsRsp, error)             // 获取订单项
	PaySuccessToOrder(ctx context.Context, req *ogen.PaySuccessToOrderReq) (*ogen.PaySuccessToOrderRsp, error) //支付成功通知订单系统
	PayCancelToOrder(ctx context.Context, req *ogen.PayCancelToOrderReq) (*ogen.PayCancelToOrderRsp, error)    //支付取消通知订单系统

	AlipayPay(alipay.TradePagePay) (*url.URL, error)                                  // 拉起支付宝支付
	AlipayClose(context.Context, alipay.TradeClose) (*alipay.TradeCloseRsp, error)    // 支付宝关闭支付
	AlipayRefund(context.Context, alipay.TradeRefund) (*alipay.TradeRefundRsp, error) // 支付宝退款
	CreatePayFlows(payFlow []*model.PayFlow) error                                    // 创建支付流水
	UpdatePayFlow(where, update *model.PayFlow) (*model.PayFlow, error)               // 修改支付流水

	CreatePayInfo(ctx context.Context, tx *gorm.DB, payInfo *model.PayInfo) error
}

var _ RepoInterface = (*Repo)(nil)

// 数据访问层实现了RepoInterface
type Repo struct {
	DB          *gorm.DB
	Cache       *redis.Client
	Alipay      *alipay.Client
	PayClient   pc.Client
	OrderClient oc.Client
	GoodsClient gc.Client
	UserClient  uc.Client
}

func NewRepo(db *gorm.DB, cache *redis.Client, alipay *alipay.Client, payClient pc.Client, orderClient oc.Client, goodsClient gc.Client, userClient uc.Client) *Repo {
	return &Repo{
		DB:          db,
		Cache:       cache,
		Alipay:      alipay,
		PayClient:   payClient,
		OrderClient: orderClient,
		GoodsClient: goodsClient,
		UserClient:  userClient,
	}
}

func (r *Repo) GetDB() *gorm.DB {
	return r.DB
}
func (r *Repo) GetCache() *redis.Client {
	return r.Cache
}

// 拉起支付宝支付
func (r *Repo) AlipayPay(param alipay.TradePagePay) (*url.URL, error) {
	param.NotifyURL = config.AppConfig.AlipayServerUrl // 异步回调连接
	param.ReturnURL = config.AppConfig.AlipayServerUrl // 支付成功跳转链接
	param.ProductCode = "FAST_INSTANT_TRADE_PAY"       // 这个写死不要动
	payPageUrl, err := r.Alipay.TradePagePay(param)
	return payPageUrl, err
}

// 拉起支付宝关闭交易
func (r *Repo) AlipayClose(ctx context.Context, param alipay.TradeClose) (*alipay.TradeCloseRsp, error) {
	param.NotifyURL = config.AppConfig.AlipayServerUrl // 异步回调连接
	return r.Alipay.TradeClose(ctx, param)
}

// 支付宝退款
func (r *Repo) AlipayRefund(ctx context.Context, param alipay.TradeRefund) (*alipay.TradeRefundRsp, error) {
	return r.Alipay.TradeRefund(ctx, param)
}

// 创建支付流水
func (r *Repo) CreatePayFlows(payFlow []*model.PayFlow) error {
	p := new(model.PayFlow)

	if err := r.DB.Model(p).Clauses(clause.Returning{}).
		Create(payFlow).Error; err != nil {
		return err
	}
	return nil
}

// 修改支付流水
func (r *Repo) UpdatePayFlow(where, update *model.PayFlow) (*model.PayFlow, error) {
	p := new(model.PayFlow)
	if err := r.DB.Model(p).Clauses(clause.Returning{}).
		Where(where).
		Updates(update).Error; err != nil {
		return nil, err
	}
	return p, nil
}

// 获取订单项
func (r *Repo) GetOrderItems(ctx context.Context, req *ogen.GetOrderItemsReq) (*ogen.GetOrderItemsRsp, error) {
	return r.OrderClient.GetOrderItems(ctx, req)
}

// 支付成功通知订单系统
func (r *Repo) PaySuccessToOrder(ctx context.Context, req *ogen.PaySuccessToOrderReq) (*ogen.PaySuccessToOrderRsp, error) {
	return r.OrderClient.PaySuccessToOrder(ctx, req)
}

// 支付取消通知订单系统
func (r *Repo) PayCancelToOrder(ctx context.Context, req *ogen.PayCancelToOrderReq) (*ogen.PayCancelToOrderRsp, error) {
	return r.OrderClient.PayCancelToOrder(ctx, req)
}

func (r *Repo) CreatePayInfo(ctx context.Context, tx *gorm.DB, payInfo *model.PayInfo) error {
	return tx.WithContext(ctx).Create(payInfo).Error
}
