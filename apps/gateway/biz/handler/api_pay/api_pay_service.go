// Code generated by hertz generator.

package api_pay

import (
	"context"
	"fmt"
	"net/url"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	_ "github.com/hertz-contrib/swagger"
	api_pay "github.com/lzl-here/bt-shop-backend/apps/gateway/biz/model/api_pay"
	"github.com/lzl-here/bt-shop-backend/apps/gateway/global"
	pgen "github.com/lzl-here/bt-shop-backend/kitex_gen/pay"
	_ "github.com/swaggo/files"
)

func Pay(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api_pay.ApiPayReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	rpcReq := &pgen.PayReq{
		Subject:     req.Subject,
		TotalAmount: req.TotalAmount,
		TradeNo:     req.TradeNo,
	}
	rpcResp, err := global.PayClient.Pay(ctx, rpcReq)

	if err != nil {
		c.JSON(consts.StatusInternalServerError, &pgen.PayRsp{Code: 500, Msg: err.Error()})
		return
	}
	if rpcResp.Code != 1 {
		c.JSON(consts.StatusInternalServerError, &pgen.PayRsp{Code: 500, Msg: rpcResp.Msg})
		return
	}

	resp := &pgen.PayRsp{
		Code:  int32(rpcResp.Code),
		Msg:   rpcResp.Msg,
		LogId: "",
		Data: &pgen.PayRsp_PayRspData{
			PayUrl:  rpcResp.Data.PayUrl,
			TradeNo: rpcResp.Data.TradeNo,
		},
	}
	c.JSON(consts.StatusOK, resp)
}

// AlipayWebhook .
// @router /pay/webhook/alipay [POST]
func AlipayWebhook(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api_pay.ApiAlipayWebhookReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// 提取出 body
	s := string(c.Request.Body())
	fmt.Println(s)
	values, err := url.ParseQuery(s)
	if err != nil {
		fmt.Println(err)
		c.JSON(consts.StatusInternalServerError, &api_pay.ApiAlipayWebhookRsp{Code: 500, Msg: err.Error()})
		return
	}
	rpcReq := &pgen.AlipayWebhookReq{}
	rpcReq.AuthAppId = values.Get("auth_app_id")
	rpcReq.NotifyTime = values.Get("notify_time")
	rpcReq.NotifyType = values.Get("notify_type")
	rpcReq.NotifyId = values.Get("notify_id")
	rpcReq.AppId = values.Get("app_id")
	rpcReq.Charset = values.Get("charset")
	rpcReq.Version = values.Get("version")
	rpcReq.SignType = values.Get("sign_type")
	rpcReq.Sign = values.Get("sign")
	rpcReq.TradeNo = values.Get("trade_no")
	rpcReq.OutTradeNo = values.Get("out_trade_no")
	rpcReq.OutRequestNo = values.Get("out_request_no")
	rpcReq.OutBizNo = values.Get("out_biz_no")
	rpcReq.BuyerId = values.Get("buyer_id")
	rpcReq.BuyerLogonId = values.Get("buyer_logon_id")
	rpcReq.BuyerOpenId = values.Get("buyer_open_id")
	rpcReq.SellerId = values.Get("seller_id")
	rpcReq.SellerEmail = values.Get("seller_email")
	rpcReq.TradeStatus = values.Get("trade_status")
	rpcReq.RefundStatus = values.Get("refund_status")
	rpcReq.RefundReason = values.Get("refund_reason")
	rpcReq.RefundAmount = values.Get("refund_amount")
	rpcReq.TotalAmount = values.Get("total_amount")
	rpcReq.ReceiptAmount = values.Get("receipt_amount")
	rpcReq.InvoiceAmount = values.Get("invoice_amount")
	rpcReq.BuyerPayAmount = values.Get("buyer_pay_amount")
	rpcReq.PointAmount = values.Get("point_amount")
	rpcReq.RefundFee = values.Get("refund_fee")
	rpcReq.Subject = values.Get("subject")
	rpcReq.Body = values.Get("body")
	rpcReq.GmtCreate = values.Get("gmt_create")
	rpcReq.GmtPayment = values.Get("gmt_payment")
	rpcReq.GmtRefund = values.Get("gmt_refund")
	rpcReq.GmtClose = values.Get("gmt_close")
	rpcReq.FundBillList = values.Get("fund_bill_list")
	rpcReq.PassbackParams = values.Get("passback_params")
	rpcReq.VoucherDetailList = values.Get("voucher_detail_list")
	rpcReq.AgreementNo = values.Get("agreement_no")
	rpcReq.ExternalAgreementNo = values.Get("external_agreement_no")
	rpcReq.DbackStatus = values.Get("dback_status")
	rpcReq.DbackAmount = values.Get("dback_amount")
	rpcReq.BankAckTime = values.Get("bank_ack_time")
	resp, err := global.PayClient.AlipayWebhook(ctx, rpcReq)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, &api_pay.ApiAlipayWebhookRsp{Code: 500, Msg: err.Error()})
		fmt.Println(err)
		return
	}
	if resp.Code != 1 {
		c.JSON(consts.StatusInternalServerError, &api_pay.ApiAlipayWebhookRsp{Code: 500, Msg: resp.Msg})
		fmt.Printf("code: %s - code: %d", resp.Msg, resp.Code)
		return
	}
	// 返回200给通知支付宝已完成，不然支付宝会执行重试
	c.JSON(consts.StatusOK, "success")
}
