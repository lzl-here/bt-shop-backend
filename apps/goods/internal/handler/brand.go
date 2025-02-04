package handler

import (
	"context"

	"github.com/lzl-here/bt-shop-backend/apps/goods/internal/domain/model"
	ggen "github.com/lzl-here/bt-shop-backend/kitex_gen/goods"
)

/**
 * @description: 获取全部品牌列表
*/
func (h *GoodsHandler) GetBrandList(ctx context.Context, req *ggen.GetBrandListReq) (*ggen.GetBBrandListRsp, error) {
	brands, err := h.rep.GetAllBrandList(ctx)
	if err != nil {
		return nil, err
	}
	return buildBrandRsp(brands), nil
}

func buildBrandRsp(brands []*model.Brand) *ggen.GetBBrandListRsp {
	brandRsp := make([]*ggen.BaseBrand, 0)
	for _, b := range brands {
		brandRsp = append(brandRsp, &ggen.BaseBrand{
			BrandId:   b.ID,
			BrandName: b.Name,
		})
	}
	return &ggen.GetBBrandListRsp{
		Code: 1,
		Msg:  "success",
		Data: &ggen.GetBBrandListRsp_GetBBrandListDataRsp{
			BrandList: brandRsp,
		},
	}

}


