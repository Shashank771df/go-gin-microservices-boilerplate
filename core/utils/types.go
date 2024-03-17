package utils

type Types struct {
}

func (o Types) String(v string) *string {
	return &v
}

func (o Types) Bool(v bool) *bool {
	return &v
}

func (o Types) Int(v int) *int {
	return &v
}

func (o Types) Uint(v uint) *uint {
	return &v
}

func (o Types) Uint8(v uint8) *uint8 {
	return &v
}

func (o Types) Uint16(v uint16) *uint16 {
	return &v
}

func (o Types) Uint32(v uint32) *uint32 {
	return &v
}

func (o Types) Uint64(v uint64) *uint64 {
	return &v
}

func (o Types) Int8(v int8) *int8 {
	return &v
}

func (o Types) Int16(v int16) *int16 {
	return &v
}

func (o Types) Int32(v int32) *int32 {
	return &v
}

func (o Types) Int64(v int64) *int64 {
	return &v
}

func (o Types) Float32(v float32) *float32 {
	return &v
}

func (o Types) Float64(v float64) *float64 {
	return &v
}
