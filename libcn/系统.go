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

func console_log(x any) {
	fmt.Println(x)
}

func 循环(操作序列, 单条操作指令 func()) {
	console_log(操作序列)

}

func 执行指令序列(操作序列 ...func()) func() {
	return func() {
		for index, 操作 := range 操作序列 {
			fmt.Println(index)
			操作()
		}

	}
}

func I不存在(凭据 map[string]string) bool {
	if 凭据 != nil {
		return true
	} else {
		return false
	}

}
