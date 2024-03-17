package utils

import (
	"encoding/hex"
	"strconv"
	"strings"
)

type UtilString struct {
	Value string
}

func (obj UtilString) ToInt64(ok *bool) int64 {
	var err error
	var value int64

	*ok = true
	if value, err = strconv.ParseInt(obj.Value, 10, 64); err != nil {
		*ok = false
		return 0
	}

	return value
}

func (obj UtilString) ToUint64(ok *bool) uint64 {
	var err error
	var value uint64

	*ok = true
	if value, err = strconv.ParseUint(obj.Value, 10, 64); err != nil {
		*ok = false
		return 0
	}

	return value
}

func (obj UtilString) ToInt32(ok *bool) int32 {
	var err error
	var value int64

	*ok = true
	if value, err = strconv.ParseInt(obj.Value, 10, 32); err != nil {
		*ok = false
		return 0
	}

	return int32(value)
}

func (obj UtilString) ToUint32(ok *bool) uint32 {
	var err error
	var value uint64

	*ok = true
	if value, err = strconv.ParseUint(obj.Value, 10, 32); err != nil {
		*ok = false
		return 0
	}

	return uint32(value)
}

func (obj UtilString) ToBool(ok *bool) bool {
	var err error
	var value bool

	*ok = true
	if value, err = strconv.ParseBool(obj.Value); err != nil {
		*ok = false
		return false
	}

	return value
}

func (obj UtilString) ToFloat(ok *bool) float64 {
	var err error
	var value float64

	*ok = true
	if value, err = strconv.ParseFloat(obj.Value, 64); err != nil {
		*ok = false
		return 0.0
	}

	return value
}

func (obj *UtilString) FromInt(value int64) string {
	obj.Value = strconv.FormatInt(value, 10)

	return obj.Value
}

func (obj UtilString) FromUint(value uint64) string {
	obj.Value = strconv.FormatUint(value, 10)

	return obj.Value
}

func (obj *UtilString) FromInt32(value int32) string {
	obj.Value = strconv.FormatInt(int64(value), 10)

	return obj.Value
}

func (obj UtilString) FromUint16(value uint16) string {
	obj.Value = strconv.FormatUint(uint64(value), 10)

	return obj.Value
}

func (obj *UtilString) FromInt16(value int16) string {
	obj.Value = strconv.FormatInt(int64(value), 10)

	return obj.Value
}

func (obj UtilString) FromUint8(value uint8) string {
	obj.Value = strconv.FormatUint(uint64(value), 10)

	return obj.Value
}

func (obj *UtilString) FromInt8(value int8) string {
	obj.Value = strconv.FormatInt(int64(value), 10)

	return obj.Value
}

func (obj UtilString) FromUint32(value uint32) string {
	obj.Value = strconv.FormatUint(uint64(value), 10)

	return obj.Value
}

func (obj UtilString) FormatBool(value bool) string {
	obj.Value = strconv.FormatBool(value)

	return obj.Value
}

func (obj UtilString) FormatFloat(value float64, precision int) string {
	obj.Value = strconv.FormatFloat(value, 'f', precision, 64)

	return obj.Value
}

func (obj *UtilString) ToLower() string {
	obj.Value = strings.ToLower(obj.Value)

	return obj.Value
}

func (obj *UtilString) ToUpper() string {
	obj.Value = strings.ToUpper(obj.Value)

	return obj.Value
}

func (obj *UtilString) HexFormat(buffer []byte, separator string) string {
	obj.Value = hex.EncodeToString(buffer)
	digits := len(obj.ToUpper())

	var ret string = ""

	for i := 0; i < digits; i += 2 {
		ret += obj.Value[i : i+2]
		if i+2 < digits {
			ret += separator
		}
	}

	obj.Value = ret

	return ret
}

func (obj UtilString) StartsWith(value string, prefix string) bool {
	return strings.HasPrefix(value, prefix)
}

func (obj UtilString) Split(value string, sep string) []string {
	return strings.Split(value, sep)
}

func (obj UtilString) Replace(value string, char string) string {
	return strings.ReplaceAll(value, char, "")
}
