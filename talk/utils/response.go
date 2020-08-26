package utils

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

//JSONResult json返回的标准格式
func JSONResult(status string, data interface{}, code int) map[string]interface{} {
	if datas, ok := data.(gin.H); ok {
		return gin.H{
			"message": status,
			"data":    datas,
			"code":    code,
		}
	}
	if data == nil {
		return gin.H{
			"message": status,
			"data":    gin.H(map[string]interface{}{}),
			"code":    code,
		}
	}
	maps := StructToMap(data)
	return gin.H{
		"message": status,
		"data":    gin.H(maps),
		"code":    code,
	}
}

//JSONResultArr 数组
func JSONResultArr(status string, data []interface{}, code int) map[string]interface{} {
	arrayData := StructArrToMap(data)
	return gin.H{
		"message": status,
		"data":    arrayData,
		"code":    code,
	}
}

//StructToMap 结构体转map
func StructToMap(data interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	if data == nil {
		return m
	}
	refValue := reflect.ValueOf(data)
	refType := reflect.TypeOf(data)
	for i := 0; i < refValue.NumField(); i++ {
		if _, isExsist := refType.Field(i).Tag.Lookup("protobuf"); !isExsist {
			continue
		}
		m[refType.Field(i).Name] = refValue.Field(i).Interface()
	}
	return m
}

//StructArrToMap 数组结构体转数组map
func StructArrToMap(data []interface{}) []interface{} {
	result := []interface{}{}
	m := map[string]interface{}{}
	for _, v := range data {
		m = StructToMap(v)
		result = append(result, m)
	}
	return result
}
