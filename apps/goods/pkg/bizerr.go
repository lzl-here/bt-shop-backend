package pkg

import "github.com/cloudwego/kitex/pkg/kerrors"

var (
	ErrResourceNotFound = kerrors.NewBizStatusError(10001, "资源不存咋")
	ErrInternalError    = kerrors.NewBizStatusError(10002, "服务端内部异常")
)

