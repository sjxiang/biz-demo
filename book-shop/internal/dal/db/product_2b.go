package db

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	
	ProductId   int64  `json:"product_id"`
	Name        string `json:"name"`
	Pic         string `json:"pic"`
	Description string `json:"description"`
	ISBN        string `json:"isbn"`
	SpuName     string `json:"spu_name"`
	SpuPrice    int64  `json:"spu_price"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
	Status      int64  `json:"status"`
}

func (p *Product) TableName() string {
	return "t_product"
}


// 2B 
// 添加商品
// 编辑商品
// 删除商品
// 上架商品
// 下架商品
// 查询商品 
// 商品列表 

// 2C
// 批量查询商品 
// 搜索商品 

// 扣减库存
// 库存返还


func GetProductById(ctx context.Context, productId int64) (*Product, error) {
	products := make([]*Product, 0)
	err := DB.WithContext(ctx).Where("product_id = ?", productId).Find(&products).Error
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, errors.New("该商品不存在")
	}

	product := products[0]
	
	return product, nil
}

func ListProducts(ctx context.Context, filterParam map[string]interface{}) ([]*Product, error) {
	products := make([]*Product, 0)
	
	DB = DB.Debug().WithContext(ctx)
	for k, v := range filterParam {
		DB = DB.Where(k+" = ?", v)
	}
	if err := DB.Find(&products).Error; err != nil {
		return nil, err
	}
	
	return products, nil
}


func InsertProduct(ctx context.Context, product *Product) error {
	return DB.WithContext(ctx).Create(product).Error
}

func UpdateProduct(ctx context.Context, productId int64, changeMap map[string]interface{}) error {
	return DB.WithContext(ctx).Model(&Product{}).Where("product_id = ?", productId).
		Updates(changeMap).Error
}