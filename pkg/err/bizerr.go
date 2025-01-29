package err

import "github.com/cloudwego/kitex/pkg/kerrors"

var (
	ErrResourceNotFound = kerrors.NewBizStatusError(10001, "资源不存咋")
	ErrInternalError    = kerrors.NewBizStatusError(10002, "服务端内部异常")
	ErrBadRequest       = kerrors.NewBizStatusError(10003, "请求参数错误")
	ErrDownStreamError  = kerrors.NewBizStatusError(10004, "调用第三方服务失败")
)
