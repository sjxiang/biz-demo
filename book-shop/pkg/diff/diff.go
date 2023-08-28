package diff

import (
	"github.com/r3labs/diff/v2"

	"github.com/sjxiang/biz-demo/book-shop/internal/dal/db"
)

func GetChangedMap(origin, target *db.Product) map[string]interface{} {
	d, _ := diff.NewDiffer(diff.TagName("json"))
	
	changedMap := make(map[string]interface{})
	changeLog, _ := d.Diff(origin, target)
	
	for _, change := range changeLog {
		if depth := len(change.Path); depth != 1 {
			continue
		}
		if change.Type == diff.UPDATE {
			changedMap[change.Path[0]] = change.To
		}
	}
	return changedMap
}
