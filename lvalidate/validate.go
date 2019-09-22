package mvalidate

import (
	"bytes"
	"errors"
)

// Validate Validate
type Validate struct {
	data,
	code []byte
	flag Mode
}

// NewValidate Validate init
func NewValidate(data []byte, code []byte, flag int64) *Validate {
	t := &Validate{}
	t.data = data // 数据
	t.code = code // 校验码
	t.flag = flag // 校验码生成方式
	return t
}

// ValidateCode 校验码检测一致性
func (t *Validate) ValidateCode() error {
	if t.code == nil {
		return errors.New("校验码不存在")
	}

	oldCode := t.code
	if err := t.MakeCode(); err != nil {
		return errors.New("生成校验码失败")
	}

	if bytes.Equal(oldCode, t.code) {
		return nil
	}

	return errors.New("校验码不匹配")
}

// MakeCode 生成校验码
func (t *Validate) MakeCode() (err error) {
	if ByteOverlayAccumulation == t.flag {
		return t.makeCodeFlag1()
	}

	return errors.New("不支持此校验码生成方式")
}

// makeCodeFlag1 字节流累计相加方式
func (t *Validate) makeCodeFlag1() (err error) {
	dataLen := len(t.data)

	if dataLen == 0 {
		return errors.New("空数据无法生成校验码")
	}

	code := t.data[0]
	for i := 1; i < dataLen; i++ {
		code ^= t.data[i]
	}

	t.code = append([]byte{}, code)
	return
}

// GetCode 获取校验码
func (t *Validate) GetCode() []byte {
	return t.code
}
