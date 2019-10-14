package lhttp

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Post 请求
func Post(postURL, entity string, headers g.Map) (*ghttp.ClientResponse, error) {
	c := ghttp.NewClient()

	for k, v := range headers {
		c.SetHeader(k, v.(string))
	}
	return c.Post(postURL, entity)
}
