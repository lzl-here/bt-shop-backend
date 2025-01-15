package service

import (
	"context"

	pgen "github.com/lzl-here/bt-shop-backend/kitex_gen/person"
)

type PersonServiceImpl struct {
	pgen.PersonService
}

func NewPersonService() *PersonServiceImpl {
	return &PersonServiceImpl{}
}

func (s *PersonServiceImpl) GetPersonInfo(ctx context.Context, req *pgen.GetPersonReq) (resp *pgen.GetPersonRsp, err error) {
	return &pgen.GetPersonRsp{
		Code:  0,
		Msg:   "success",
		LogId: "",
		Data: &pgen.GetPersonRsp_PersonInfoData{
			Id:        1,
			NickName:  "lzl",
			AvatarUrl: "https://avatars.githubusercontent.com/u/10248107?v=4",
			Desc:      "hello world",
			CreatedAt: "2023-07-01 00:00:00",
			UpdatedAt: "2023-07-01 00:00:00",
			DeletedAt: "2023-07-01 00:00:00",
		},
	}, nil
}
