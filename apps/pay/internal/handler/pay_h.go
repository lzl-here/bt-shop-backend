package handler

import (
	"context"

	"github.com/lzl-here/bt-shop-backend/apps/pay/internal/repo"
	pgen "github.com/lzl-here/bt-shop-backend/kitex_gen/pay"
	"github.com/smartwalle/alipay/v3"
)

type PayHandler struct {
	rep *repo.Repo
}

func NewPayHandler(rep *repo.Repo) *PayHandler {
	return &PayHandler{
		rep: rep,
	}
}

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
		Code: 0,
		Msg:  "ok",
		Data: &pgen.PayRsp_PayRspData{
			PayUrl:     url.String(),
			TradeNo:    req.TradeNo,
		},
	}, nil
}