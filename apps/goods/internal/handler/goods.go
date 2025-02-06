package handler

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/lzl-here/bt-shop-backend/apps/goods/internal/domain/model"
	ggen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/goods"
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
	return buildGoodsListRsp(spus, skus, specs)
}

func buildGoodsListRsp(spus []*model.GoodsSpu, skus []*model.GoodsSku, specs []*model.Spec) (*ggen.GetGoodsListRsp, error) {
	skuSpuIDMap := make(map[uint64][]*model.GoodsSku)
	for _, sku := range skus {
		skuSpuIDMap[sku.SpuID] = append(skuSpuIDMap[sku.SpuID], sku)
	}
	rsp := &ggen.GetGoodsListRsp{}
	rsp.Code = 1
	rsp.Msg = "success"
	rsp.Data = &ggen.GetGoodsListRsp_GetGoodsListRspData{}
	specRsp := make([]*ggen.BaseSpec, 0)
	for _, s := range specs {
		specRsp = append(specRsp, &ggen.BaseSpec{
			Id:       s.ID,
			SpecName: s.SpecName,
			SpuId:    s.SpuID,
		})
	}

	for _, spu := range spus {
		skuList, err := buildSkuList(skuSpuIDMap[spu.ID])
		if err != nil {
			return nil, err
		}
		rsp.Data.SpuList = append(rsp.Data.SpuList, &ggen.GoodsSpu{
			Spu: &ggen.BaseSpu{
				Id:           spu.ID,
				SpuName:      spu.SpuName,
				SpuDesc:      spu.SpuDesc,
				SpuImgUrl:    spu.SpuImgUrl,
				SpuPrice:     spu.SpuPrice,
				CategoryId:   spu.CategoryID,
				CategoryName: spu.CategoryName,
				BrandId:      spu.BrandID,
				BrandName:    spu.BrandName,
			},
			SkuList:  skuList, // 构建sku列表
			SpecList: specRsp, // 构建spec列表
		})
	}
	return rsp, nil
}

func buildSkuList(skus []*model.GoodsSku) ([]*ggen.GoodsSku, error) {
	rsp := make([]*ggen.GoodsSku, 0)

	for _, s := range skus {
		specValues := make([]*ggen.BaseSpecValue, 0)
		if err := json.Unmarshal([]byte(s.SpecValues), &specValues); err != nil {
			return nil, err
		}
		rsp = append(rsp, &ggen.GoodsSku{
			Sku: &ggen.BaseSku{
				Id:       s.ID,
				SpuId:    s.SpuID,
				StockNum: s.StockNum,
				SkuPrice: s.SkuPrice,
			},
			SpecValueList: specValues,
		})
	}
	return rsp, nil
}

func buildAllSpecList(skus []*model.GoodsSku) []*ggen.BaseSpecValue {
	rsp := make([]*ggen.BaseSpecValue, 0)
	for _, s := range skus {
		rsp = append(rsp, buildSpecList(s.SpecValues)...)
	}
	return rsp
}

func buildSpecList(specValueStr string) []*ggen.BaseSpecValue {
	specValues := new([]*model.SpecValue)
	json.Unmarshal([]byte(specValueStr), specValues)
	rsp := make([]*ggen.BaseSpecValue, 0)
	for _, s := range *specValues {
		rsp = append(rsp, &ggen.BaseSpecValue{
			SpecId:    s.ID,
			SpecName:  s.SpecName,
			SpecValue: s.SpecValue,
		})
	}
	return rsp
}

/**
 * 获取商品详情，直接复用列表接口
 */
func (h *GoodsHandler) GetGoodsDetail(ctx context.Context, req *ggen.GetGoodsDetailReq) (*ggen.GetGoodsDetailRsp, error) {
	spus, err := h.rep.GetSpuListByIDs(ctx, []uint64{req.SpuId})
	if err != nil {
		return nil, err
	}
	if len(spus) == 0 {
		return &ggen.GetGoodsDetailRsp{
			Code: 1,
			Msg:  "success",
		}, nil
	}
	spu := spus[0]
	spuIDs := make([]uint64, 0, len(spus))
	spuIDs = append(spuIDs, spu.ID)
	skus, err := h.rep.GetSkuListBySpuIDs(ctx, spuIDs)
	if err != nil {
		return nil, err
	}
	specs, err := h.rep.GetSpecBySpuIDs(ctx, spuIDs)
	if err != nil {
		return nil, err
	}
	return buildGoodsDetailRsp(spu, skus, specs)
}

func buildGoodsDetailRsp(spu *model.GoodsSpu, skus []*model.GoodsSku, specs []*model.Spec) (*ggen.GetGoodsDetailRsp, error) {

	spuRsp := &ggen.GoodsSpu{
		Spu: &ggen.BaseSpu{
			Id:           spu.ID,
			SpuName:      spu.SpuName,
			SpuDesc:      spu.SpuDesc,
			CategoryId:   spu.CategoryID,
			CategoryName: spu.CategoryName,
			BrandId:      spu.BrandID,
			BrandName:    spu.BrandName,
			SpuPrice:     spu.SpuPrice,
			Enabled:      spu.Enabled,
			SpuImgUrl:    spu.SpuImgUrl,
		},
	}
	// 构建sku 和 规格项
	specMap := make(map[uint64]*ggen.BaseSpec)
	skuRsps := make([]*ggen.GoodsSku, 0)
	for _, s := range skus {
		// sku
		skuRsp := &ggen.GoodsSku{
			Sku: &ggen.BaseSku{
				Id:       s.ID,
				SpuId:    s.SpuID,
				StockNum: s.StockNum,
				SkuPrice: s.SkuPrice,
			},
		}
		specValues := make([]*ggen.BaseSpecValue, 0)
		if err := json.Unmarshal([]byte(s.SpecValues), &specValues); err != nil {
			return nil, err
		}
		skuRsp.SpecValueList = specValues
		skuRsps = append(skuRsps, skuRsp)
		// 规格项
		for _, s := range skuRsp.SpecValueList {
			specMap[s.SpecId] = &ggen.BaseSpec{
				Id:       s.SpecId,
				SpecName: s.SpecName,
				SpuId:    spu.ID,
			}
		}
	}
	spuRsp.SkuList = skuRsps
	// 构建规格项
	specList := make([]*ggen.BaseSpec, 0)
	for _, s := range specMap {
		specList = append(specList, s)
	}
	spuRsp.SpecList = specList
	return &ggen.GetGoodsDetailRsp{
		Code: 1,
		Msg:  "success",
		Data: &ggen.GetGoodsDetailRsp_GetGoodsDetailRspData{
			Spu: spuRsp,
		},
	}, nil
}

/**
 * 发布商品
 * 1. 发布参数
 * 2. 发布规格
 * 3. 发布spu 和 sku
 */
func (h *GoodsHandler) PublishGoods(ctx context.Context, req *ggen.PublishGoodsReq) (*ggen.PublishGoodsRsp, error) {
	var err error
	var spu *model.GoodsSpu

	err = h.rep.Transaction(ctx, func(ctx context.Context, tx *gorm.DB) error {
		// 存储属性
		attrs := make([]*model.Attribute, 0)
		for _, a := range req.AttributeList {
			attrs = append(attrs, &model.Attribute{
				Name:  a.AttributeName,
				Value: a.AttributeValue,
			})
		}
		attrs, err = h.rep.CreateAttributes(ctx, h.rep.GetDB(), attrs)
		if err != nil {
			return err
		}
		// 存储规格项和规格值
		specs, specValues, err := h.storeSpecsAndValues(ctx, req)
		if err != nil {
			return err
		}
		if spu, err = h.storeSpuAndSkus(ctx, req, specs, specValues, attrs); err != nil {
			return err
		}
		specIDs := make([]uint64, 0, len(specs))
		for _, s := range specs {
			specIDs = append(specIDs, s.ID)
		}
		if err = h.rep.UpdateSpecs(ctx, h.rep.GetDB(), specIDs, &model.Spec{SpuID: spu.ID}); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	//	SPU 添加到es
	// go func() {
	err = h.rep.AddSpuToES(ctx, spu)
	if err != nil {
		klog.Error("add spu to es error", err)
	}
	// }()

	return &ggen.PublishGoodsRsp{
		Code: 1,
		Msg:  "success",
		Data: &ggen.PublishGoodsRsp_PublishGoodsRspData{
			SpuId: spu.ID,
		},
	}, nil
}

func getIDs(specs []*model.Spec, attrs []*model.Attribute) (string, string) {
	specIDs := ""
	attrIDs := ""
	if len(specs) > 0 {
		specIDs = strconv.FormatUint(specs[0].ID, 10)
		for i := 1; i < len(specs); i++ {
			specIDs += "," + strconv.FormatUint(specs[i].ID, 10)
		}
	}
	if len(attrs) > 0 {
		attrIDs = strconv.FormatUint(attrs[0].ID, 10)
		for i := 1; i < len(attrs); i++ {
			attrIDs += "," + strconv.FormatUint(attrs[i].ID, 10)
		}
	}
	return specIDs, attrIDs
}

// 存储规格项和规格值，返回带有id的规格项和规格值
func (h *GoodsHandler) storeSpecsAndValues(ctx context.Context, req *ggen.PublishGoodsReq) ([]*model.Spec, []*model.SpecValue, error) {
	var err error
	specValues := make([]*model.SpecValue, 0)
	for _, sku := range req.SkuList {
		for _, s := range sku.SpecKeyValue {
			specValues = append(specValues, &model.SpecValue{
				SpecName:  s.SpecName,
				SpecValue: s.SpecValue,
				SpecID:    0,
			})
		}
	}
	specs := make([]*model.Spec, 0)
	specMap := make(map[string]struct{})
	for _, s := range specValues {
		if _, ok := specMap[s.SpecName]; !ok {
			specMap[s.SpecName] = struct{}{}
		}
	}
	for n := range specMap {
		specs = append(specs, &model.Spec{
			SpecName: n,
		})
	}

	if specs, err = h.rep.CreateSpecs(ctx, h.rep.GetDB(), specs); err != nil {
		return nil, nil, err
	}
	for _, s := range specs {
		for _, v := range specValues {
			if s.SpecName == v.SpecName {
				v.SpecID = s.ID
			}
		}
	}
	if err = h.rep.CreateSpecValues(ctx, h.rep.GetDB(), specValues); err != nil {
		return nil, nil, err
	}
	return specs, specValues, nil
}

// 存储spu 和 sku
func (h *GoodsHandler) storeSpuAndSkus(ctx context.Context, req *ggen.PublishGoodsReq, specs []*model.Spec, specValues []*model.SpecValue, attrs []*model.Attribute) (*model.GoodsSpu, error) {
	var err error
	// 存储spu
	attrIDs, specIDs := getIDs(specs, attrs)
	brand, err := h.rep.GetBrandByID(ctx, req.BrandId)
	if err != nil {
		return nil, err
	}
	spu := &model.GoodsSpu{
		SellerID:     req.UserId,
		ShopID:       req.ShopId,
		SpuName:      req.SpuName,
		SpuDesc:      req.SpuDesc,
		SpuImgUrl:    req.SpuImgUrl,
		SpuPrice:     req.SpuPrice,
		CategoryID:   req.CategoryId,
		CategoryName: req.CategoryName,
		BrandID:      req.BrandId,
		BrandName:    brand.Name,
		BrandIconUrl: brand.IconUrl,
		Enabled:      true,
		AttributeIDs: attrIDs,
		SpecIDs:      specIDs,
	}
	if spu, err = h.rep.CreateSpu(ctx, h.rep.GetDB(), spu); err != nil {
		return nil, err
	}
	// 存储sku
	// 前端传来的规格和规格项是没有id的，只有存入DB后才有id，这里存入规格值json的时候需要存入带id的
	// 规格值的 name_value 是独一无二的，根据这个作为key来获取带有id的规格值
	specValueMap := make(map[string]*model.SpecValue)
	for _, s := range specValues {
		specValueMap[s.SpecName+"_"+s.SpecValue] = s
	}
	skus := make([]*model.GoodsSku, 0)
	for _, s := range req.SkuList {
		skuSpecValues := make([]*model.SpecValue, 0)
		for _, sv := range s.SpecKeyValue {
			skuSpecValues = append(skuSpecValues, specValueMap[sv.SpecName+"_"+sv.SpecValue])
		}
		bytes, err := json.Marshal(skuSpecValues)
		if err != nil {
			return nil, err
		}
		sku := model.GoodsSku{
			SpuID:      spu.ID,
			StockNum:   s.StockNum,
			SkuPrice:   s.SkuPrice,
			SpecValues: string(bytes),
			Enabled:    true,
		}
		skus = append(skus, &sku)
	}
	if err = h.rep.CreateSkus(ctx, h.rep.GetDB(), skus); err != nil {
		return nil, err
	}
	return spu, nil
}

func (h *GoodsHandler) GetSellerGoodsList(ctx context.Context, req *ggen.GetSellerGoodsListReq) (*ggen.GetSellerGoodsListRsp, error) {
	spus, err := h.rep.GetSpuListByShopID(ctx, req.ShopId, req.PageNo, req.PageSize)
	if err != nil {
		return nil, err
	}
	total, err := h.rep.GetSpuTotalByShopID(ctx, req.ShopId)
	if err != nil {
		return nil, err
	}
	return buildSellerGoodsListRsp(spus, total, req.PageNo, req.PageSize)
}

func buildSellerGoodsListRsp(spus []*model.GoodsSpu, total int32, pageNo int32, pageSize int32) (*ggen.GetSellerGoodsListRsp, error) {
	rsp := &ggen.GetSellerGoodsListRsp{}
	rsp.Code = 1
	rsp.Msg = "success"
	rsp.Data = &ggen.GetSellerGoodsListRsp_GetSellerGoodsListRspData{
		Total:    total,
		PageNo:   pageNo,
		PageSize: pageSize,
	}
	for _, s := range spus {
		spu, err := s.ToGen()
		if err != nil {
			return nil, err
		}
		rsp.Data.SpuList = append(rsp.Data.SpuList, spu)
	}
	return rsp, nil
}
