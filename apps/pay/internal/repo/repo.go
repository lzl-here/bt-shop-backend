package repo

import (
	"net/url"

	"github.com/lzl-here/bt-shop-backend/apps/pay/internal/config"
	"github.com/lzl-here/bt-shop-backend/apps/pay/internal/domain/model"
	"github.com/redis/go-redis/v9"
	"github.com/smartwalle/alipay/v3"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 在mock数据时，在newServer的时候替换成mockImpl
type RepoInterface interface {
	AlipayPay(alipay.TradePagePay) (*url.URL, error)
	CreatePayFlow(payFlow *model.PayFlow) (*model.PayFlow, error)
}

// 数据访问层实现了RepoInterface
type Repo struct {
	DB     *gorm.DB
	Cache  *redis.Client
	Alipay *alipay.Client
}

func NewRepo(db *gorm.DB, cache *redis.Client, alipay *alipay.Client) *Repo {
	return &Repo{
		DB:     db,
		Cache:  cache,
		Alipay: alipay,
	}
}

// 支付宝支付
func (r *Repo) AlipayPay(param alipay.TradePagePay) (*url.URL, error) {
	param.NotifyURL = config.AppConfig.AlipayServerDomain + "/pay/notify"
	param.ReturnURL = config.AppConfig.AlipayServerDomain + "/pay/return"
	param.ProductCode = "FAST_INSTANT_TRADE_PAY"
	payPageUrl, err := r.Alipay.TradePagePay(param)
	return payPageUrl, err
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
