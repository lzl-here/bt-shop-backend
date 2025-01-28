package handler

import (
	"context"

	"github.com/lzl-here/bt-shop-backend/apps/order/internal/repo"
	ogen "github.com/lzl-here/bt-shop-backend/kitex_gen/order"
)

type OrderHandler struct {
	rep repo.RepoInterface
}

func NewOrderHandler(rep repo.RepoInterface) *OrderHandler {
	return &OrderHandler{
		rep: rep,
	}
}

func (h *OrderHandler) GenSettlePage(ctx context.Context, req *ogen.GenSettlePageReq) (res *ogen.GenSettlePageRsp, err error) {
	return nil, nil
}

func (h *OrderHandler) CreateTrade(ctx context.Context, req *ogen.CreateTradeReq) (res *ogen.CreateTradeRsp, err error) {
	return nil, nil
}

func (h *OrderHandler) CancelTrade(ctx context.Context, req *ogen.CancelTradeReq) (res *ogen.CancelTradeRsp, err error) {
	return nil, nil
}
