package handler

import (
	"context"

	ggen "github.com/lzl-here/bt-shop-backend/kitex_gen/goods"
)

func (h *GoodsHandler) AddToCart(ctx context.Context, req *ggen.AddToCartReq) (*ggen.AddToCartRsp, error) {
	return nil, nil
}
