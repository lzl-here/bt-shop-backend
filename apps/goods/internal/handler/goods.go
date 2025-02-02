package handler

import (
	"context"
	"encoding/json"

	"github.com/lzl-here/bt-shop-backend/apps/goods/internal/domain/model"
	ggen "github.com/lzl-here/bt-shop-backend/kitex_gen/goods"
)

func (h *GoodsHandler) GetGoodsList(ctx context.Context, req *ggen.GetGoodsListReq) (*ggen.GetGoodsListRsp, error) {
	spus, err := h.rep.GetSpuListByIDs(ctx, req.SpuIds)
	if err != nil {
		return nil, err
	}
	spuIDs := make([]uint64, 0, len(spus))
	for _, s := range spus {
		spuIDs = append(spuIDs, s.ID)
	}
	skus, err := h.rep.GetSkuListBySpuIDs(ctx, spuIDs)
	if err != nil {
		return nil, err
	}
	specs, err := h.rep.GetSpecBySpuIDs(ctx, spuIDs)
	if err != nil {
		return nil, err
	}

	// 拼接数据给前端
	return buildRsp(spus, skus, specs), nil
}

func buildRsp(spus []*model.GoodsSpu, skus []*model.GoodsSku, specs []*model.Spec) *ggen.GetGoodsListRsp {
	skuSpuIDMap := make(map[uint64][]*model.GoodsSku)
	for _, sku := range skus {
		skuSpuIDMap[sku.SpuID] = append(skuSpuIDMap[sku.SpuID], sku)
	}
	rsp := &ggen.GetGoodsListRsp{}
	rsp.Code = 1
	rsp.Msg = "success"
	rsp.Data = &ggen.GetGoodsListRsp_GetGoodsListRspData{}
	specRsp := make([]*ggen.SpecInfo, 0)
	for _, s := range specs {
		specRsp = append(specRsp, &ggen.SpecInfo{
			SpecId:   s.ID,
			SpecName: s.Name,
			SpuId:    s.SpuID,
		})
	}
	for _, spu := range spus {
		rsp.Data.SpuList = append(rsp.Data.SpuList, &ggen.SpuInfo{
			SpuId:           spu.ID,
			SpuName:         spu.SpuName,
			SpuDesc:         spu.SpuDesc,
			SpuImgUrl:       spu.SpuImgUrl,
			SpuMinAmount:    spu.MinAmount,
			SpuMaxAmount:    spu.MaxAmount,
			SpuCategoryId:   spu.CategoryID,
			SpuCategoryName: spu.CategoryName,
			BrandId:         spu.BrandID,
			BrandName:       spu.BrandName,
			SkuList:         buildSkuList(skuSpuIDMap[spu.ID]), // 构建sku列表
			SpecList:        specRsp,                           // 构建spec列表
			SpecValueList:   buildAllSpecList(skus),            // 构建specValue列表
		})
	}
	return rsp
}

func buildSkuList(skus []*model.GoodsSku) []*ggen.SkuInfo {
	rsp := make([]*ggen.SkuInfo, 0)
	for _, s := range skus {
		rsp = append(rsp, &ggen.SkuInfo{
			SkuId:    s.ID,
			SpuId:    s.SpuID,
			SpecStr:  s.SpecValues,
			StockNum: s.StockNum,
		})
	}
	return rsp
}

func buildAllSpecList(skus []*model.GoodsSku) []*ggen.SpecValueInfo {
	rsp := make([]*ggen.SpecValueInfo, 0)
	for _, s := range skus {
		rsp = append(rsp, buildSpecList(s.SpecValues)...)
	}
	return rsp
}

func buildSpecList(specValueStr string) []*ggen.SpecValueInfo {
	specValues := new([]*model.SpecValue)
	json.Unmarshal([]byte(specValueStr), specValues)
	rsp := make([]*ggen.SpecValueInfo, 0)
	for _, s := range *specValues {
		rsp = append(rsp, &ggen.SpecValueInfo{
			SpecValueId: s.ID,
			SpecName:    s.Name,
			SpecValue:   s.Value,
		})
	}
	return rsp
}


func (h *GoodsHandler) GetGoodsDetail(ctx context.Context, req *ggen.GetGoodsDetailReq) (*ggen.GetGoodsDetailRsp, error) {
	return nil, nil
}

func (h *GoodsHandler) PublishGoods(ctx context.Context, req *ggen.PublishGoodsReq) (*ggen.PublishGoodsRsp, error) {
	return &ggen.PublishGoodsRsp{}, nil
}
