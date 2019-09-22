package ltype

import "bytes"

// BytesToString []byte 去除空数据(0), (ASCII)再转成字符
func BytesToString(origin []byte) string {
	index := bytes.IndexByte(origin, 0)

	if index == -1 {
		return string(origin[0:])
	}

	return string(origin[0:index])
}
