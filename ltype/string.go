package ltype

import (
	"bytes"
	"encoding/hex"
	"errors"
	"github.com/axgle/mahonia"
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

// EncodeToByte 中文字符转成 hex 字符串
func EncodeToByte(str string, encode string) (resp []byte) {
	if encode == "" {
		encode = "gbk"
	}

	enc := mahonia.NewEncoder(encode)
	resp = []byte(enc.ConvertString(str))
	return
}

// DecodeToByte hex 字符串转中文字节流
// n 字节长度, cdata 字节流, err 错误
func DecodeToByte(str string, decode string) (n int, cdata []byte, err error) {
	if decode == "" {
		decode = "gbk"
	}

	respByte, _ := hex.DecodeString(str)
	dec := mahonia.NewDecoder(decode)
	return dec.Translate(respByte, false)
}
