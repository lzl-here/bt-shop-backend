package handler

import (
	"context"

	"github.com/lzl-here/bt-shop-backend/apps/order/internal/constant"
	"github.com/lzl-here/bt-shop-backend/apps/order/internal/domain/model"
	ogen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/order"
	"gorm.io/gorm"
)

// 支付成功
// 流转状态
func (h *OrderHandler) PaySuccessToOrder(ctx context.Context, req *ogen.PaySuccessToOrderReq) (*ogen.PaySuccessToOrderRsp, error) {
	h.rep.GetDB().Transaction(func(tx *gorm.DB) error {
		{
			where := &model.Trade{TradeNo: req.TradeNo}
			update := &model.Trade{TradeStatus: constant.TradeStatusPaid}
			err := h.rep.UpdateTrade(ctx, h.rep.GetDB(), where, update)
			if err != nil {
				return err
			}
		}
		{
			where := &model.Order{TradeNo: req.TradeNo}
			update := &model.Order{OrderStatus: constant.OrderStatusPaid}
			err := h.rep.UpdateOrder(ctx, h.rep.GetDB(), where, update)
			if err != nil {
				return err
			}
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
		{
			where := &model.Trade{TradeNo: req.TradeNo}
			update := &model.Trade{TradeStatus: constant.TradeStatusCancel}
			err := h.rep.UpdateTrade(ctx, h.rep.GetDB(), where, update)
			if err != nil {
				return err
			}
		}
		{
			where := &model.Order{TradeNo: req.TradeNo}
			update := &model.Order{OrderStatus: constant.OrderStatusCanceled}
			err := h.rep.UpdateOrder(ctx, h.rep.GetDB(), where, update)
			if err != nil {
				return err
			}
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
/**	
* @description: 获取交易列表
*/
func (h *OrderHandler) GetTradeList(ctx context.Context, req *ogen.GetTradeListReq) (*ogen.GetTradeListRsp, error) {
	return nil, nil
}
