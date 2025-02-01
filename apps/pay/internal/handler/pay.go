package handler

import (
	"context"

	pgen "github.com/lzl-here/bt-shop-backend/kitex_gen/pay"
	"github.com/smartwalle/alipay/v3"
)

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
	return &pgen.PayRsp{
		Code: 1,
		Msg:  "ok",
		Data: &pgen.PayRsp_PayRspData{
			PayUrl:  url.String(),
			TradeNo: req.TradeNo,
		},
	}, nil
}
