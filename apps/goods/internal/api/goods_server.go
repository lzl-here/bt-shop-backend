package api

import (
	"context"

	"github.com/lzl-here/bt-shop-backend/apps/goods/internal/handler"
	"github.com/lzl-here/bt-shop-backend/apps/goods/internal/repo"
	ggen "github.com/lzl-here/bt-shop-backend/kitex_gen/goods"
)

// rpc入口

var _ (ggen.GoodsService) = (*GoodsServer)(nil)

type GoodsServer struct {
	rep *repo.Repo
}

func NewGoodsServer(rep *repo.Repo) *GoodsServer {
	return &GoodsServer{
		rep: rep,
	}
}

func (s *GoodsServer) GetGoodsList(ctx context.Context, req *ggen.GetGoodsListReq) (res *ggen.GetGoodsListRsp, err error) {
	return handler.NewGoodsHandler(s.rep).GetGoodsList(ctx, req)
}
