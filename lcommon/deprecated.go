package lcommon

import (
	"github.com/skiy/gfutils/lcrypto"
	"github.com/skiy/gfutils/lfile"
	"mime/multipart"
)

// MD5 MD5
// Deprecated
func MD5(str string) string {
	return lcrypto.MD5(str)
}

// FileMD5 FileMD5
// Deprecated
func FileMD5(file *multipart.FileHeader) (md5Str string, err error) {
	return lfile.MD5(file)
}
