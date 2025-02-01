package handler

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	pgen "github.com/lzl-here/bt-shop-backend/kitex_gen/pay"
	bizerr "github.com/lzl-here/bt-shop-backend/pkg/err"
	"github.com/smartwalle/alipay/v3"
)

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
		return nil, kerrors.NewBizStatusError(bizerr.ErrDownStream.BizStatusCode(), alipayResp.Msg)
	}

	// // TODO 状态机 和 部分退款
	// p := &model.PayFlow{
	// 	TradeNo:      req.TradeNo,
	// 	ThirdTradeNo: alipayResp.TradeNo,
	// 	Status:     constant.PayStatusRefund,
	// }
	// if _, err := h.rep.CreatePayFlows(p); err != nil {
	// 	return nil, err
	// }
	// TODO 同步order
	return &pgen.RefundPayRsp{
		Code: 1,
		Msg:  "ok",
	}, nil
}
