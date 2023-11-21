package main

import (
	"errors"
	"fmt"
	"goapiPrj/lib"
	"log"
	"reflect"
	"strconv"
)

func main() {

	fmt.Println("hello22233")
	f := "C:\\modyfing\\apiprj\\jbbot\\zmng\\db\\111523\\userColl.json"
	jsonRows := lib.ReadToJsonArr(f)
	idx := lib.RdmMax(len(jsonRows))

	fmt.Println(idx)
	idx = 128
	row := jsonRows[idx]

	uid := row["userid"]

	uidStr := fmt.Sprintf("%.0f", uid)

	fmt.Println("uid==>" + uidStr)
	rzt := sumAllbet(uidStr)
	fmt.Println("sumbet==>" + fmt.Sprintf("%.0f", rzt))
	lib.WriteJsonArr("ucol.json", jsonRows)

}

func sumAllbet(userid string) float64 {
	file := "C:\\modyfing\\apiprj\\jbbot\\zmng\\db_zhudan\\zhudan_uid" + userid + ".json"
	data_rows := lib.ReadToJsonArr(file)

	modifiedArr := array_column(data_rows, "ValidBet")

	arraySum, err := ArraySum(modifiedArr)
	if err != nil {
		log.Fatal("Error ==>", err)
	}
	return arraySum

}

func array_column(structSlice []map[string]any, key string) []any {
	rt := reflect.TypeOf(structSlice)
	rv := reflect.ValueOf(structSlice)
	if rt.Kind() == reflect.Slice { //切片类型
		var sliceColumn []interface{}
		elemt := rt.Elem() //获取切片元素类型
		for i := 0; i < rv.Len(); i++ {
			inxv := rv.Index(i)
			if elemt.Kind() == reflect.Struct {
				for i := 0; i < elemt.NumField(); i++ {
					if elemt.Field(i).Name == key {
						strf := inxv.Field(i)
						switch strf.Kind() {
						case reflect.String:
							sliceColumn = append(sliceColumn, strf.String())
						case reflect.Float64:
							sliceColumn = append(sliceColumn, strf.Float())
						case reflect.Int, reflect.Int64:
							sliceColumn = append(sliceColumn, strf.Int())
						default:
							//do nothing
						}
					}
				}
			}
		}
		return sliceColumn
	}
	return nil
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
