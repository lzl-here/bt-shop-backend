package handler

import (
	"context"
	"time"

	"github.com/lzl-here/bt-shop-backend/apps/order/internal/repo"
	"github.com/lzl-here/bt-shop-backend/pkg/utils"

	ogen "github.com/lzl-here/bt-shop-backend/kitex_gen/order"
	bizerr "github.com/lzl-here/bt-shop-backend/pkg/err"
)

type OrderHandler struct {
	rep repo.RepoInterface
}

func NewOrderHandler(rep repo.RepoInterface) *OrderHandler {
	return &OrderHandler{
		rep: rep,
	}
}

func (h *OrderHandler) CreateTrade(ctx context.Context, req *ogen.CreateTradeReq) (res *ogen.CreateTradeRsp, err error) {
	// 幂等处理
	ok, cleaner, _, err := utils.NoDuplicate(ctx, h.rep.GetCache(), "order", req.Token, 30*time.Second)
	if !ok {
		return nil, bizerr.ErrDuplicateReq
	}
	if err != nil {
		cleaner(5)
		return nil, err
	}

	// TODO 校验商品信息
	// 创建交易、订单、订单项
	
	return nil, nil
}

func (h *OrderHandler) CancelTrade(ctx context.Context, req *ogen.CancelTradeReq) (res *ogen.CancelTradeRsp, err error) {
	return nil, nil
}
