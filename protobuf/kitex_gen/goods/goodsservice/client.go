// Code generated by Kitex v0.12.1. DO NOT EDIT.

package goodsservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	goods "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/goods"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetKeywordDownList(ctx context.Context, Req *goods.GetKeywordDownListReq, callOptions ...callopt.Option) (r *goods.GetKeywordDownListRsp, err error)
	SearchSpuList(ctx context.Context, Req *goods.SearchSpuListReq, callOptions ...callopt.Option) (r *goods.SearchSpuListRsp, err error)
	GetGoodsDetail(ctx context.Context, Req *goods.GetGoodsDetailReq, callOptions ...callopt.Option) (r *goods.GetGoodsDetailRsp, err error)
	GetSellerGoodsList(ctx context.Context, Req *goods.GetSellerGoodsListReq, callOptions ...callopt.Option) (r *goods.GetSellerGoodsListRsp, err error)
	GetGoodsList(ctx context.Context, Req *goods.GetGoodsListReq, callOptions ...callopt.Option) (r *goods.GetGoodsListRsp, err error)
	GetCategoryList(ctx context.Context, Req *goods.GetCategoryListReq, callOptions ...callopt.Option) (r *goods.GetCategoryListRsp, err error)
	GetBrandList(ctx context.Context, Req *goods.GetBrandListReq, callOptions ...callopt.Option) (r *goods.GetBrandListRsp, err error)
	AddToCart(ctx context.Context, Req *goods.AddToCartReq, callOptions ...callopt.Option) (r *goods.AddToCartRsp, err error)
	PublishGoods(ctx context.Context, Req *goods.PublishGoodsReq, callOptions ...callopt.Option) (r *goods.PublishGoodsRsp, err error)
	StockReduce(ctx context.Context, Req *goods.StockReduceReq, callOptions ...callopt.Option) (r *goods.StockReduceRsp, err error)
	StockIncrease(ctx context.Context, Req *goods.StockIncreaseReq, callOptions ...callopt.Option) (r *goods.StockIncreaseRsp, err error)
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

func (p *kGoodsServiceClient) GetKeywordDownList(ctx context.Context, Req *goods.GetKeywordDownListReq, callOptions ...callopt.Option) (r *goods.GetKeywordDownListRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetKeywordDownList(ctx, Req)
}

func (p *kGoodsServiceClient) SearchSpuList(ctx context.Context, Req *goods.SearchSpuListReq, callOptions ...callopt.Option) (r *goods.SearchSpuListRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SearchSpuList(ctx, Req)
}

func (p *kGoodsServiceClient) GetGoodsDetail(ctx context.Context, Req *goods.GetGoodsDetailReq, callOptions ...callopt.Option) (r *goods.GetGoodsDetailRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetGoodsDetail(ctx, Req)
}

func (p *kGoodsServiceClient) GetSellerGoodsList(ctx context.Context, Req *goods.GetSellerGoodsListReq, callOptions ...callopt.Option) (r *goods.GetSellerGoodsListRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetSellerGoodsList(ctx, Req)
}

func (p *kGoodsServiceClient) GetGoodsList(ctx context.Context, Req *goods.GetGoodsListReq, callOptions ...callopt.Option) (r *goods.GetGoodsListRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetGoodsList(ctx, Req)
}

func (p *kGoodsServiceClient) GetCategoryList(ctx context.Context, Req *goods.GetCategoryListReq, callOptions ...callopt.Option) (r *goods.GetCategoryListRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetCategoryList(ctx, Req)
}

func (p *kGoodsServiceClient) GetBrandList(ctx context.Context, Req *goods.GetBrandListReq, callOptions ...callopt.Option) (r *goods.GetBrandListRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetBrandList(ctx, Req)
}

func (p *kGoodsServiceClient) AddToCart(ctx context.Context, Req *goods.AddToCartReq, callOptions ...callopt.Option) (r *goods.AddToCartRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddToCart(ctx, Req)
}

func (p *kGoodsServiceClient) PublishGoods(ctx context.Context, Req *goods.PublishGoodsReq, callOptions ...callopt.Option) (r *goods.PublishGoodsRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishGoods(ctx, Req)
}

func (p *kGoodsServiceClient) StockReduce(ctx context.Context, Req *goods.StockReduceReq, callOptions ...callopt.Option) (r *goods.StockReduceRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.StockReduce(ctx, Req)
}

func (p *kGoodsServiceClient) StockIncrease(ctx context.Context, Req *goods.StockIncreaseReq, callOptions ...callopt.Option) (r *goods.StockIncreaseRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.StockIncrease(ctx, Req)
}
