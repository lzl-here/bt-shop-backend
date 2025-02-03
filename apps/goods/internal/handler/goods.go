package handler

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/lzl-here/bt-shop-backend/apps/goods/internal/domain/model"
	ggen "github.com/lzl-here/bt-shop-backend/kitex_gen/goods"
	"gorm.io/gorm"
)

/**
 * 获取商品列表
 */
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
			SpuPrice:        spu.Price,
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

/**
 * 获取商品详情，直接复用列表接口
 */
func (h *GoodsHandler) GetGoodsDetail(ctx context.Context, req *ggen.GetGoodsDetailReq) (*ggen.GetGoodsDetailRsp, error) {
	listRsp, err := h.GetGoodsList(ctx, &ggen.GetGoodsListReq{
		SpuIds: []uint64{req.SpuId},
	})
	if err != nil {
		return nil, err
	}
	return &ggen.GetGoodsDetailRsp{
		Code: 1,
		Msg:  "success",
		Data: &ggen.GetGoodsDetailRsp_GetGoodsDetailRspData{
			SpuInfo: listRsp.Data.SpuList[0],
		},
	}, nil
}

/**
 * 发布商品
 * 1. 发布参数
 * 2. 发布规格
 * 3. 发布spu 和 sku
 */
//TODO (mysql 和 es 都存一份, es中用和desc和name做keyword的模糊匹配, categoryID和brandID做强匹配)
func (h *GoodsHandler) PublishGoods(ctx context.Context, req *ggen.PublishGoodsReq) (*ggen.PublishGoodsRsp, error) {
	var err error
	var spu *model.GoodsSpu

	err = h.rep.Transaction(ctx, func(ctx context.Context, tx *gorm.DB) error {
		// 存储属性
		attrs := make([]*model.Attribute, 0)
		for _, a := range req.AttributeInfoList {
			attrs = append(attrs, &model.Attribute{
				Name:  a.AttributeName,
				Value: a.AttributeValue,
			})
		}
		attrs, err = h.rep.CreateAttributes(ctx, attrs)
		if err != nil {
			return err
		}
		// 存储规格
		specValues := make([]*model.SpecValue, 0)
		for _, sku := range req.SkuInfoList {
			for _, s := range sku.SpecKeyValue {
				specValues = append(specValues, &model.SpecValue{
					Name:   s.SpecName,
					Value:  s.SpecValue,
					SpecID: 0,
				})
			}
		}
		specs := make([]*model.Spec, 0)
		specMap := make(map[string]struct{})
		for _, s := range specValues {
			if _, ok := specMap[s.Name]; !ok {
				specMap[s.Name] = struct{}{}
			}
		}
		for n := range specMap {
			specs = append(specs, &model.Spec{
				Name: n,
			})
		}

		if specs, err = h.rep.CreateSpecs(ctx, specs); err != nil {
			return err
		}
		for _, s := range specs {
			for _, v := range specValues {
				if s.Name == v.Name {
					v.SpecID = s.ID
				}
			}
		}
		if err = h.rep.CreateSpecValues(ctx, specValues); err != nil {
			return err
		}

		// 存储spu
		specIDs := ""
		attrIDs := ""
		for _, s := range specs {
			specIDs += "," + strconv.FormatUint(s.ID, 10)
		}
		for _, a := range attrs {
			attrIDs += "," + strconv.FormatUint(a.ID, 10)
		}
		spu = &model.GoodsSpu{
			SpuName:      req.SpuName,
			SpuDesc:      req.SpuDesc,
			SpuImgUrl:    req.SpuImgUrl,
			Price:        req.SpuPrice,
			CategoryID:   req.CategoryId,
			CategoryName: req.CategoryName,
			BrandID:      req.BrandId,
			BrandName:    req.BrandName,
			Enabled:      true,
			AttributeIDs: attrIDs,
			SpecIDs:      specIDs,
		}
		if spu, err = h.rep.CreateSpu(ctx, spu); err != nil {
			return err
		}
		// 存储sku
		skus := make([]*model.GoodsSku, 0)
		for _, s := range req.SkuInfoList {
			bytes, err := json.Marshal(s.SpecKeyValue)
			if err != nil {
				return err
			}
			sku := model.GoodsSku{
				SpuID:      spu.ID,
				StockNum:   s.StockNum,
				SpecValues: string(bytes),
				Enabled:    true,
			}
			skus = append(skus, &sku)
		}
		if err = h.rep.CreateSkus(ctx, skus); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	//	SPU 添加到es
	go func() {
		err = h.rep.AddSpuToES(ctx, spu)
		if err != nil {
			klog.Error("add spu to es error", err)
		}
	}()

	return &ggen.PublishGoodsRsp{
		Code: 1,
		Msg:  "success",
		Data: &ggen.PublishGoodsRsp_PublishGoodsRspData{
			SpuId: spu.ID,
		},
	}, nil
}
