package db

import (
	"context"
	"errors"

	"gorm.io/gorm/clause"
)


func IncrStock(ctx context.Context, productId, stockNum int64) error {
	return updateStock(ctx, productId, stockNum, "incr")
}

func DecrStock(ctx context.Context, productId, stockNum int64) error {
	return updateStock(ctx, productId, stockNum, "decr")
}

func updateStock(ctx context.Context, productId, stockNum int64, updateType string) error {
	products := make([]*Product, 0)

	tx := DB.Begin().WithContext(ctx)
	if tx.Error != nil {
		return tx.Error
	}

	// select for update
	if err := tx.Clauses(clause.Locking{Strength: "Update"}).Where("product_id = ?", productId).Find(&products).Error; err != nil {
		tx.Rollback()
		return err
	}
	if len(products) == 0 {
		tx.Rollback()
		return errors.New("item not found")
	}

	product := products[0]
	curStockNum := product.Stock

	if updateType == "incr" {
		curStockNum += stockNum
	} else if updateType == "decr" {
		curStockNum -= stockNum
	}
	
	if curStockNum < 0 {
		tx.Rollback()
		return errors.New("库存不足")
	}

	if err := tx.Model(&Product{}).Where("product_id = ?", productId).
		Updates(map[string]interface{}{
			"stock": curStockNum,
		}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
