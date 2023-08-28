package service

import (
	"context"

	"github.com/sjxiang/biz-demo/book-shop/internal/dal/db"
	"github.com/sjxiang/biz-demo/book-shop/internal/dal/es"
)


type ProductService struct {
	ctx context.Context
}

func NewProductService(ctx context.Context) *ProductService {
	return &ProductService{
		ctx: ctx,
	}
}

func (p *ProductService) GetProduct(ctx context.Context, productId int64) (*db.Product, error) {
	return db.GetProductById(ctx, productId)
}

func (p *ProductService) ListProducts(ctx context.Context, name, spuName *string, status *int64) ([]*db.Product, error) {
	filterParam := make(map[string]interface{})
	if name != nil {
		filterParam["name"] = *name
	}
	if spuName != nil {
		filterParam["spu_name"] = *spuName
	}
	if status != nil {
		filterParam["status"] = *status
	}

	products, err := db.ListProducts(ctx, filterParam)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductService) MGetProducts2C(ctx context.Context, productIds []int64) ([]*db.Product, error) {
	products, err := es.BatchGetProductById(ctx, productIds)
	return products, err
}

func (p *ProductService) Search(ctx context.Context, name, description, spuName *string) ([]*db.Product, error) {
	filterMap := make(map[string]interface{}, 0)
	if name != nil {
		filterMap["name"] = *name
	}
	if description != nil {
		filterMap["description"] = *description
	}
	if spuName != nil {
		filterMap["spu_name"] = *spuName
	}

	products, err := es.SearchProduct(ctx, filterMap)
	return products, err
}
