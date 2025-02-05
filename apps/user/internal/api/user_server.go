package api

import (
	"context"

	"github.com/lzl-here/bt-shop-backend/apps/user/internal/handler"
	"github.com/lzl-here/bt-shop-backend/apps/user/internal/repo"
	ugen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/user"
)

// rpc入口

var _ (ugen.UserService) = (*UserServer)(nil)

type UserServer struct {
	rep *repo.Repo
}

func NewUserServer(rep *repo.Repo) *UserServer {
	return &UserServer{
		rep: rep,
	}
}

// 普通登录
func (s *UserServer) NormalLogin(ctx context.Context, req *ugen.NormalLoginReq) (*ugen.NormalLoginRsp, error) {
	return handler.NewUserHandler(s.rep).NormalLogin(ctx, req)
}

// 普通注册
func (s *UserServer) NormalRegister(ctx context.Context, req *ugen.NormalRegisterReq) (*ugen.NormalRegisterRsp, error) {
	return handler.NewUserHandler(s.rep).NormalRegister(ctx, req)
}

// 登出
func (s *UserServer) Logout(ctx context.Context, req *ugen.LogoutReq) (*ugen.LogoutRsp, error) {
	return handler.NewUserHandler(s.rep).Logout(ctx, req)
}

// 获取用户信息
func (s *UserServer) GetUserInfo(ctx context.Context, req *ugen.GetUserInfoReq) (*ugen.GetUserInfoRsp, error) {
	return handler.NewUserHandler(s.rep).GetUserInfo(ctx, req)
}

// 更新用户信息
func (s *UserServer) UpdateUserInfo(ctx context.Context, req *ugen.UpdateUserInfoReq) (*ugen.UpdateUserInfoRsp, error) {
	return handler.NewUserHandler(s.rep).UpdateUserInfo(ctx, req)
}
