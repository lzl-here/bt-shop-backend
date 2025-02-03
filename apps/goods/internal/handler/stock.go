package handler

import (
	"context"

	ggen "github.com/lzl-here/bt-shop-backend/kitex_gen/goods"
)

// TODO 扣减库存
func (h *GoodsHandler) StockReduce(ctx context.Context, req *ggen.StockReduceReq) (*ggen.StockReduceRsp, error) {

	return nil, nil
}

// TODO 扣减库存回滚
func (h *GoodsHandler) StockIncrease(ctx context.Context, req *ggen.StockIncreaseReq) (*ggen.StockIncreaseRsp, error) {
	return nil, nil
}
