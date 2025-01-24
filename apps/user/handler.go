package main

import (
	"context"
	user "github.com/lzl-here/bt-shop-backend/kitex_gen/user"
)

type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResponse, err error) {
	resp = new(user.Register)

	if err := mysql.CreateUser([]*model.User{
		{
			Name:      req.Name,
			Gender:    int64(req.Gender),
			Age:       req.Age,
			Introduce: req.Introduce,
		},
	}); err != nil {
		resp.Msg = err.Error()
		resp.Code = user.Code_DBErr
	}

	resp.Code = user.Code_Success
	return
}
