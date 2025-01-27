package api

import (
	"context"

	pgen "github.com/lzl-here/bt-shop-backend/kitex_gen/example"
	"github.com/lzl-here/bt-shop-backend/apps/pay/internal/handler"
	"github.com/lzl-here/bt-shop-backend/apps/pay/internal/repo"
)

// rpc入口

var _ (pgen.ExampleService) = (*ExampleServer)(nil)

type ExampleServer struct {
	pgen.ExampleService
	pingHandler   *handler.PingHandler
}

func NewExampleServer(rep *repo.Repo) *ExampleServer {
	return &ExampleServer{
		pingHandler:   handler.NewPingHandler(),
	}
}

func (s *ExampleServer) PingPong(ctx context.Context, req *pgen.PingReq) (resp *pgen.PingRsp, err error) {
	return s.pingHandler.PingPong(ctx, req)
}
