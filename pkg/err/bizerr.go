package err

import "github.com/cloudwego/kitex/pkg/kerrors"

// 通用异常
var (
	ErrResourceNotFound = kerrors.NewBizStatusError(10001, "资源不存咋")
	ErrInternal         = kerrors.NewBizStatusError(10002, "服务端内部异常")
	ErrBadRequest       = kerrors.NewBizStatusError(10003, "请求参数错误")
	ErrDownStream       = kerrors.NewBizStatusError(10004, "调用第三方服务失败")
	ErrDuplicateReq     = kerrors.NewBizStatusError(1005, "重复请求")
)

// 支付异常
var (
	ErrPayFailed = kerrors.NewBizStatusError(20001, "支付失败")
)

// 商品异常
var (
	ErrGoodsNotExist      = kerrors.NewBizStatusError(30001, "商品不存在")
	ErrGoodsNotCorrect    = kerrors.NewBizStatusError(30002, "商品信息不符")
	ErrInvalidTradeAmount = kerrors.NewBizStatusError(30003, "交易金额有误")
)
