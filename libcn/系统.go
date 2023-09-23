package 类库包

import "fmt"

var I检查结果_通过 string = "pass"

func I显示(s any) {
	fmt.Println(s)
}

func I检查通过(检查结果 string) bool {
	if 检查结果 == I检查结果_通过 {
		return true
	}

	return false
}
func I检查不通过(检查结果 string) bool {
	return !检查通过(检查结果)
}
