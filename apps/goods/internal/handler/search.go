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
	if req.PageNo < 1 {
		req.PageNo = 1
	}
	if req.PageSize > 100 {
		req.PageSize = 100
	}
	spus, err := h.rep.SearchSpu(ctx, req)
	if err != nil {
		return nil, err
	}
	return buildSearchRsp(spus, req.PageSize, req.PageNo), nil
}

func buildSearchRsp(spus []*model.GoodsSpu, pageSize int32, pageNo int32) *ggen.SearchSpuListRsp {
	spusRsp := make([]*ggen.BaseSpuInfo, 0)
	for _, s := range spus {
		spusRsp = append(spusRsp, &ggen.BaseSpuInfo{
			SpuId:           s.ID,
			SpuName:         s.SpuName,
			SpuDesc:         s.SpuDesc,
			SpuImgUrl:       s.SpuImgUrl,
			SpuPrice:        s.Price,
			SpuCategoryId:   s.CategoryID,
			SpuCategoryName: s.CategoryName,
			BrandId:         s.BrandID,
			BrandName:       s.BrandName,
		})
	}

	return &ggen.SearchSpuListRsp{
		Code: 1,
		Msg:  "success",
		Data: &ggen.SearchSpuListRsp_SearchSpuListRspData{
			PageSize: pageSize,
			PageNo:   pageNo,
			SpuList:  spusRsp,
		},
	}
}

/**
* @Description: 获取搜索关键词列表, 同时增加搜索关键词次数
 */
func (h *GoodsHandler) GetKeywordDownList(ctx context.Context, req *ggen.GetKeywordDownListReq) (*ggen.GetKeywordDownListRsp, error) {
	// 异步增加搜索关键词次数
	// TODO 写聚合
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
