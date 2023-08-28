package service

import (
	"context"
	"errors"
	"log"

	"github.com/sjxiang/biz-demo/book-shop/internal/dal/db"
	"github.com/sjxiang/biz-demo/book-shop/internal/dal/es"
	"github.com/sjxiang/biz-demo/book-shop/pkg/diff"
)

func (p *ProductService) AddProduct(ctx context.Context, product *db.Product) error {
	if product == nil {
		return errors.New("插入数据，不可为空")
	}
	
	// 异步 es 更新
	go func() {
		err := es.UpsertProductToES(ctx, product.ProductId, product)
		if err != nil {
			log.Printf("UpsertProductES err: %v", err)
		}
	}()
	
	return db.InsertProduct(ctx, product)
}

func (p *ProductService) EditProduct(ctx context.Context, origin, target *db.Product) error {
	productId := target.ProductId
	
	// update es async
	go func() {
		err := es.UpsertProductToES(ctx, productId, target)
		if err != nil {
			log.Printf("UpsertProductES err: %v", err)
		}
	}()

	// 值得研究
	changeMap := diff.GetChangedMap(origin, target)

	return db.UpdateProduct(ctx, productId, changeMap)
}