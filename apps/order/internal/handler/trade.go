package handler

import (
	"context"
	"encoding/json"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/lzl-here/bt-shop-backend/apps/order/internal/common"
	"github.com/lzl-here/bt-shop-backend/apps/order/internal/constant"
	"github.com/lzl-here/bt-shop-backend/apps/order/internal/domain/model"
	ggen "github.com/lzl-here/bt-shop-backend/kitex_gen/goods"
	ogen "github.com/lzl-here/bt-shop-backend/kitex_gen/order"
	pgen "github.com/lzl-here/bt-shop-backend/kitex_gen/pay"
	bizerr "github.com/lzl-here/bt-shop-backend/pkg/err"
	"github.com/lzl-here/bt-shop-backend/pkg/utils"
	"gorm.io/gorm"
)

// 1. 校验商品信息合法性
// 2. 创建交易、订单、订单项
// 3. 扣减库存
// 4. 拉起支付
// 5. 存储支付信息url
// 6. 返回支付页面给前端
func (h *OrderHandler) CreateTrade(ctx context.Context, req *ogen.CreateTradeReq) (res *ogen.CreateTradeRsp, err error) {
	// 幂等处理
	ok, cleaner, _, err := utils.NoDuplicate(ctx, h.rep.GetCache(), "order", req.Token, 30*time.Second)
	if !ok {
		return nil, bizerr.ErrDuplicateReq
	}
	if err != nil {
		cleaner()
		return nil, err
	}

	// 1. 校验商品信息
	if err := h.checkParam(ctx, req); err != nil {
		cleaner()
		return nil, err
	}
	// 2. 创建交易、订单、订单项
	var tradeNo string
	err = h.rep.GetDB().Transaction(func(tx *gorm.DB) error {
		var err error
		var trade *model.Trade
		var orders []*model.Order
		var orderItems []*model.OrderItem
		trade, orders, orderItems, tradeNo, err = buildModels(req)
		if err != nil {
			return err
		}
		// 创建交易
		err = h.rep.CreateTrade(ctx, h.rep.GetDB(), trade)
		if err != nil {
			return err
		}
		// 创建订单
		err = h.rep.CreateOrders(ctx, h.rep.GetDB(), orders)
		if err != nil {
			return err
		}
		// 创建订单项
		err = h.rep.CreateOrderItems(ctx, h.rep.GetDB(), orderItems)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	// 3, 扣减库存
	// TODO 扣减失败进行订单的状态修改
	items := make([]*ggen.StockReduceItem, 0)
	for _, o := range req.Trade.OrderList {
		for _, oi := range o.OrderItemList {
			items = append(items, &ggen.StockReduceItem{
				SkuId:  oi.SkuId,
				BuyNum: oi.BuyNum,
			})
		}
	}
	_, err = h.rep.ReduceStock(ctx, &ggen.StockReduceReq{
		TradeNo: tradeNo,
		Items:   items,
	})
	if err != nil {
		return nil, err
	}
	// 4. 拉起支付
	payRsp, err := h.rep.Pay(ctx, &pgen.PayReq{
		Subject:     "这是一个支付的标题",
		TotalAmount: req.Trade.Trade.TradeAmount,
		TradeNo:     tradeNo,
	})

	if err != nil {
		return nil, err
	}

	if payRsp.Code != 1 {
		return nil, kerrors.NewBizStatusError(payRsp.Code, payRsp.Msg)
	}

	// 5. 存储支付详情
	err = h.rep.CreateTradeInfo(ctx, h.rep.GetDB(), &model.TradeInfo{
		TradeNo: tradeNo,
		PageUrl: payRsp.Data.PayPageUrl,
	})
	if err != nil {
		return nil, err
	}
	// 6. 返回支付页面url
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
	spuIDs := make([]uint64, 0)
	for _, o := range req.Trade.OrderList {
		for _, s := range o.OrderItemList {
			spuIDs = append(spuIDs, s.SpuId)
		}
	}
	goodsRsp, err := h.rep.GetGoodsList(ctx, &ggen.GetGoodsListReq{
		SpuIds: spuIDs,
	})
	if err != nil {
		return nil
	}
	// 后端存的数据
	itemSkuIDMap := make(map[uint64]*ggen.BaseSku)
	for _, spu := range goodsRsp.Data.SpuList {
		for _, sku := range spu.SkuList {
			itemSkuIDMap[sku.Sku.Id] = sku.Sku
		}
	}
	// 前端传来的数据
	for _, o := range req.Trade.OrderList {
		for _, oi := range o.OrderItemList {
			// 不存在
			if _, ok := itemSkuIDMap[oi.SkuId]; !ok {
				return bizerr.ErrGoodsNotExist
			}
			// 金额有问题
			skuFromFront := itemSkuIDMap[oi.SkuId]
			if skuFromFront.SkuPrice != oi.SkuAmount {
				return bizerr.ErrGoodsNotCorrect
			}
		}
	}
	// 这里sku的金额没问题了，校验order和trade的金额
	// 订单金额 = sku金额之和
	// 交易金额 = 订单金额之和
	totalTradeAmount := "0"
	for _, o := range req.Trade.OrderList {
		totalOrderAmount := "0"
		for _, s := range o.OrderItemList {
			totalOrderAmount = common.AmountPlus(totalOrderAmount, s.SkuAmount)
		}
		if !common.AmountEqual(totalOrderAmount, o.Order.OrderAmount) {
			return bizerr.ErrInvalidTradeAmount
		}
		totalTradeAmount = common.AmountPlus(totalTradeAmount, o.Order.OrderAmount)
	}
	if !common.AmountEqual(totalTradeAmount, req.Trade.Trade.TradeAmount) {
		return bizerr.ErrInvalidTradeAmount
	}
	return nil
}

// 构建model
func buildModels(req *ogen.CreateTradeReq) (*model.Trade, []*model.Order, []*model.OrderItem, string, error) {
	tradeNo := common.GenTradeNo()
	orderRsp := make([]*model.Order, 0)
	itemRSp := make([]*model.OrderItem, 0)
	// trade
	trade := &model.Trade{
		TradeNo:     tradeNo,
		TradeAmount: req.Trade.Trade.TradeAmount,
		TradeStatus: constant.TradeStatusPaying,
		BuyerID:     req.Trade.Trade.BuyerId,
		PayType:     req.Trade.Trade.PayType,
	}
	reqOrders := req.Trade.OrderList
	// order
	for _, o := range reqOrders {
		orderNo := common.GenOrderNo()
		orderRsp = append(orderRsp, &model.Order{
			TradeNo:     tradeNo,
			OrderNo:     orderNo,
			OrderAmount: o.Order.OrderAmount,
			OrderStatus: constant.OrderStatusPaying,
			SellerID:    o.Order.SellerId,
			BuyerID:     o.Order.BuyerId,
			PayType:     req.Trade.Trade.PayType,
		})
		// items
		for _, s := range o.OrderItemList {
			bytes, err := json.Marshal(s.SpecValueList)
			if err != nil {
				return nil, nil, nil, "", err
			}
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
				SpecValues:   string(bytes),
				SkuAmount:    s.SkuAmount,

				IsDelivered: false,
				IsConfirmed: false,
			})
		}
	}
	return trade, orderRsp, itemRSp, tradeNo, nil
}

func (h *OrderHandler) CancelTrade(ctx context.Context, req *ogen.CancelTradeReq) (res *ogen.CancelTradeRsp, err error) {
	return nil, nil
}

/**
* @decription: 重新拉起支付 (返回url)
 */
func (h *OrderHandler) ReTrade(ctx context.Context, req *ogen.ReTradeReq) (*ogen.ReTradeRsp, error) {
	tradeInfo, err := h.rep.GetTradeInfo(ctx, &model.TradeInfo{TradeNo: req.TradeNo})
	if err != nil {
		return nil, err
	}
	return &ogen.ReTradeRsp{
		Code: 1,
		Msg:  "success",
		Data: &ogen.ReTradeRsp_ReTradeRspData{
			TradeInfo: &ogen.BaseTradeInfo{
				Id:      tradeInfo.ID,
				TradeNo: tradeInfo.TradeNo,
				PageUrl: tradeInfo.PageUrl,
			},
		},
	}, nil
}
