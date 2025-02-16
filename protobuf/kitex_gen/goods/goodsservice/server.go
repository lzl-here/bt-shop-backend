// Code generated by Kitex v0.12.1. DO NOT EDIT.
package goodsservice

import (
	server "github.com/cloudwego/kitex/server"
	goods "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/goods"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler goods.GoodsService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler goods.GoodsService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
