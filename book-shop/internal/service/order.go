package service

import (
	"context"
	"strconv"

	"github.com/sjxiang/biz-demo/book-shop/internal/dal/db"
	"github.com/sjxiang/biz-demo/book-shop/internal/types"
)

type OrderService struct {
	ctx        context.Context
	productSvc *ProductService
}

func NewOrderService(ctx context.Context, productSvc *ProductService) *OrderService {
	return &OrderService{
		ctx:        ctx,
		productSvc: productSvc,
	}
}

func (o OrderService) ListOrder(userId int64, status *int64) ([]*db.Order, error) {
	filter := make(map[string]interface{})
	filter["user_id"] = userId
	if status != nil {
		filter["status"] = *status
	}
	res, err := db.ListOrders(o.ctx, filter)
	return res, err
}

func (o OrderService) GetOrderById(orderId int64) (*db.Order, error) {
	po, err := db.GetOrderById(o.ctx, orderId)
	return po, err
}


func (o OrderService) CreateOrder(req *types.CreateOrderReq) error {
	productId, _ := strconv.ParseInt(req.ProductId, 10, 64)
	
	// 扣减库存
	err := o.productSvc.DecreaseStockNum(o.ctx, productId, req.StockNum)
	if err != nil {
		return err
	}
	
	list := make([]*db.Order, 0)
	list = append(list, &db.Order{ProductId: productId})  // 象征性写点数据

	// 插入数据
	err = db.InsertOrder(o.ctx, list)
	if err != nil {
		// 回滚
		o.createRollback(req)
		return err
	}

	return nil
}

func (o OrderService) createRollback(req *types.CreateOrderReq) {
	productId, _ := strconv.ParseInt(req.ProductId, 10, 64)
	
	// 库存返还
	_ = o.productSvc.IncreaseStockNum(o.ctx, productId, req.StockNum)
}


func (o OrderService) CancelOrder(req *types.CancelOrderReq) error {
	updateMap := map[string]interface{}{
		"status": int64(1),  // 订单取消
	}
	orderId, _ := strconv.ParseInt(req.OrderId, 10, 64)
	order, err := db.GetOrderById(o.ctx, orderId)
	if err != nil {
		return err
	}

	// 库存返还
	err = o.productSvc.IncreaseStockNum(o.ctx, order.ProductId, order.StockNum)
	if err != nil {
		return err
	}
	
	// 修改状态
	err = db.UpdateOrder(o.ctx, orderId, updateMap)
	if err != nil {
		o.cancelRollback(order.ProductId, order.StockNum)
		return err
	}
	return nil

}

func (o OrderService) cancelRollback(productId, stockNum int64) {
	// 库存返还
	_ = o.productSvc.DecreaseStockNum(o.ctx, productId, stockNum)
}