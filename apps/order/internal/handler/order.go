package handler

import (
	"context"

	"github.com/lzl-here/bt-shop-backend/apps/order/internal/domain/model"
	ogen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/order"
)

// TODO 订单详情
/*
* @description: 获取订单详情
**/
func (h *OrderHandler) GetOrderDetail(ctx context.Context, req *ogen.GetOrderDetailReq) (*ogen.GetOrderDetailRsp, error) {
	return nil, nil
}

/*
* @description: 获取卖家订单列表
**/
func (h *OrderHandler) GetSellerOrderList(ctx context.Context, req *ogen.GetSellerOrderListReq) (*ogen.GetSellerOrderListRsp, error) {
	// 获取卖家订单列表
	orders, err := h.rep.PageOrderList(ctx, &model.Order{
		ShopID:      req.ShopId,
		OrderStatus: req.OrderStatus,
	}, int(req.PageSize), int(req.PageNo))
	if err != nil {
		return nil, err
	}
	total, err := h.rep.CountOrder(ctx, &model.Order{
		ShopID:      req.ShopId,
		OrderStatus: req.OrderStatus,
	})
	if err != nil {
		return nil, err
	}
	// 将订单列表转换为响应
	return buildSellerOrderListRsp(orders, int(total), int(req.PageNo), int(req.PageSize))
}

func buildSellerOrderListRsp(orders []*model.Order, total, pageNo, pageSize int) (*ogen.GetSellerOrderListRsp, error) {
	res := &ogen.GetSellerOrderListRsp{}
	res.Code = 1
	res.Msg = "success"
	res.Data = &ogen.GetSellerOrderListRsp_GetSellerOrderListRspData{
		OrderList: []*ogen.BaseOrder{},
		Total:     int32(total),
		PageNo:    int32(pageNo),
		PageSize:  int32(pageSize),
	}
	for _, order := range orders {
		orderRsp, err := order.ToGen()
		if err != nil {
			return nil, err
		}
		res.Data.OrderList = append(res.Data.OrderList, orderRsp)
	}
	return res, nil
}
