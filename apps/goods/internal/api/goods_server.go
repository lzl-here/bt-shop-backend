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

// 获取商品详情
func (s *GoodsServer) GetGoodsDetail(ctx context.Context, req *ggen.GetGoodsDetailReq) (res *ggen.GetGoodsDetailRsp, err error) {
	return handler.NewGoodsHandler(s.rep).GetGoodsDetail(ctx, req)
}

// 获取商品列表的详情
func (s *GoodsServer) GetGoodsList(ctx context.Context, req *ggen.GetGoodsListReq) (res *ggen.GetGoodsListRsp, err error) {
	return handler.NewGoodsHandler(s.rep).GetGoodsList(ctx, req)
}

// 获取关键词下拉框
func (s *GoodsServer) GetKeywordDownList(ctx context.Context, req *ggen.GetKeywordDownListReq) (res *ggen.GetKeywordDownListRsp, err error) {
	return handler.NewGoodsHandler(s.rep).GetKeywordDownList(ctx, req)
}

// 获取分类列表
func (s *GoodsServer) GetCategoryList(ctx context.Context, req *ggen.GetCategoryListReq) (*ggen.GetCategoryListRsp, error) {
	return handler.NewGoodsHandler(s.rep).GetCategoryList(ctx, req)
}

func (s *GoodsServer) GetBrandList(ctx context.Context, req *ggen.GetBrandListReq) (*ggen.GetBBrandListRsp, error) {
	return handler.NewGoodsHandler(s.rep).GetBrandList(ctx, req)
}

// 搜索商品列表
func (s *GoodsServer) SearchSpuList(ctx context.Context, req *ggen.SearchSpuListReq) (res *ggen.SearchSpuListRsp, err error) {
	return handler.NewGoodsHandler(s.rep).SearchSpuList(ctx, req)
}

// 添加到购物车
func (s *GoodsServer) AddToCart(ctx context.Context, req *ggen.AddToCartReq) (res *ggen.AddToCartRsp, err error) {
	return handler.NewGoodsHandler(s.rep).AddToCart(ctx, req)
}

// 发布商品
func (s *GoodsServer) PublishGoods(ctx context.Context, req *ggen.PublishGoodsReq) (res *ggen.PublishGoodsRsp, err error) {
	return handler.NewGoodsHandler(s.rep).PublishGoods(ctx, req)
}
