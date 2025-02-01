package handler

import (
	"github.com/lzl-here/bt-shop-backend/apps/order/internal/repo"
)

type OrderHandler struct {
	rep repo.RepoInterface
}

func NewOrderHandler(rep repo.RepoInterface) *OrderHandler {
	return &OrderHandler{
		rep: rep,
	}
}
