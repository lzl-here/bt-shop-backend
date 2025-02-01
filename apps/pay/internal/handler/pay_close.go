package handler

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/lzl-here/bt-shop-backend/apps/pay/internal/constant"
	"github.com/lzl-here/bt-shop-backend/apps/pay/internal/domain/model"
	pgen "github.com/lzl-here/bt-shop-backend/kitex_gen/pay"
	bizerr "github.com/lzl-here/bt-shop-backend/pkg/err"

	"github.com/smartwalle/alipay/v3"
)

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
		return nil, kerrors.NewBizStatusError(bizerr.ErrDownStream.BizStatusCode(), alipayResp.Msg)
	}
	// TODO 给每个订单项都生成
	_, err = h.rep.CreatePayFlow(&model.PayFlow{
		TradeNo:      req.TradeNo,
		Status:       constant.PayStatusCancel,
		ThirdTradeNo: alipayResp.TradeNo,

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
