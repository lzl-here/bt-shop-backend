package middleware

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	bizerr "github.com/lzl-here/bt-shop-backend/pkg/err"
)

// ErrorMiddleWare 错误处理中间件
// 确保业务上返回给客户端的err一定是kitex的bizStatusError
func ErrorMiddleWare(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) error {
		// handler返回异常：

		// 1. 返回其他异常, 包装成internal error返回给客户端
		err := next(ctx, req, resp)
		if err != nil {
			klog.Error(err)
			bizErr := bizerr.ErrInternalError
			ri := rpcinfo.GetRPCInfo(ctx)
			if setter, ok := ri.Invocation().(rpcinfo.InvocationSetter); ok {
				setter.SetBizStatusErr(bizErr)
				klog.Error(bizErr)
			}
			return nil
		}

		// 2. 返回kitex的bizStatusErr异常，不会在上面的next中返回，只能通过rpcinfo获取，直接返回给客户端
		bizErr := rpcinfo.GetRPCInfo(ctx).Invocation().BizStatusErr()
		klog.Error(bizErr)

		return nil
	}
}
