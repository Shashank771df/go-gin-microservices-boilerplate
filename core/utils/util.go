package utils

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"math"
	"reflect"
)

type Util struct {
}

func (obj Util) Round(value float64, decimals int) float64 {
	f := math.Pow10(decimals)
	k := int64(value * f)

	return math.Round(float64(k) / f)
}

func (obj Util) Tracev4() string {
	uuid2 := make([]byte, 16)
	var rander = rand.Reader

	_, err := io.ReadFull(rander, uuid2[:])
	if err != nil {
		return ""
	}
	uuid2[6] = (uuid2[6] & 0x0f) | 0x40 // Version 4
	uuid2[8] = (uuid2[8] & 0x3f) | 0x80 // Variant is 10

	return obj.encodeUUID(uuid2)
}

func (obj Util) encodeUUID(uuid []byte) string {
	dst := make([]byte, 36)

	hex.Encode(dst, uuid[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], uuid[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], uuid[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], uuid[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], uuid[10:])

	return string(dst)
}

func (obj Util) JsonScript(m *interface{}) *string {
	if m == nil {
		return nil
	}

	mJson, err := json.Marshal(m)
	if err != nil {
		return nil
	}

	jsonStr := string(mJson)
	return &jsonStr
}

func (obj Util) SetValue(dst interface{}, src interface{}) {
	dstValue := reflect.ValueOf(dst)
	srcValue := reflect.ValueOf(src)

	if dstValue.IsNil() || srcValue.IsNil() {
		return
	}

	dstValue.Elem().Set(srcValue.Elem())
}

func (obj Util) ToMoney(amount float64) uint64 {
	return uint64(amount * 10000)
}
