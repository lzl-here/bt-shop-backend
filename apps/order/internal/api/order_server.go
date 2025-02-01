package api

import (
	"context"

	"github.com/lzl-here/bt-shop-backend/apps/order/internal/handler"
	"github.com/lzl-here/bt-shop-backend/apps/order/internal/repo"
	ogen "github.com/lzl-here/bt-shop-backend/kitex_gen/order"
)

// rpc入口

var _ (ogen.OrderService) = (*OrderServer)(nil)

type OrderServer struct {
	rep *repo.Repo
}

func NewOrderServer(rep *repo.Repo) *OrderServer {
	return &OrderServer{
		rep: rep,
	}
}

// 创建一笔交易
func (s *OrderServer) CreateTrade(ctx context.Context, req *ogen.CreateTradeReq) (res *ogen.CreateTradeRsp, err error) {
	return handler.NewOrderHandler(s.rep).CreateTrade(ctx, req)
}

// 取消一笔交易
func (s *OrderServer) CancelTrade(ctx context.Context, req *ogen.CancelTradeReq) (res *ogen.CancelTradeRsp, err error) {
	return handler.NewOrderHandler(s.rep).CancelTrade(ctx, req)
}

// 获取订单项列表
func (s *OrderServer) GetOrderItems(ctx context.Context, req *ogen.GetOrderItemsReq) (res *ogen.GetOrderItemsRsp, err error) {
	return handler.NewOrderHandler(s.rep).GetOrderItems(ctx, req)
}

// 支付成功回调
func (s *OrderServer) PaySuccessToOrder(ctx context.Context, req *ogen.PaySuccessToOrderReq) (res *ogen.PaySuccessToOrderRsp, err error) {
	return handler.NewOrderHandler(s.rep).PaySuccessToOrder(ctx, req)
}
