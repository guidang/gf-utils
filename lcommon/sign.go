package lcommon

import (
	"github.com/gogf/gf/frame/g"
	"net/url"
	"sort"
	"strings"
)

// SignEncode Sign Encode
func SignEncode(data, filter g.Map, passNull, urlencode bool) string {
	var paramKeys = make([]string, 0)

	temp := g.MapStrStr{}

	for k, v := range data {
		// 过滤
		if filter := filter[k]; filter != nil && filter.(bool) {
			continue
		}

		value := v.(string)

		// 过滤空值
		if passNull && value == "" {
			continue
		}

		// 编码
		if urlencode {
			value = url.QueryEscape(value)
		}

		temp[k] = value
		paramKeys = append(paramKeys, k)
	}

	var encryptArr []string // 加密串
	sort.Strings(paramKeys)
	for _, v := range paramKeys {
		encryptArr = append(encryptArr, v+"="+temp[v])
	}

	return strings.Join(encryptArr, "&")
}
