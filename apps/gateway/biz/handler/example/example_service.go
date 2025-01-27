// Code generated by hertz generator.

package example

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	example "github.com/lzl-here/bt-shop-backend/apps/gateway/biz/model/example"
)

// PingPong .
// @router /ping [GET]
func PingPong(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.PingReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(example.PingRsp)
	resp.Code = 1
	resp.Msg = "success"
	resp.LogId = "123"
	resp.Data = &example.PingRsp_PingData{
		Id:        1,
		Message:   "hello",
		CreatedAt: "2023-07-01 00:00:00",
		UpdatedAt: "2023-07-01 00:00:00",
		DeletedAt: "2023-07-01 00:00:00",
	}
	resp.Data = nil
	c.JSON(consts.StatusOK, resp)
}

// GetPerson .
// @router /get_person [GET]
func GetPerson(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.GetPersonReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(example.GetPersonRsp)

	c.JSON(consts.StatusOK, resp)
}
