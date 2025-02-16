package api

import (
	"context"

	"github.com/lzl-here/bt-shop-backend/apps/order/internal/handler"
	"github.com/lzl-here/bt-shop-backend/apps/order/internal/repo"
	ogen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/order"
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

// 支付取消回调
func (s *OrderServer) PayCancelToOrder(ctx context.Context, req *ogen.PayCancelToOrderReq) (res *ogen.PayCancelToOrderRsp, err error) {
	return handler.NewOrderHandler(s.rep).PayCancelToOrder(ctx, req)
}

// 获取交易列表
func (s *OrderServer) GetTradeList(ctx context.Context, req *ogen.GetTradeListReq) (res *ogen.GetTradeListRsp, err error) {
	return handler.NewOrderHandler(s.rep).GetTradeList(ctx, req)
}

// 交易详情
func (s *OrderServer) GetTradeDetail(ctx context.Context, req *ogen.GetTradeDetailReq) (res *ogen.GetTradeDetailRsp, err error) {
	return handler.NewOrderHandler(s.rep).GetTradeDetail(ctx, req)
}

// 订单详情
func (s *OrderServer) GetOrderDetail(ctx context.Context, req *ogen.GetOrderDetailReq) (res *ogen.GetOrderDetailRsp, err error) {
	return handler.NewOrderHandler(s.rep).GetOrderDetail(ctx, req)
}

// 获取卖家订单列表
func (s *OrderServer) GetSellerOrderList(ctx context.Context, req *ogen.GetSellerOrderListReq) (res *ogen.GetSellerOrderListRsp, err error) {
	return handler.NewOrderHandler(s.rep).GetSellerOrderList(ctx, req)
}
