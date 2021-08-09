package lcommon

import (
	"github.com/lesebox/gf-utils/lcrypto"
	"github.com/lesebox/gf-utils/lfile"
	"github.com/lesebox/gf-utils/lfilepath"
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

// Ext returns the real file name extension used by path.
// Deprecated
func Ext(path string) string {
	return lfilepath.Ext(path)
}
