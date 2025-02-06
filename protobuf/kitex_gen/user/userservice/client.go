// Code generated by Kitex v0.12.1. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	user "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	NormalLogin(ctx context.Context, Req *user.NormalLoginReq, callOptions ...callopt.Option) (r *user.NormalLoginRsp, err error)
	NormalRegister(ctx context.Context, Req *user.NormalRegisterReq, callOptions ...callopt.Option) (r *user.NormalRegisterRsp, err error)
	Logout(ctx context.Context, Req *user.LogoutReq, callOptions ...callopt.Option) (r *user.LogoutRsp, err error)
	GetUserInfo(ctx context.Context, Req *user.GetUserInfoReq, callOptions ...callopt.Option) (r *user.GetUserInfoRsp, err error)
	UpdateUserInfo(ctx context.Context, Req *user.UpdateUserInfoReq, callOptions ...callopt.Option) (r *user.UpdateUserInfoRsp, err error)
	GetShopDetail(ctx context.Context, Req *user.GetShopDetailReq, callOptions ...callopt.Option) (r *user.GetShopDetailRsp, err error)
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
	return &kUserServiceClient{
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

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) NormalLogin(ctx context.Context, Req *user.NormalLoginReq, callOptions ...callopt.Option) (r *user.NormalLoginRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.NormalLogin(ctx, Req)
}

func (p *kUserServiceClient) NormalRegister(ctx context.Context, Req *user.NormalRegisterReq, callOptions ...callopt.Option) (r *user.NormalRegisterRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.NormalRegister(ctx, Req)
}

func (p *kUserServiceClient) Logout(ctx context.Context, Req *user.LogoutReq, callOptions ...callopt.Option) (r *user.LogoutRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Logout(ctx, Req)
}

func (p *kUserServiceClient) GetUserInfo(ctx context.Context, Req *user.GetUserInfoReq, callOptions ...callopt.Option) (r *user.GetUserInfoRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, Req)
}

func (p *kUserServiceClient) UpdateUserInfo(ctx context.Context, Req *user.UpdateUserInfoReq, callOptions ...callopt.Option) (r *user.UpdateUserInfoRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUserInfo(ctx, Req)
}

func (p *kUserServiceClient) GetShopDetail(ctx context.Context, Req *user.GetShopDetailReq, callOptions ...callopt.Option) (r *user.GetShopDetailRsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetShopDetail(ctx, Req)
}
