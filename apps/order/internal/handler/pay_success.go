package handler

import (
	"context"

	"github.com/lzl-here/bt-shop-backend/apps/order/internal/constant"
	"github.com/lzl-here/bt-shop-backend/apps/order/internal/domain/model"
	ogen "github.com/lzl-here/bt-shop-backend/kitex_gen/order"
	"gorm.io/gorm"
)

// 支付成功
// 流转状态
func (h *OrderHandler) PaySuccessToOrder(ctx context.Context, req *ogen.PaySuccessToOrderReq) (*ogen.PaySuccessToOrderRsp, error) {
	h.rep.GetDB().Transaction(func(tx *gorm.DB) error {
		err := h.rep.UpdateTrade(ctx, &model.Trade{TradeNo: req.TradeNo}, &model.Trade{TradeStatus: constant.TradeStatusPaid})
		if err != nil {
			return err
		}
		err = h.rep.UpdateOrder(ctx, &model.Order{TradeNo: req.TradeNo}, &model.Order{OrderStatus: constant.OrderStatusPaid})
		if err != nil {
			return err
		}
		return nil
	})
	return &ogen.PaySuccessToOrderRsp{
		Code: 1,
		Msg:  "success",
	}, nil
}

// 支付取消
// 流转状态
func (h *OrderHandler) PayCancelToOrder(ctx context.Context, req *ogen.PayCancelToOrderReq) (*ogen.PayCancelToOrderRsp, error) {
	h.rep.GetDB().Transaction(func(tx *gorm.DB) error {
		err := h.rep.UpdateTrade(ctx, &model.Trade{TradeNo: req.TradeNo}, &model.Trade{TradeStatus: constant.TradeStatusCancel})
		if err != nil {
			return err
		}
		err = h.rep.UpdateOrder(ctx, &model.Order{TradeNo: req.TradeNo}, &model.Order{OrderStatus: constant.OrderStatusCanceled})
		if err != nil {
			return err
		}
		return nil
	})
	return &ogen.PayCancelToOrderRsp{
		Code: 1,
		Msg:  "success",
		Data: &ogen.PayCancelToOrderRsp_PayCancelToOrderRspData{
			TradeNo: req.TradeNo,
		},
	}, nil
}
