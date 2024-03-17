package utils

import (
	"encoding/json"
	"reflect"
)

type UtilsMap struct {
}

func (o UtilsMap) ErrorToStruct(in error, out interface{}) error {
	var err error

	if err = json.Unmarshal([]byte(in.Error()), &out); err != nil {
		return err
	}

	return nil
}

func (o UtilsMap) InterfaceToStruct(in interface{}, out interface{}) error {
	var err error
	var data []byte
	if data, err = json.Marshal(in); err != nil {
		return err
	}

	err = json.Unmarshal(data, &out)
	return err
}

func (o UtilsMap) GetStringValue(in map[string]interface{}, key string) string {
	if val, ok := in[key]; ok {
		if reflect.ValueOf(val).Kind() == reflect.String {
			return val.(string)
		}
	}

	return ""
}

func (o UtilsMap) GetMapValue(in map[string]interface{}, key string) map[string]interface{} {
	if val, ok := in[key]; ok {
		if reflect.ValueOf(val).Kind() == reflect.Map {
			return val.(map[string]interface{})
		}
	}

	return nil
}

func (o UtilsMap) MapToArray(in interface{}) []interface{} {
	ret := []interface{}{}

	data := reflect.ValueOf(in)
	for _, key := range data.MapKeys() {
		ret = append(ret, data.MapIndex(key).Interface())
	}

	return ret
}

func (o UtilsMap) MapToString(in interface{}) string {
	var err error
	var data []byte

	if data, err = json.Marshal(in); err != nil {
		return ""
	}

	return string(data)
}

func (o UtilsMap) StringToInterface(data string, out interface{}) error {
	var err error

	if err = json.Unmarshal([]byte(data), &out); err != nil {
		return err
	}

	return nil
}
