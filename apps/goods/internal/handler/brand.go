package handler

import (
	"context"

	"github.com/lzl-here/bt-shop-backend/apps/goods/internal/domain/model"
	ggen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/goods"
)

/**
 * @description: 获取全部品牌列表
 */
func (h *GoodsHandler) GetBrandList(ctx context.Context, req *ggen.GetBrandListReq) (*ggen.GetBrandListRsp, error) {
	brands, err := h.rep.GetAllBrandList(ctx)
	if err != nil {
		return nil, err
	}
	return buildBrandRsp(brands), nil
}

func buildBrandRsp(brands []*model.Brand) *ggen.GetBrandListRsp {
	brandRsp := make([]*ggen.BaseBrand, 0)
	for _, b := range brands {
		brandRsp = append(brandRsp, &ggen.BaseBrand{
			BrandId:   b.ID,
			BrandName: b.Name,
			IconUrl:   b.IconUrl,
		})
	}
	return &ggen.GetBrandListRsp{
		Code: 1,
		Msg:  "success",
		Data: &ggen.GetBrandListRsp_GetBrandListDataRsp{
			BrandList: brandRsp,
		},
	}

}
