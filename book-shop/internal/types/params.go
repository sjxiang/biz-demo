package types


type UserParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type LoginResponse struct {
	Code   int64  `json:"code"`
	Expire string `json:"expire"`
	Token  string `json:"token"`
}

type AddProductRequest struct {
	Name        string `json:"name"`
	Pic         string `json:"pic"`
	Description string `json:"description"`
	ISBN        string `json:"isbn"`
	SpuName     string `json:"spu_name"`
	SpuPrice    int64  `json:"spu_price"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
}

type EditProductRequest struct {
	ProductId   string  `json:"product_id"`
	Name        string `json:"name"`
	Pic         string `json:"pic"`
	Description string `json:"description"`
	ISBN        string `json:"isbn"`
	SpuName     string `json:"spu_name"`
	SpuPrice    int64  `json:"spu_price"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
}

type OperateProductReq struct {
	ProductId string `json:"product_id"`
}

type SearchProductReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	SpuName     string `json:"spu_name"`
}

type ListProductReq struct {
	Name    string `json:"name"`
	SpuName string `json:"spu_name"`
	Status  int64  `json:"status"`
}

type CreateOrderReq struct {
	Address   string `json:"address"`
	ProductId string `json:"product_id"`
	StockNum  int64  `json:"stock_num"`
}

type CancelOrderReq struct {
	OrderId string `json:"order_id"`
}

type ListOrderReq struct {
	Status int64 `json:"status"`
}
