package lcommon

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
)

// MD5 字符串生成 MD5
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// FileMD5 文件 MD5
func FileMD5(file *multipart.FileHeader) (md5Str string, err error) {
	f, err := file.Open()
	if err != nil {
		return
	}
	defer f.Close()

	h := md5.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
