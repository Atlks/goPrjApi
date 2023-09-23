package lib

func Array_sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
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
