package handler

import (
	"context"

	pgen "github.com/lzl-here/bt-shop-backend/kitex_gen/example"
	"github.com/lzl-here/bt-shop-backend/pkg/verr"
	"github.com/lzl-here/bt-shop-backend/services/example/internal/repo"
)

// 做业务处理的部分

type PersonHandler struct {
	rep *repo.Repo
}

func NewPersonHandler(rep *repo.Repo) *PersonHandler {
	return &PersonHandler{rep: rep}
}

func (h *PersonHandler) GetPerson(ctx context.Context, req *pgen.GetPersonReq) (*pgen.GetPersonRsp, error) {
	person, err := h.rep.GetPersonByID(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	if person == nil {
		return nil, verr.ErrInternal
	}
	return &pgen.GetPersonRsp{
		Code: 0,
		Msg:  "ok",
		Data: &pgen.GetPersonRsp_PersonData{
			Uid:       person.ID,
			Name:      person.Name,
			Email:     person.Email,
			Password:  person.Password,
			Age:       person.Age,
			CreatedAt: person.CreatedString(),
			UpdatedAt: person.UpdatedString(),
			DeletedAt: person.DeletedString(),
		},
	}, nil
}
