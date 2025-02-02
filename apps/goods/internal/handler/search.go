package handler

import (
	"context"

	"github.com/lzl-here/bt-shop-backend/apps/goods/internal/domain/model"
	ggen "github.com/lzl-here/bt-shop-backend/kitex_gen/goods"
)

/**
* @Description: 搜索商品
 */
func (h *GoodsHandler) SearchSpuList(ctx context.Context, req *ggen.SearchSpuListReq) (*ggen.SearchSpuListRsp, error) {
	// TODO es
	return nil, nil
}

/**
* @Description: 获取搜索关键词列表, 同时增加搜索关键词次数
 */
func (h *GoodsHandler) GetKeywordDownList(ctx context.Context, req *ggen.GetKeywordDownListReq) (*ggen.GetKeywordDownListRsp, error) {
	// 异步增加搜索关键词次数
	go func() {
		h.rep.UpdateOrSaveKeywordTimes(ctx, req.Prefix)
	}()

	list, err := h.rep.GetKeywordDownList(ctx, req.Prefix, 10)
	if err != nil {
		return nil, err
	}
	return buildKeyDownListRsp(list), nil
}

func buildKeyDownListRsp(list []*model.Keyword) *ggen.GetKeywordDownListRsp {
	rsp := &ggen.GetKeywordDownListRsp{
		Code: 1,
		Msg:  "success",
		Data: &ggen.GetKeywordDownListRsp_GetKeyDownListRspData{
			KeywordList: make([]*ggen.Keyword, 0),
		},
	}
	for _, keyword := range list {
		rsp.Data.KeywordList = append(rsp.Data.KeywordList, &ggen.Keyword{
			Id:    keyword.ID,
			Value: keyword.Value,
		})
	}
	return rsp
}
