package handler

import (
	"context"
	"time"

	"github.com/lzl-here/bt-shop-backend/apps/goods/internal/constant"
	"github.com/lzl-here/bt-shop-backend/apps/goods/internal/domain/model"
	ggen "github.com/lzl-here/bt-shop-backend/protobuf/kitex_gen/goods"
	bizerr "github.com/lzl-here/bt-shop-backend/pkg/err"
	"github.com/lzl-here/bt-shop-backend/pkg/utils"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

// 扣减库存
// 生成一行扣减库存记录，trade_no和记录类型作为唯一索引来兜底
// 为了性能在加上redis set一个key来快速判断
// 和扣减库存在同一事务中，保证扣减库存的幂等性
// TODO 使用redis扣减库存后，db异步落库
func (h *GoodsHandler) StockReduce(ctx context.Context, req *ggen.StockReduceReq) (*ggen.StockReduceRsp, error) {
	ok, cleaner, _, _ := utils.NoDuplicate(ctx, h.rep.GetCache(), "stock_reduce", req.TradeNo, 30*time.Second)
	if !ok {
		return nil, bizerr.ErrDuplicateReq
	}

	err := h.stockOperate(ctx, req.TradeNo, req.Items, constant.StockOpTypeReduce)

	if err != nil {
		cleaner()
		return nil, err
	}

	return &ggen.StockReduceRsp{
		Code: 1,
		Msg:  "success",
		Data: &ggen.StockReduceRsp_StockReduceRspData{},
	}, nil
}

// 增加库存，和扣减库存逻辑一样
func (h *GoodsHandler) StockIncrease(ctx context.Context, req *ggen.StockIncreaseReq) (*ggen.StockIncreaseRsp, error) {
	ok, cleaner, _, _ := utils.NoDuplicate(ctx, h.rep.GetCache(), "stock_increase", req.TradeNo, 30*time.Second)
	if !ok {
		return nil, bizerr.ErrDuplicateReq
	}

	err := h.stockOperate(ctx, req.TradeNo, req.Items, constant.StockOpTypeIncrease)

	// 异常回滚
	if err != nil {
		cleaner()
		return nil, err
	}
	return nil, nil
}

// 操作库存
// 1. 生成操作记录保证强唯一性
// 2. 操作库存
func (h *GoodsHandler) stockOperate(ctx context.Context, tradeNo string, items []*ggen.StockReduceItem, op string) error {
	err := h.rep.Transaction(ctx, func(ctx context.Context, tx *gorm.DB) error {
		// 插入库存记录，opType + tradeNo 作为唯一索引，保证的强幂等性
		record := &model.StockOpRecord{
			TradeNo: tradeNo,
			OpType:  op,
		}
		if err := h.rep.CreateStockRecord(ctx, tx, record); err != nil {
			return err
		}
		// 并行操作库存
		eg, ctx := errgroup.WithContext(ctx)
		for _, i := range items {
			item := i // 需要重新赋值，避免闭包问题
			var f func() error

			if op == constant.StockOpTypeReduce {
				// 扣减库存的情况
				f = func() error {
					ok, err := h.rep.ReduceStock(ctx, tx, item.SkuId, item.BuyNum)
					if err != nil {
						return err
					}
					if !ok { //库存不足
						return bizerr.ErrStockNotEnough
					}
					return nil
				}
			} else if op == constant.StockOpTypeIncrease {
				// 释放库存的情况
				f = func() error {
					ok, err := h.rep.ReduceStock(ctx, tx, item.SkuId, item.BuyNum)
					if err != nil {
						return err
					}
					if !ok {
						return bizerr.ErrStockNotEnough
					}
					return nil
				}
			} else {
				return bizerr.ErrInternal
			}
			// 执行库存操作
			eg.Go(func() error {
				return f()
			})
		}
		return eg.Wait()
	})
	return err
}
