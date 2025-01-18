package middleware

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

func ErrorMiddleWare(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) error {
		err := next(ctx, req, resp)
		if err != nil {
			klog.Error(err)
			return err
		}
		err = rpcinfo.GetRPCInfo(ctx).Invocation().BizStatusErr()
		if err != nil {
			klog.Error(err)
		}
		return nil
	}
}
