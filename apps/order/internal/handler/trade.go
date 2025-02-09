package handler

import (
	"context"
	"encoding/json"
	"time"

	"github.com/lzl-here/bt-shop-backend/apps/order/internal/common"
	"github.com/lzl-here/bt-shop-backend/apps/order/internal/constant"
	"github.com/lzl-here/bt-shop-backend/apps/order/internal/domain/model"
	bizerr "github.com/lzl-here/bt-shop-backend/pkg/err"
	"github.com/lzl-here/bt-shop-backend/pkg/utils"
	ggen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/goods"
	ogen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/order"
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

	return &ogen.CreateTradeRsp{
		Code: 1,
		Msg:  "success",
		Data: &ogen.CreateTradeRsp_CreateTradeRspData{
			TradeNo: tradeNo,
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

/*
* @description: 取消交易
**/
func (h *OrderHandler) CancelTrade(ctx context.Context, req *ogen.CancelTradeReq) (res *ogen.CancelTradeRsp, err error) {
	return nil, nil
}

/*
* @description: 获取交易详情
**/
func (h *OrderHandler) GetTradeDetail(ctx context.Context, req *ogen.GetTradeDetailReq) (*ogen.GetTradeDetailRsp, error) {
	var err error
	var trade *model.Trade
	trade, err = h.rep.GetTrade(ctx, &model.Trade{TradeNo: req.TradeNo})
	if err != nil {
		return nil, err
	}
	if trade == nil {
		return nil, bizerr.ErrResourceNotFound
	}
	// TODO 校验权限 if trade.BuyerID != 0 {
	// 	return nil, bizerr.ErrResourceNotFound
	// }
	var orderList []*model.Order
	if orderList, err = h.rep.GetOrderList(ctx, &model.Order{TradeNo: req.TradeNo}); err != nil {
		return nil, err
	}
	var orderItems []*model.OrderItem
	if orderItems, err = h.rep.GetOrderItems(ctx, &model.OrderItem{TradeNo: req.TradeNo}); err != nil {
		return nil, err
	}
	return buildTradeRsp(trade, orderList, orderItems)
}

func buildTradeRsp(trade *model.Trade, orderList []*model.Order, orderItems []*model.OrderItem) (*ogen.GetTradeDetailRsp, error) {
	baseTrade, err := trade.ToGen()
	if err != nil {
		return nil, err
	}
	orderItemMap := make(map[string][]*ogen.BaseOrderItem)
	for _, oi := range orderItems {
		baseOrderItem, err := oi.ToGen()
		if err != nil {
			return nil, err
		}
		orderItemMap[oi.OrderNo] = append(orderItemMap[oi.OrderNo], baseOrderItem)
	}
	orderRspList := make([]*ogen.TradeOrder, 0)
	for _, order := range orderList {
		baseOrder, err := order.ToGen()
		if err != nil {
			return nil, err
		}
		orderRspList = append(orderRspList, &ogen.TradeOrder{
			Order:         baseOrder,
			OrderItemList: orderItemMap[order.OrderNo],
		})
	}
	data := &ogen.GetTradeDetailRsp_GetTradeDetailRspData{
		Trade: &ogen.TradeTrade{
			Trade:     baseTrade,
			OrderList: orderRspList,
		},
	}
	rsp := &ogen.GetTradeDetailRsp{
		Code: 1,
		Msg:  "success",
		Data: data,
	}
	return rsp, nil
}

/**
* @description: 获取交易列表
 */
func (h *OrderHandler) GetTradeList(ctx context.Context, req *ogen.GetTradeListReq) (*ogen.GetTradeListRsp, error) {
	tradeList, err := h.rep.PageTradeList(ctx, &model.Trade{}, int(req.PageSize), int(req.PageNo))
	if err != nil {
		return nil, err
	}
	tradeNos := make([]string, 0)
	for _, t := range tradeList {
		tradeNos = append(tradeNos, t.TradeNo)
	}
	orderList, err := h.rep.GetOrderListByTradeNo(ctx, tradeNos)
	if err != nil {
		return nil, err
	}
	orderItems, err := h.rep.GetOrderItemsByTradeNo(ctx, tradeNos)
	if err != nil {
		return nil, err
	}
	orderMap := make(map[string][]*model.Order)
	for _, o := range orderList {
		orderMap[o.TradeNo] = append(orderMap[o.TradeNo], o)
	}
	orderItemsMap := make(map[string][]*model.OrderItem)
	for _, oi := range orderItems {
		orderItemsMap[oi.OrderNo] = append(orderItemsMap[oi.TradeNo], oi)
	}
	count, err := h.rep.CountTrade(ctx, &model.Trade{})
	if err != nil {
		return nil, err
	}
	return buildTradeListRsp(tradeList, orderMap, orderItemsMap, req.PageSize, req.PageNo, count)
}

func buildTradeListRsp(tradeList []*model.Trade, orderMap map[string][]*model.Order, orderItemsMap map[string][]*model.OrderItem, pageSize int32, pageNo int32, count int64) (*ogen.GetTradeListRsp, error) {
	tradeRspList := make([]*ogen.TradeTrade, 0)

	// trade
	for _, t := range tradeList {
		baseTrade, err := t.ToGen()
		if err != nil {
			return nil, err
		}
		// order
		orderRspList := make([]*ogen.TradeOrder, 0)
		orderList := orderMap[t.TradeNo]
		for _, o := range orderList {
			baseOrder, err := o.ToGen()
			if err != nil {
				return nil, err
			}
			// orderItem
			orderItemRspList := make([]*ogen.BaseOrderItem, 0)
			orderItems := orderItemsMap[o.OrderNo]
			for _, oi := range orderItems {
				baseOrderItem, err := oi.ToGen()
				if err != nil {
					return nil, err
				}
				orderItemRspList = append(orderItemRspList, baseOrderItem)
			}
			orderRspList = append(orderRspList, &ogen.TradeOrder{
				Order:         baseOrder,
				OrderItemList: orderItemRspList,
			})
		}
		tradeRspList = append(tradeRspList, &ogen.TradeTrade{
			Trade:     baseTrade,
			OrderList: orderRspList,
		})
	}
	rsp := &ogen.GetTradeListRsp{
		Code: 1,
		Msg:  "success",
		Data: &ogen.GetTradeListRsp_GetTradeListRspData{
			TradeList: tradeRspList,
			Count:     int32(count),
			PageSize:  pageSize,
			PageNo:    pageNo,
		},
	}
	return rsp, nil
}
