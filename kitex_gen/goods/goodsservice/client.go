// Code generated by Kitex v0.12.1. DO NOT EDIT.

package goodsservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	goods "github.com/lzl-here/bt-shop-backend/kitex_gen/goods"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetGoodsList(ctx context.Context, Req *goods.GetGoodsListReq, callOptions ...callopt.Option) (r *goods.GetGoodsListRsp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kGoodsServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kGoodsServiceClient struct {
	*kClient
}

func (p *kGoodsServiceClient) GetGoodsList(ctx context.Context, Req *goods.GetGoodsListReq, callOptions ...callopt.Option) (r *goods.GetGoodsListRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetGoodsList(ctx, Req)
}
