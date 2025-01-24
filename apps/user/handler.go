package main

import (
	"context"
	user "github.com/lzl-here/bt-shop-backend/kitex_gen/user"
)

type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {

	resp = new(user.RegisterResp)

	return
}

func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp = new(user.LoginResp)

	return
}
