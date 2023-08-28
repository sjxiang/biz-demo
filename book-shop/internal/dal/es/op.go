package es

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/olivere/elastic/v7"
	"github.com/sjxiang/biz-demo/book-shop/internal/dal/db"
)

// 插入 doc
func UpsertProductToES(ctx context.Context, productId int64, product *db.Product) error {
	// 拷贝
	doc := getDocFromPo(product)
	
	_, err := GetESClient().Update().
		Index("product").
		Id(strconv.FormatInt(productId, 10)).
		Doc(doc).
		Upsert(doc).
		Refresh("true").
		Do(ctx)

	return err
}


// 批量获取 doc 
func BatchGetProductById(ctx context.Context, productIds []int64) ([]*db.Product, error) {
	mgetSvc := GetESClient().MultiGet()

	for _, id := range productIds {
		mgetSvc.Add(elastic.NewMultiGetItem().
			Index("product").
			Id(strconv.FormatInt(id, 10)))
	}

	rsp, err := mgetSvc.Do(ctx)
	if err != nil {
		return nil, err
	}
	products := make([]*db.Product, 0)
	for _, doc := range rsp.Docs {
		products = append(products, getPoFromSource(string(doc.Source)))
	}
	return products, nil
}

// 查询 doc 
func SearchProduct(ctx context.Context, filter map[string]interface{}) ([]*db.Product, error) {
	boolQuery := elastic.NewBoolQuery()
	for k, v := range filter {
		boolQuery.Must(elastic.NewMatchQuery(k, v))
	}
	searchResult, err := GetESClient().Search().
		Index("product").
		Query(boolQuery).
		Size(10000).
		From(0).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	if searchResult.Hits.TotalHits.Value <= 0 || len(searchResult.Hits.Hits) <= 0 {
		return []*db.Product{}, nil
	}

	ret := make([]*db.Product, 0)
	for _, hit := range searchResult.Hits.Hits {
		ret = append(ret, getPoFromSource(string(hit.Source)))
	}

	return ret, nil
}

// convertToPo
func getPoFromSource(source string) *db.Product {

	sourceMap := make(map[string]interface{}, 0)
	_ = json.Unmarshal([]byte(source), &sourceMap)
	
	ret := &db.Product{
		ProductId:   int64(sourceMap["product_id"].(float64)),
		Name:        sourceMap["name"].(string),
		Pic:         sourceMap["pic"].(string),
		Description: sourceMap["description"].(string),	
		ISBN:        sourceMap["isbn"].(string),
		SpuName:     sourceMap["spu_name"].(string),
		SpuPrice:    int64(sourceMap["spu_price"].(float64)),
		Price:       int64(sourceMap["price"].(float64)),
		Stock:       int64(sourceMap["stock"].(float64)),
		Status:      int64(sourceMap["status"].(float64)),
	}
	
	return ret
}

// convertToDoc
func getDocFromPo(e *db.Product) map[string]interface{} {
	ret := map[string]interface{}{
		"product_id":  e.ProductId,
		"name":        e.Name,
		"pic":         e.Pic,
		"description": e.Description,
		"price":       e.Price,
		"stock":       e.Stock,
		"status":      e.Status,
		"isbn":        e.ISBN,
		"spu_name":    e.SpuName,
		"spu_price":   e.SpuPrice,
	}

	return ret
}
