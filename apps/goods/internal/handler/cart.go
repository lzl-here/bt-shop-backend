package handler

import (
	"context"

	ggen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/goods"
)

// TODO 添加到购物车
func (h *GoodsHandler) AddToCart(ctx context.Context, req *ggen.AddToCartReq) (*ggen.AddToCartRsp, error) {
	return nil, nil
}
