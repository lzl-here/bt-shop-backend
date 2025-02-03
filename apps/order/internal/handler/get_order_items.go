package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/lzl-here/bt-shop-backend/apps/order/internal/domain/model"
	ogen "github.com/lzl-here/bt-shop-backend/kitex_gen/order"
)

func (h *OrderHandler) GetOrderItems(ctx context.Context, req *ogen.GetOrderItemsReq) (*ogen.GetOrderItemsRsp, error) {
	items, err := h.rep.GetOrderItems(ctx, &model.OrderItem{TradeNo: req.TradeNo})
	if err != nil {
		return nil, err
	}

	itemRsp := make([]*ogen.BaseOrderItem, 0)
	for _, item := range items {
		specValues := make([]*ogen.BaseSpecValue, 0)
		err = json.Unmarshal([]byte(item.SpecValues), &specValues)
		if err != nil {
			return nil, err
		}
		itemRsp = append(itemRsp, &ogen.BaseOrderItem{
			SpuId:         item.SpuID,
			SkuId:         item.SkuID,
			SpuName:       item.SpuName,
			CategoryId:    item.CategoryID,
			CategoryName:  item.CategoryName,
			BrandId:       item.BrandID,
			BrandName:     item.BrandName,
			SkuImgUrl:     item.ItemImgUrl,
			SpecValueList: specValues,
			SkuAmount:     item.SkuAmount,
			TradeNo:       item.TradeNo,
			OrderNo:       item.OrderNo,
			ShopId:        item.ShopID,
			SellerId:      item.SellerID,
			BuyerId:       item.BuyerID,
		})
	}
	return &ogen.GetOrderItemsRsp{
		Code: 1,
		Msg:  "success",
		Data: &ogen.GetOrderItemsRsp_GetOrderItemsRspData{
			OrderItemList: itemRsp,
		},
	}, nil
}
