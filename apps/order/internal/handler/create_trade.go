package handler

import (
	"context"
	"time"

	"github.com/lzl-here/bt-shop-backend/apps/order/internal/common"
	"github.com/lzl-here/bt-shop-backend/apps/order/internal/constant"
	"github.com/lzl-here/bt-shop-backend/apps/order/internal/domain/model"
	"github.com/lzl-here/bt-shop-backend/pkg/utils"
	"gorm.io/gorm"

	ggen "github.com/lzl-here/bt-shop-backend/kitex_gen/goods"
	ogen "github.com/lzl-here/bt-shop-backend/kitex_gen/order"
	pgen "github.com/lzl-here/bt-shop-backend/kitex_gen/pay"
	bizerr "github.com/lzl-here/bt-shop-backend/pkg/err"
)

// 1. 校验商品信息合法性
// 2. 创建交易、订单、订单项
// 3. 拉起支付
// 4. 返回支付页面给前端
func (h *OrderHandler) CreateTrade(ctx context.Context, req *ogen.CreateTradeReq) (res *ogen.CreateTradeRsp, err error) {
	// 幂等处理
	ok, cleaner, _, err := utils.NoDuplicate(ctx, h.rep.GetCache(), "order", req.Token, 30*time.Second)
	if !ok {
		return nil, bizerr.ErrDuplicateReq
	}
	if err != nil {
		cleaner(5)
		return nil, err
	}

	// 1. 校验商品信息
	if err := h.checkParam(ctx, req); err != nil {
		cleaner(5)
		return nil, err
	}
	// 2. 创建交易、订单、订单项
	var tradeNo string
	err = h.rep.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error
		var trade *model.Trade
		var orders []*model.Order
		var orderItems []*model.OrderItem
		trade, orders, orderItems, tradeNo = buildModels(req)
		// 创建交易
		err = h.rep.CreateTrade(ctx, trade)
		if err != nil {
			return err
		}
		// 创建订单
		err = h.rep.CreateOrders(ctx, orders)
		if err != nil {
			return err
		}
		// 创建订单项
		err = h.rep.CreateOrderItems(ctx, orderItems)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	// 3. 拉起支付
	payRsp, err := h.rep.Pay(ctx, &pgen.PayReq{
		Subject:     "这是支付的标题",
		TotalAmount: req.TradeInfo.TradeAmount,
		TradeNo:     tradeNo,
	})
	if err != nil {
		return nil, err
	}
	if payRsp.Code != 1 {
		return nil, bizerr.ErrPayFailed
	}
	// 4. 返回支付页面url
	return &ogen.CreateTradeRsp{
		Code: 1,
		Msg:  "success",
		Data: &ogen.CreateTradeRsp_CreateTradeRspData{
			TradeNo:    tradeNo,
			PayPageUrl: payRsp.Data.PayPageUrl,
		},
	}, nil
}

// 检查商品合法性:
// 1. 前端传来的sku是否存在
// 2. 前端传来的金额是否合法
func (h *OrderHandler) checkParam(ctx context.Context, req *ogen.CreateTradeReq) error {
	goodsRsp, err := h.rep.GetGoodsList(ctx, &ggen.GetGoodsListReq{})
	if err != nil {
		return nil
	}
	if goodsRsp.Code != 1 {
		return bizerr.ErrDownStream
	}
	// 后端存的数据
	itemSkuIDMap := make(map[uint64]*ggen.SkuInfo)
	for _, spu := range goodsRsp.Data.SpuList {
		for _, sku := range spu.SkuList {
			itemSkuIDMap[sku.SkuId] = sku
		}
	}
	// 前端传来的数据
	for _, o := range req.TradeInfo.OrderInfoList {
		for _, oi := range o.OrderItemList {
			// 不存在
			if _, ok := itemSkuIDMap[oi.SkuId]; !ok {
				return bizerr.ErrGoodsNotExist
			}
			// 金额有问题
			skuFromFront := itemSkuIDMap[oi.SkuId]
			if skuFromFront.SkuAmount != oi.SkuAmount {
				return bizerr.ErrGoodsNotCorrect
			}
		}
	}
	// 这里sku的金额没问题了，校验order和trade的金额
	// 订单金额 = sku金额之和
	// 交易金额 = 订单金额之和
	totalTradeAmount := "0"
	for _, o := range req.TradeInfo.OrderInfoList {
		if err != nil {
			return bizerr.ErrInvalidTradeAmount
		}
		totalOrderAmount := "0"
		for _, s := range o.OrderItemList {
			totalOrderAmount = common.AmountPlus(totalOrderAmount, s.SkuAmount)
		}
		if !common.AmountEqual(totalOrderAmount, o.OrderAmount) {
			return bizerr.ErrInvalidTradeAmount
		}
		totalTradeAmount = common.AmountPlus(totalTradeAmount, o.OrderAmount)
	}
	if !common.AmountEqual(totalTradeAmount, req.TradeInfo.TradeAmount) {
		return bizerr.ErrInvalidTradeAmount
	}
	return nil
}

// 构建model
func buildModels(req *ogen.CreateTradeReq) (*model.Trade, []*model.Order, []*model.OrderItem, string) {
	tradeNo := common.GenTradeNo()
	orderRsp := make([]*model.Order, 0)
	itemRSp := make([]*model.OrderItem, 0)
	// trade
	trade := &model.Trade{
		TradeNo:     tradeNo,
		TradeAmount: req.TradeInfo.TradeAmount,
		TradeStatus: constant.TradeStatusPaying,
	}
	reqOrders := req.TradeInfo.OrderInfoList
	// order
	for _, o := range reqOrders {
		orderNo := common.GenOrderNo()
		orderRsp = append(orderRsp, &model.Order{
			TradeNo:     tradeNo,
			OrderNo:     orderNo,
			OrderAmount: o.OrderAmount,
			OrderStatus: constant.OrderStatusPaying,
			SellerID:    o.SellerId,
			BuyerID:     o.BuyerId,
			PayType:     req.PayType,
		})
		// items
		for _, s := range o.OrderItemList {
			itemRSp = append(itemRSp, &model.OrderItem{
				TradeNo:      tradeNo,
				OrderNo:      orderNo,
				SkuID:        s.SkuId,
				SpuID:        s.SpuId,
				SpuName:      s.SpuName,
				CategoryID:   s.CategoryId,
				CategoryName: s.CategoryName,
				BrandID:      s.BrandId,
				BrandName:    s.BrandName,
				ItemImgUrl:   s.SkuImgUrl,
				SpecValues:   s.SpecValues,
				SkuAmount:    s.SkuAmount,

				IsDelivered: false,
				IsConfirmed: false,
			})
		}
	}
	return trade, orderRsp, itemRSp, tradeNo
}
