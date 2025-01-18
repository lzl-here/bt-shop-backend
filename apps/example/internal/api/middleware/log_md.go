package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

// 日志中间件
// TODO log_id
// TODO 结构化打印 req, resp
func LogMiddleWare(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) error {
		rpcInfo := rpcinfo.GetRPCInfo(ctx)
		address := rpcInfo.From().Address()
		methodName := rpcInfo.Invocation().MethodName()
		serviceName := rpcInfo.Invocation().ServiceName()

		startTime := time.Now()
		// 打印请求路径和请求体
		klog.Infof("请求路径: %s, 请求地址: %s, 请求体: %v",
			serviceName+"."+methodName,
			address.String(),
			fmt.Sprintf("%+v", req),
		)

		err := next(ctx, req, resp)

		duration := time.Since(startTime)

		// 打印响应体和耗时时间
		klog.Infof("响应体: %v, 耗时: %vms", fmt.Sprintf("%+v", resp), duration.Milliseconds())

		return err
	}
}
