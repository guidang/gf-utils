package lcommon

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"net/url"
)

// IsURL 是否为 URL
func IsURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// Domain 获得当前站点地址
func Domain(r *ghttp.Request, https bool) string {
	scheme := "http"
	if https || r.TLS != nil {
		scheme = "https"
	}
	return fmt.Sprintf(`%s://%s`, scheme, r.Host)
}
