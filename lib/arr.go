package lib

import (
	"errors"
	"reflect"
	"strconv"
)

func Array_sum(nums []any) float64 {
	var sum float64 = 0
	for _, num := range nums {
		numCur, _ := toFloat64(num.Interface())
		sum += numCur
	}
	return sum
}

// found := inArray("apple", []interface{}{"banana", "orange", "apple"})
// fmt.Println(found) // 输出：true
func InArray(needle interface{}, haystack []interface{}) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func Array_column(structSlice []map[string]any, key string) []any {

	var sliceColumn []any
	for _, value := range structSlice {
		//fmt.Println("uid==>" + uidStr)
		sliceColumn = append(sliceColumn, value[key])
	}

	return sliceColumn
}

//func array_sum(arr []any) float64 {
//	return
//}

// 数组切片求和
func ArraySum(input interface{}) (sum float64, err error) {
	list := reflect.ValueOf(input)
	switch reflect.TypeOf(input).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < list.Len(); i++ {
			val := list.Index(i)
			v, err := toFloat64(val.Interface())
			if err != nil {
				return 0, err
			}
			sum += v
		}
	default:
		return 0, errors.New("input must be slice or array")
	}
	return
}

func toFloat64(in interface{}) (f64 float64, err error) {
	switch val := in.(type) {
	case float64:
		return val, nil
	case float32:
		return float64(val), nil
	case uint8:
		return float64(val), nil
	case uint16:
		return float64(val), nil
	case uint32:
		return float64(val), nil
	case uint64:
		return float64(val), nil
	case uint:
		if strconv.IntSize == 32 || strconv.IntSize == 64 {
			return float64(val), nil
		}
		return 0, errors.New("convert uint to float64 failed")
	case int8:
		return float64(val), nil
	case int16:
		return float64(val), nil
	case int32:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case int:
		if strconv.IntSize == 32 || strconv.IntSize == 64 {
			return float64(val), nil
		}
		return 0, errors.New("convert int to float64 failed")
	case bool:
		if val {
			f64 = 1
		}
		return
	case string:
		f64, err = strconv.ParseFloat(val, 64)
		if err == nil {
			return
		}
		return 0, errors.New("convert string to float64 failed")
	default:
		return 0, errors.New("convert to float64 failed")
	}
}
