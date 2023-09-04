package utils

import (
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sjxiang/biz-demo/tiktok_demo/biz/mw/minio"
	"go.uber.org/zap"
)

// NewFileName Splicing user_id and time to make unique filename
func NewFileName(user_id, time int64) string {
	return fmt.Sprintf("%d.%d", user_id, time)
}

// URLconvert Convert the path in the database into a complete url accessible by the front end
func URLconvert(ctx context.Context, c *gin.Context, path string) (fullURL string) {
	if len(path) == 0 {
		return ""
	}
	arr := strings.Split(path, "/")

	u, err := minio.GetObjURL(ctx, arr[0], arr[1])
	if err != nil {
		zap.L().Info(err.Error())
		return ""
	}

	u.Scheme = string(c.Request.URL.Scheme)
	u.Host = string(c.Request.URL.Host)
	u.Path = "/src" + u.Path
	return u.String()
}
