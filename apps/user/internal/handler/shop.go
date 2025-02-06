package handler

import (
	"context"

	"github.com/lzl-here/bt-shop-backend/apps/user/internal/domain/model"
	ugen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/user"
)

func (h *UserHandler) GetShopDetail(ctx context.Context, req *ugen.GetShopDetailReq) (*ugen.GetShopDetailRsp, error) {
	shop, err := h.rep.GetShopDetail(ctx, &model.Shop{
		ID: req.ShopId,
	})

	if err != nil {
		return nil, err
	}

	baseShop, err := shop.ToGen()
	if err != nil {
		return nil, err
	}
	return &ugen.GetShopDetailRsp{
		Code: 1,
		Msg:  "success",
		Data: &ugen.GetShopDetailRsp_Data{
			Shop: baseShop,
		},
	}, nil
}
