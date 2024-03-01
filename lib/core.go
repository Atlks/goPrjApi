package lib

// cvt
func AnyToMap(v any) map[string]any {
	map1 := v.(map[string]any)
	return map1
}

func AnyToInt(v any) int32 {
	map1 := v.(int32)
	return map1
}

func AnyToNumFlt64(v any) float64 {
	map1 := v.(float64)
	return map1
}
