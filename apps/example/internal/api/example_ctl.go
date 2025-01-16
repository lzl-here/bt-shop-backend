package api

import (
	"context"

	pgen "github.com/lzl-here/bt-shop-backend/kitex_gen/example"
	"github.com/lzl-here/bt-shop-backend/apps/example/internal/service/handler"
	"github.com/lzl-here/bt-shop-backend/apps/example/internal/repo"
)

// rpc入口

var _ (pgen.ExampleService) = (*ExampleServer)(nil)

type ExampleServer struct {
	pgen.ExampleService
	pingHandler   *handler.PingHandler
	personHandler *handler.PersonHandler
}

func NewExampleServer(rep *repo.Repo) *ExampleServer {
	return &ExampleServer{
		pingHandler:   handler.NewPingHandler(),
		personHandler: handler.NewPersonHandler(rep),
	}
}

func (s *ExampleServer) PingPong(ctx context.Context, req *pgen.PingReq) (resp *pgen.PingRsp, err error) {
	return s.pingHandler.PingPong(ctx, req)
}

func (s *ExampleServer) GetPerson(ctx context.Context, req *pgen.GetPersonReq) (res *pgen.GetPersonRsp, err error) {
	return s.personHandler.GetPerson(ctx, req)
}
