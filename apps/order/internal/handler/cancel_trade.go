package handler

import (
	"context"

	ogen "github.com/lzl-here/bt-shop-backend/kitex_gen/order"
)

func (h *OrderHandler) CancelTrade(ctx context.Context, req *ogen.CancelTradeReq) (res *ogen.CancelTradeRsp, err error) {
	return nil, nil
}
