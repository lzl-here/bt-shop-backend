package handler

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/lzl-here/bt-shop-backend/apps/pay/internal/constant"
	"github.com/lzl-here/bt-shop-backend/apps/pay/internal/domain/model"
	"github.com/lzl-here/bt-shop-backend/apps/pay/internal/repo"
	pgen "github.com/lzl-here/bt-shop-backend/kitex_gen/pay"
	bizerr "github.com/lzl-here/bt-shop-backend/pkg/err"
	"github.com/smartwalle/alipay/v3"
)

type PayHandler struct {
	rep repo.RepoInterface
}

func NewPayHandler(rep repo.RepoInterface) *PayHandler {
	return &PayHandler{
		rep: rep,
	}
}

// 统一支付接口:
// 调用第三方支付，创建本地支付流水
// TODO 根据tradeNo做幂等
func (h *PayHandler) Pay(ctx context.Context, req *pgen.PayReq) (res *pgen.PayRsp, err error) {
	param := alipay.TradePagePay{}
	param.OutTradeNo = req.TradeNo
	param.Subject = req.Subject
	param.TotalAmount = req.TotalAmount
	url, err := h.rep.AlipayPay(param)
	if err != nil {
		return nil, err
	}
	if _, err = h.rep.CreatePayFlow(&model.PayFlow{
		TradeNo:     req.TradeNo,
		PayState:    constant.PayStatePaying,
		TotalAmount: req.TotalAmount,
	}); err != nil {
		return nil, err
	}
	return &pgen.PayRsp{
		Code: 1,
		Msg:  "ok",
		Data: &pgen.PayRsp_PayRspData{
			PayUrl:  url.String(),
			TradeNo: req.TradeNo,
		},
	}, nil
}

// 支付宝回调
// 本地支付流水状态流转 和 订单状态流转
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
	return nil, nil

}

// 支付宝支付成功回调: 修改本地状态，通知order同步订单状态
func (h *PayHandler) paySuccess(ctx context.Context, req *pgen.AlipayWebhookReq) (*pgen.AlipayWebhookRsp, error) {
	// TODO 状态机
	p := &model.PayFlow{
		TradeNo:       req.OutTradeNo,
		ThirdTradeNo:  req.TradeNo,
		PayState:      constant.PayStatePayed,
		ThirdBuyerID:  req.BuyerId,
		ThirdSellerID: req.SellerId,
	}
	if _, err := h.rep.UpdatePayFlow(p); err != nil {
		return nil, err
	}

	// TODO 通知order同步订单状态
	return &pgen.AlipayWebhookRsp{
		Code: 1,
		Msg:  "ok",
	}, nil
}

func (h *PayHandler) payPartRefund(ctx context.Context, req *pgen.AlipayWebhookReq) (*pgen.AlipayWebhookRsp, error) {
	// TODO 状态机
	p := &model.PayFlow{
		TradeNo:      req.OutTradeNo,
		ThirdTradeNo: req.TradeNo,
		PayState:     constant.PayStatePartRefunded,
	}
	if _, err := h.rep.UpdatePayFlow(p); err != nil {
		return nil, err
	}

	// TODO 通知order同步订单状态
	return &pgen.AlipayWebhookRsp{
		Code: 1,
		Msg:  "ok",
	}, nil
}

func (h *PayHandler) payFullRefund(ctx context.Context, req *pgen.AlipayWebhookReq) (*pgen.AlipayWebhookRsp, error) {
	// TODO 状态机
	p := &model.PayFlow{
		TradeNo:      req.OutTradeNo,
		ThirdTradeNo: req.TradeNo,
		PayState:     constant.PayStateFullRefunded,
	}
	if _, err := h.rep.UpdatePayFlow(p); err != nil {
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
		return nil, kerrors.NewBizStatusError(bizerr.ErrDownStreamError.BizStatusCode(), alipayResp.Msg)
	}
	// TODO 状态机
	p := &model.PayFlow{
		TradeNo:       req.TradeNo,
		ThirdTradeNo:  alipayResp.TradeNo,
		PayState:      constant.PayStateCanceled,
		ThirdBuyerID:  req.BuyerId,
		ThirdSellerID: req.SellerId,
	}
	if _, err := h.rep.UpdatePayFlow(p); err != nil {
		return nil, err
	}

	// TODO 通知order同步订单状态
	return &pgen.AlipayWebhookRsp{
		Code: 1,
		Msg:  "ok",
	}, nil
}

// 取消支付，本地流转为退款中，通知order为退款中，
// 假如退款时间过长，用户进入到订单页面看到的状态为已完成，但是应该为退款中
func (h *PayHandler) ClosePay(ctx context.Context, req *pgen.ClosePayReq) (*pgen.ClosePayRsp, error) {
	param := alipay.TradeClose{
		OutTradeNo: req.TradeNo,
	}
	alipayResp, err := h.rep.AlipayClose(ctx, param)
	if err != nil {
		return nil, err
	}
	if !alipayResp.IsSuccess() {
		return nil, kerrors.NewBizStatusError(bizerr.ErrDownStreamError.BizStatusCode(), alipayResp.Msg)
	}
	_, err = h.rep.UpdatePayFlow(&model.PayFlow{
		TradeNo:  req.TradeNo,
		PayState: constant.PayStateRefunding,
	})
	if err != nil {
		return nil, err
	}

	// TODO 通知order为退款中
	return &pgen.ClosePayRsp{
		Code: 1,
		Msg:  "ok",
	}, nil
}

// 退款，本地支付流水状态流转 和 订单状态流转
func (h *PayHandler) RefundPay(ctx context.Context, req *pgen.RefundPayReq) (res *pgen.RefundPayRsp, err error) {
	param := alipay.TradeRefund{
		OutTradeNo:   req.TradeNo,
		RefundAmount: req.RefundAmount,
		RefundReason: req.RefundReason,
		OutRequestNo: req.OutRequestNo,
	}
	alipayResp, err := h.rep.AlipayRefund(ctx, param)
	if err != nil {
		return nil, err
	}
	if !alipayResp.IsSuccess() {
		return nil, kerrors.NewBizStatusError(bizerr.ErrDownStreamError.BizStatusCode(), alipayResp.Msg)
	}

	// TODO 状态机 和 部分退款
	p := &model.PayFlow{
		TradeNo:      req.TradeNo,
		ThirdTradeNo: alipayResp.TradeNo,
		PayState:     constant.PayStateFullRefunded,
	}
	if _, err := h.rep.UpdatePayFlow(p); err != nil {
		return nil, err
	}
	// TODO 同步order
	return &pgen.RefundPayRsp{
		Code: 1,
		Msg:  "ok",
	}, nil
}
