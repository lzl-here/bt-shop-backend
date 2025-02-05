package handler

import (
	"context"

	ogen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/order"
)

// TODO 订单详情
/*
* @description: 获取订单详情
**/
func (h *OrderHandler) GetOrderDetail(ctx context.Context, req *ogen.GetOrderDetailReq) (*ogen.GetOrderDetailRsp, error) {
	return nil, nil
}
