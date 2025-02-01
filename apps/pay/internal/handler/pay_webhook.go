package handler

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/lzl-here/bt-shop-backend/apps/pay/internal/constant"
	"github.com/lzl-here/bt-shop-backend/apps/pay/internal/domain/model"
	ogen "github.com/lzl-here/bt-shop-backend/kitex_gen/order"
	pgen "github.com/lzl-here/bt-shop-backend/kitex_gen/pay"
	bizerr "github.com/lzl-here/bt-shop-backend/pkg/err"
	"github.com/lzl-here/bt-shop-backend/pkg/utils"
	"github.com/smartwalle/alipay/v3"
)

// 支付宝回调
// 本地支付流水状态流转 和 订单状态流转
// TODO 状态机
func (h *PayHandler) AlipayWebhook(ctx context.Context, req *pgen.AlipayWebhookReq) (res *pgen.AlipayWebhookRsp, err error) {
	if req.NotifyType == constant.AlipayNoticeSyncTrade {
		if req.TradeStatus == constant.AlipayTradeSuccess && req.GmtRefund == "" {
			// 支付成功
			return h.paySuccess(ctx, req)
		} else if req.TradeStatus == constant.AlipayTradeSuccess && req.GmtRefund != "" {
			// 部分退款
			return h.payPartRefund(ctx, req)
		} else if req.TradeStatus == constant.AlipayTradeClosed && req.GmtRefund == "" {
			// 支付关闭
			return h.payClosed(ctx, req)
		} else if req.TradeStatus == constant.AlipayTradeClosed && req.GmtRefund != "" {
			// 全额退款
			return h.payFullRefund(ctx, req)
		}
	}
	// TODO 回调的返回值应该是什么？
	return nil, nil

}

// 支付宝支付成功回调
// 1. 生成支付流水
// 2. 通知order同步订单状态
func (h *PayHandler) paySuccess(ctx context.Context, req *pgen.AlipayWebhookReq) (*pgen.AlipayWebhookRsp, error) {

	// 1. 根据订单项生成流水, 一个订单项对应一笔支付流水
	getItemsRsp, err := h.rep.GetOrderItems(ctx, &ogen.GetOrderItemsReq{
		TradeNo: req.OutTradeNo,
	})
	if err != nil {
		return nil, err
	}
	if getItemsRsp.Code != 1 {
		return nil, bizerr.ErrDownStream
	}
	payFlows := make([]*model.PayFlow, 0)
	for _, i := range getItemsRsp.Data.OrderItemList {
		p := &model.PayFlow{
			PayNo:         utils.GenPayNo(),
			TradeNo:       i.TradeNo,
			OrderNo:       i.OrderNo,
			ThirdTradeNo:  req.TradeNo,
			Status:        constant.PayStatusSuccess,
			OrderItemID:   i.Id,
			InOutType:     constant.PayInType,
			ThirdBuyerID:  req.BuyerId,
			ThirdSellerID: req.SellerId,
			TotalAmount:   i.SkuAmount,
			ShopID:        i.ShopId,
			SellerID:      i.SellerId,
			BuyerID:       i.BuyerId,
		}
		payFlows = append(payFlows, p)
	}
	if err := h.rep.CreatePayFlows(payFlows); err != nil {
		return nil, err
	}

	// 2. 通知order同步订单状态
	payRsp, err := h.rep.PaySuccessToOrder(ctx, &ogen.PaySuccessToOrderReq{
		TradeNo: req.OutTradeNo,
	})
	if err != nil {
		return nil, err
	}
	if payRsp.Code != 1 {
		return nil, bizerr.ErrDownStream
	}
	return &pgen.AlipayWebhookRsp{
		Code: 1,
		Msg:  "ok",
	}, nil
}

func (h *PayHandler) payPartRefund(ctx context.Context, req *pgen.AlipayWebhookReq) (*pgen.AlipayWebhookRsp, error) {
	p := &model.PayFlow{
		TradeNo:      req.OutTradeNo,
		ThirdTradeNo: req.TradeNo,
	}
	if _, err := h.rep.UpdatePayFlow(&model.PayFlow{TradeNo: req.OutTradeNo}, p); err != nil {
		return nil, err
	}

	// TODO 通知order同步订单状态
	return &pgen.AlipayWebhookRsp{
		Code: 1,
		Msg:  "ok",
	}, nil
}

func (h *PayHandler) payFullRefund(ctx context.Context, req *pgen.AlipayWebhookReq) (*pgen.AlipayWebhookRsp, error) {
	p := &model.PayFlow{
		TradeNo:      req.OutTradeNo,
		ThirdTradeNo: req.TradeNo,
	}
	if _, err := h.rep.UpdatePayFlow(&model.PayFlow{TradeNo: req.OutTradeNo}, p); err != nil {
		return nil, err
	}

	// TODO 通知order同步订单状态
	return &pgen.AlipayWebhookRsp{
		Code: 1,
		Msg:  "ok",
	}, nil
}

// 支付宝支付关闭回调：关闭订单
func (h *PayHandler) payClosed(ctx context.Context, req *pgen.AlipayWebhookReq) (res *pgen.AlipayWebhookRsp, err error) {
	param := alipay.TradeClose{
		OutTradeNo: req.TradeNo,
	}
	alipayResp, err := h.rep.AlipayClose(ctx, param)
	if err != nil {
		return nil, err
	}
	if !alipayResp.IsSuccess() {
		return nil, kerrors.NewBizStatusError(bizerr.ErrDownStream.BizStatusCode(), alipayResp.Msg)
	}
	p := &model.PayFlow{
		TradeNo:       req.OutTradeNo,
		ThirdTradeNo:  req.TradeNo,
		ThirdBuyerID:  req.BuyerId,
		ThirdSellerID: req.SellerId,
		Status:        constant.PayStatusCancel,
		TotalAmount:   req.TotalAmount,
	}
	if _, err := h.rep.UpdatePayFlow(&model.PayFlow{TradeNo: req.OutTradeNo}, p); err != nil {
		return nil, err
	}

	// TODO 通知order同步订单状态
	return &pgen.AlipayWebhookRsp{
		Code: 1,
		Msg:  "ok",
	}, nil
}
