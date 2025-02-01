package err

import "github.com/cloudwego/kitex/pkg/kerrors"

var (
	ErrResourceNotFound = kerrors.NewBizStatusError(10001, "资源不存咋")
	ErrInternal         = kerrors.NewBizStatusError(10002, "服务端内部异常")
	ErrBadRequest       = kerrors.NewBizStatusError(10003, "请求参数错误")
	ErrDownStream       = kerrors.NewBizStatusError(10004, "调用第三方服务失败")
	ErrDuplicateReq     = kerrors.NewBizStatusError(1005, "重复请求")
)
