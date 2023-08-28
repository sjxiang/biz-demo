package service

import (
	"context"

	"github.com/sjxiang/biz-demo/book-shop/internal/dal/db"
)

func (p *ProductService) IncreaseStockNum(ctx context.Context, productId, incrNum int64) error {
	return db.IncrStock(ctx, productId, incrNum)
}

func (p *ProductService) DecreaseStockNum(ctx context.Context, productId, decrNum int64) error {
	return db.DecrStock(ctx, productId, decrNum)
}
