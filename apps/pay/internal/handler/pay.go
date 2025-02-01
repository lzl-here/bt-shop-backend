package handler

import (
	"context"
	"time"

	pgen "github.com/lzl-here/bt-shop-backend/kitex_gen/pay"
	bizerr "github.com/lzl-here/bt-shop-backend/pkg/err"
	"github.com/lzl-here/bt-shop-backend/pkg/utils"
	"github.com/smartwalle/alipay/v3"
)

// 统一支付接口:
// 请求第三方支付机构创建一笔交易，并且返回支付页面给前端
func (h *PayHandler) Pay(ctx context.Context, req *pgen.PayReq) (res *pgen.PayRsp, err error) {
	// 根据tradeNo做幂等
	ok, cleaner, _, _ := utils.NoDuplicate(ctx, h.rep.GetCache(), "pay", req.TradeNo, 30*time.Second)
	if !ok {
		return nil, bizerr.ErrDuplicateReq
	}

	param := alipay.TradePagePay{}
	param.OutTradeNo = req.TradeNo
	param.Subject = req.Subject
	param.TotalAmount = req.TotalAmount
	url, err := h.rep.AlipayPay(param)
	if err != nil {
		cleaner(5)
		return nil, err
	}
	return &pgen.PayRsp{
		Code: 1,
		Msg:  "success",
		Data: &pgen.PayRsp_PayRspData{
			PayPageUrl: url.String(),
			TradeNo:    req.TradeNo,
		},
	}, nil
}
