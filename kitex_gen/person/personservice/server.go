// Code generated by Kitex v0.12.1. DO NOT EDIT.
package personservice

import (
	server "github.com/cloudwego/kitex/server"
	person "github.com/lzl-here/bt-shop-backend/kitex_gen/person"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler person.PersonService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler person.PersonService, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}