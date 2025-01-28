package api

import (
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
