package handler

import (
	"context"

	"github.com/lzl-here/bt-shop-backend/apps/example/internal/repo"
	bizerr "github.com/lzl-here/bt-shop-backend/apps/example/pkg"
	pgen "github.com/lzl-here/bt-shop-backend/kitex_gen/example"
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
		return nil, bizerr.ErrResourceNotFound
	}
	// bizErr := kerrors.NewGRPCBizStatusErrorWithExtra(404, "not foudadadawednd", map[string]string{"key" : "value"})
	// grpcStatusErr := bizErr.(kerrors.GRPCStatusIface)
	// st, _ := grpcStatusErr.GRPCStatus().WithDetails()
	// grpcStatusErr.SetGRPCStatus(st)
	// return nil, bizErr

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
