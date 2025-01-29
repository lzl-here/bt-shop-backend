package handler

import (
	"github.com/lzl-here/bt-shop-backend/apps/pay/internal/repo"
)

type PayHandler struct {
	rep repo.RepoInterface
}

func NewPayHandler(rep repo.RepoInterface) *PayHandler {
	return &PayHandler{
		rep: rep,
	}
}
