package handler

import (
	"github.com/lzl-here/bt-shop-backend/apps/goods/internal/repo"
)

type GoodsHandler struct {
	rep repo.RepoInterface
}

func NewGoodsHandler(rep repo.RepoInterface) *GoodsHandler {
	return &GoodsHandler{
		rep: rep,
	}
}
