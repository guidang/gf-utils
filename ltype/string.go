package ltype

import (
	"bytes"
	"encoding/hex"
	"errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

// HexToBytes 字符串转Bytes
func HexToBytes(str string) (resp []byte, err error) {
	return hex.DecodeString(str)
}

// StringToByte 字符串转Bytes
func StringToByte(str string) (resp []byte, err error) {
	strLen := len(str)
	if strLen == 0 {
		return nil, errors.New("no data")
	}

	for i := 0; i < strLen; i++ {
		resp = append(resp, str[i]&0xFF)
	}
	return resp, nil
}

// UTF8ToGBK UTF8ToGBK
func UTF8ToGBK(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
