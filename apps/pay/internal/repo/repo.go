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

	gc "github.com/lzl-here/bt-shop-backend/kitex_gen/goods/goodsservice"
	oc "github.com/lzl-here/bt-shop-backend/kitex_gen/order/orderservice"
	pc "github.com/lzl-here/bt-shop-backend/kitex_gen/pay/payservice"
	uc "github.com/lzl-here/bt-shop-backend/kitex_gen/user/userservice"
)

// 在mock数据时，在newServer的时候替换成mockImpl
type RepoInterface interface {
	AlipayPay(alipay.TradePagePay) (*url.URL, error)
	AlipayClose(context.Context, alipay.TradeClose) (*alipay.TradeCloseRsp, error)
	AlipayRefund(context.Context, alipay.TradeRefund) (*alipay.TradeRefundRsp, error)
	CreatePayFlow(payFlow *model.PayFlow) (*model.PayFlow, error)
	UpdatePayFlow(payFlow *model.PayFlow) (*model.PayFlow, error)
}

// 数据访问层实现了RepoInterface
type Repo struct {
	DB          *gorm.DB
	Cache       *redis.Client
	Alipay      *alipay.Client
	PayClient   *pc.Client
	OrderClient *oc.Client
	GoodsClient *gc.Client
	UserClient  *uc.Client
}

func NewRepo(db *gorm.DB, cache *redis.Client, alipay *alipay.Client, payClient *pc.Client, orderClient *oc.Client, goodsClient *gc.Client, userClient *uc.Client) *Repo {
	return &Repo{
		DB:        db,
		Cache:     cache,
		Alipay:    alipay,
		PayClient: payClient,
		OrderClient: orderClient,
		GoodsClient: goodsClient,
		UserClient: userClient,
	}
}

// 支付宝支付
func (r *Repo) AlipayPay(param alipay.TradePagePay) (*url.URL, error) {
	param.NotifyURL = config.AppConfig.AlipayServerUrl // 异步回调连接
	param.ReturnURL = config.AppConfig.AlipayServerUrl // 支付成功跳转链接
	param.ProductCode = "FAST_INSTANT_TRADE_PAY"       // 这个写死不要动
	payPageUrl, err := r.Alipay.TradePagePay(param)
	return payPageUrl, err
}

// 支付宝关闭交易
func (r *Repo) AlipayClose(ctx context.Context, param alipay.TradeClose) (*alipay.TradeCloseRsp, error) {
	param.NotifyURL = config.AppConfig.AlipayServerUrl // 异步回调连接
	return r.Alipay.TradeClose(ctx, param)
}

// 支付宝退款
func (r *Repo) AlipayRefund(ctx context.Context, param alipay.TradeRefund) (*alipay.TradeRefundRsp, error) {
	return r.Alipay.TradeRefund(ctx, param)
}

// 创建支付流水
func (r *Repo) CreatePayFlow(payFlow *model.PayFlow) (*model.PayFlow, error) {
	p := new(model.PayFlow)
	if err := r.DB.Model(p).Clauses(clause.Returning{}).
		Create(payFlow).Error; err != nil {
		return nil, err
	}
	return p, nil
}

// 修改支付流水
func (r *Repo) UpdatePayFlow(payFlow *model.PayFlow) (*model.PayFlow, error) {
	p := new(model.PayFlow)
	if err := r.DB.Model(p).Clauses(clause.Returning{}).
		Where("trade_no = ?", payFlow.TradeNo).
		Updates(payFlow).Error; err != nil {
		return nil, err
	}
	return p, nil
}
