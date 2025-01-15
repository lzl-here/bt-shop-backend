package handler

import (
	"context"

	pgen "github.com/lzl-here/bt-shop-backend/kitex_gen/example"
)

type PingHandler struct {
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) PingPong(ctx context.Context, req *pgen.PingReq) (resp *pgen.PingRsp, err error) {
	return &pgen.PingRsp{
		Code: 0,
		Msg:  "ok",
		Data: &pgen.PingRsp_PingData{
			Id:        1,
			Message:   "data message",
			CreatedAt: "2023-08-08 00:00:00",
			UpdatedAt: "2023-08-08 00:00:00",
			DeletedAt: "2023-08-08 00:00:00",
		},
	}, nil
}
