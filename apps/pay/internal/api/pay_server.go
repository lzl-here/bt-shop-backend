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

func (s *PayServer) Pay(ctx context.Context, req *pgen.PayReq) (res *pgen.PayRsp, err error) {
	return handler.NewPayHandler(s.rep).Pay(ctx, req)
}

func (s *PayServer) AlipayWebhook(ctx context.Context, req *pgen.AlipayWebhookReq) (res *pgen.AlipayWebhookRsp, err error) { 
	return handler.NewPayHandler(s.rep).AlipayWebhook(ctx, req)
}