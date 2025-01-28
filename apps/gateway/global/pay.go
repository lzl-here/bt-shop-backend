package global

import (
	"github.com/smartwalle/alipay/v3"
	cfg "github.com/lzl-here/bt-shop-backend/pkg/config"
	pc "github.com/lzl-here/bt-shop-backend/kitex_gen/pay/payservice"
)

var AlipayClient *alipay.Client

var AppConfig *cfg.AppGatewayConfig

var PayClient pc.Client