package utils

import (
	"bytes"
)

type ByteArray struct {
	Buffer []byte
}

func (obj *ByteArray) Init(len int) *ByteArray {
	obj.Buffer = make([]byte, len)
	return obj
}

func (obj *ByteArray) PushBackSingle(data byte) {
	obj.Buffer = append(obj.Buffer, data)
}

func (obj ByteArray) IsEmpty(data byte) bool {
	return obj.Len() == 0
}

func (obj *ByteArray) PushBack(data []byte) {
	obj.Buffer = append(obj.Buffer, data...)
}

func (obj *ByteArray) PushFrontSingle(data byte) {
	internal := []byte{byte(data)}
	obj.Buffer = append(internal, obj.Buffer...)
}

func (obj *ByteArray) PushFront(data []byte) {
	obj.Buffer = append(data, obj.Buffer...)
}

func (obj *ByteArray) PopFront(qty int) {
	if obj.Len() <= qty {
		return
	}

	obj.Buffer = obj.Buffer[qty:]
}

func (obj *ByteArray) PopBack(qty int) {
	if obj.Len() <= qty {
		return
	}

	obj.Buffer = obj.Buffer[:obj.Len()-qty]
}

func (obj ByteArray) ToHex() string {
	data := UtilString{}
	return data.HexFormat(obj.Buffer, " ")
}

func (obj ByteArray) Len() int {
	return len(obj.Buffer)
}

func (obj ByteArray) LastIndex() int {
	return len(obj.Buffer) - 1
}

func (obj ByteArray) At(pos int) byte {
	return obj.Buffer[pos]
}

func (obj ByteArray) Part(init int, end int) ByteArray {
	return ByteArray{Buffer: obj.Buffer}
}

func (obj *ByteArray) AddBack(value byte, qty int) {
	buff := bytes.Repeat([]byte{byte(value)}, qty)
	obj.Buffer = append(obj.Buffer, buff...)
}

func (obj *ByteArray) AddFront(value byte, qty int) {
	buff := bytes.Repeat([]byte{byte(value)}, qty)
	obj.Buffer = append(buff, obj.Buffer...)
}

func (obj *ByteArray) Reverse() {
	lenBuff := obj.Len()

	for i := 0; i < lenBuff/2; i++ {
		j := lenBuff - i - 1
		obj.Buffer[i], obj.Buffer[j] = obj.Buffer[j], obj.Buffer[i]
	}
}
