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

func (s *GoodsServer) GetKeywordDownList(ctx context.Context, req *ggen.GetKeywordDownListReq) (res *ggen.GetKeywordDownListRsp, err error) {
	return handler.NewGoodsHandler(s.rep).GetKeywordDownList(ctx, req)
}

func (s *GoodsServer) GetGoodsDetail(ctx context.Context, req *ggen.GetGoodsDetailReq) (res *ggen.GetGoodsDetailRsp, err error) {
	return handler.NewGoodsHandler(s.rep).GetGoodsDetail(ctx, req)
}

func (s *GoodsServer) GetCategoryList(ctx context.Context, req *ggen.GetCategoryListReq) (*ggen.GetCategoryListRsp, error) {
	return handler.NewGoodsHandler(s.rep).GetCategoryList(ctx, req)
}

func (s *GoodsServer) SearchSpuList(ctx context.Context, req *ggen.SearchSpuListReq) (res *ggen.SearchSpuListRsp, err error) {
	return handler.NewGoodsHandler(s.rep).SearchSpuList(ctx, req)
}

func (s *GoodsServer) AddToCart(ctx context.Context, req *ggen.AddToCartReq) (res *ggen.AddToCartRsp, err error) {
	return handler.NewGoodsHandler(s.rep).AddToCart(ctx, req)
}

func (s *GoodsServer) PublishGoods(ctx context.Context, req *ggen.PublishGoodsReq) (res *ggen.PublishGoodsRsp, err error) {
	return handler.NewGoodsHandler(s.rep).PublishGoods(ctx, req)
}
