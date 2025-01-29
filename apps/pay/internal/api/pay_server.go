package api

import (
	"context"

	"github.com/lzl-here/bt-shop-backend/apps/pay/internal/handler"
	"github.com/lzl-here/bt-shop-backend/apps/pay/internal/repo"
	pgen "github.com/lzl-here/bt-shop-backend/kitex_gen/pay"
)

// rpc入口

var _ (pgen.PayService) = (*PayServer)(nil)

type PayServer struct {
	rep *repo.Repo
}

func NewPayServer(rep *repo.Repo) *PayServer {
	return &PayServer{
		rep: rep,
	}
}

// 支付
func (s *PayServer) Pay(ctx context.Context, req *pgen.PayReq) (res *pgen.PayRsp, err error) {
	return handler.NewPayHandler(s.rep).Pay(ctx, req)
}

// 支付宝回调
func (s *PayServer) AlipayWebhook(ctx context.Context, req *pgen.AlipayWebhookReq) (res *pgen.AlipayWebhookRsp, err error) {
	
	return handler.NewPayHandler(s.rep).AlipayWebhook(ctx, req)
}

// 取消支付
func (s *PayServer) ClosePay(ctx context.Context, req *pgen.ClosePayReq) (res *pgen.ClosePayRsp, err error) {
	return handler.NewPayHandler(s.rep).ClosePay(ctx, req)
}

// 退款
func (s *PayServer) RefundPay(ctx context.Context, req *pgen.RefundPayReq) (res *pgen.RefundPayRsp, err error) {
	return handler.NewPayHandler(s.rep).RefundPay(ctx, req)
}
