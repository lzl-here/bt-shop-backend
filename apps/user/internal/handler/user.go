package handler

import (
	"context"

	"github.com/lzl-here/bt-shop-backend/apps/user/internal/domain/model"
	bizerr "github.com/lzl-here/bt-shop-backend/pkg/err"
	ugen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/user"
)

func (h *UserHandler) GetUserInfo(ctx context.Context, req *ugen.GetUserInfoReq) (*ugen.GetUserInfoRsp, error) {
	user, err := h.rep.GetUser(ctx, &model.User{ID: req.UserId})
	if err != nil {
		return nil, err
	}
	baseUser, err := user.ToGen()
	if err != nil {
		return nil, err
	}
	return &ugen.GetUserInfoRsp{
		Code: 1,
		Msg:  "success",
		Data: &ugen.GetUserInfoRsp_Data{
			User: baseUser,
		},
	}, nil
}

func (h *UserHandler) UpdateUserInfo(ctx context.Context, req *ugen.UpdateUserInfoReq) (*ugen.UpdateUserInfoRsp, error) {
	user, err := h.rep.GetUser(ctx, &model.User{ID: req.UserId})
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, bizerr.ErrResourceNotFound
	}
	user.AvatarUrl = req.AvatarUrl
	user.ShopID = req.ShopId
	user, err = h.rep.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	baseUser, err := user.ToGen()
	if err != nil {
		return nil, err
	}
	return &ugen.UpdateUserInfoRsp{
		Code: 1,
		Msg:  "success",
		Data: &ugen.UpdateUserInfoRsp_Data{
			User: baseUser,
		},
	}, nil
}
