package handler

import (
	"github.com/lzl-here/bt-shop-backend/apps/user/internal/repo"
)

type UserHandler struct {
	rep repo.RepoInterface
}

func NewUserHandler(rep repo.RepoInterface) *UserHandler {
	return &UserHandler{
		rep: rep,
	}
}
