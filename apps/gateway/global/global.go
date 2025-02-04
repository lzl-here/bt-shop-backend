package global

import (
	gc "github.com/lzl-here/bt-shop-backend/kitex_gen/goods/goodsservice"
	oc "github.com/lzl-here/bt-shop-backend/kitex_gen/order/orderservice"
	pc "github.com/lzl-here/bt-shop-backend/kitex_gen/pay/payservice"
	uc "github.com/lzl-here/bt-shop-backend/kitex_gen/user/userservice"
	cfg "github.com/lzl-here/bt-shop-backend/pkg/config"
	"github.com/smartwalle/alipay/v3"
)

var AlipayClient *alipay.Client

var AppConfig *cfg.AppGatewayConfig

var PayClient pc.Client
var GoodsClient gc.Client
var OrderClient oc.Client
var UserClient uc.Client
