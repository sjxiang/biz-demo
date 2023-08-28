package db

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderId         int64  `json:"order_id"`
	UserId          int64  `json:"user_id"`
	Address         string `json:"address"`
	ProductId       int64  `json:"product_id"`
	StockNum        int64  `json:"stock_num"`
	ProductSnapshot string `json:"product_snapshot"`
	Status          int64  `json:"status"`
}

func (o *Order) TableName() string {
	return "t_order"
}


func InsertOrder(ctx context.Context, orders []*Order) error {
	return DB.WithContext(ctx).Create(orders).Error
}

func UpdateOrder(ctx context.Context, orderId int64, updateMap map[string]interface{}) error {
	return DB.WithContext(ctx).Model(&Order{}).Where("order_id = ?", orderId).
		Updates(updateMap).Error
}

func ListOrders(ctx context.Context, filterMap map[string]interface{}) ([]*Order, error) {
	res := make([]*Order, 0)
	db := DB.WithContext(ctx)
	for k, v := range filterMap {
		db = db.Where(k+" = ?", v)
	}
	err := db.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetOrderById(ctx context.Context, orderId int64) (*Order, error) {
	res := make([]*Order, 0)
	err := DB.WithContext(ctx).Where("order_id = ?", orderId).Find(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if len(res) == 0 {
		return nil, errors.New("不存在该订单")
	}
	return res[0], nil
}