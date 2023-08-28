package es

import (
	"sync"

	"github.com/olivere/elastic/v7"

	"github.com/sjxiang/biz-demo/book-shop/pkg/conf"
)

// ES client
var (
	esOnce sync.Once
	esCli  *elastic.Client
)

// GetESClient get ES client
func GetESClient() *elastic.Client {
	if esCli != nil {
		return esCli
	}

	esOnce.Do(func() {
		cli, err := elastic.NewSimpleClient(elastic.SetURL(conf.ESAddress))
		if err != nil {
			panic("new es client failed, err=" + err.Error())
		}
		esCli = cli
	})

	return esCli
}


