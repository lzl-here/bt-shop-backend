package handler

import (
	"context"

	"github.com/lzl-here/bt-shop-backend/apps/pay/internal/constant"
	"github.com/lzl-here/bt-shop-backend/apps/pay/internal/domain/model"
	"github.com/lzl-here/bt-shop-backend/apps/pay/internal/repo"
	pgen "github.com/lzl-here/bt-shop-backend/kitex_gen/pay"
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
		OutTradeNo:  req.TradeNo,
		PayState:    constant.PayStatePaying,
		TotalAmount: req.TotalAmount,
	}); err != nil {
		return nil, err
	}
	return &pgen.PayRsp{
		Code: 0,
		Msg:  "ok",
		Data: &pgen.PayRsp_PayRspData{
			PayUrl:  url.String(),
			TradeNo: req.TradeNo,
		},
	}, nil
}

// TODO implement

// 支付宝回调
// 本地支付流水状态流转 和 订单状态流转
func (h *PayHandler) AlipayWebhook(ctx context.Context, req *pgen.AlipayWebhookReq) (res *pgen.AlipayWebhookRsp, err error) {
	panic("implement me")
}

// 取消支付，本地支付流水状态流转 和 订单状态流转
func (h *PayHandler) CancelPay(ctx context.Context, req *pgen.CancelPayReq) (res *pgen.CancelPayRsp, err error) {
	panic("implement me")
}

// 退款，本地支付流水状态流转 和 订单状态流转
func (h *PayHandler) RefundPay(ctx context.Context, req *pgen.RefundPayReq) (res *pgen.RefundPayRsp, err error) {
	panic("implement me")
}
