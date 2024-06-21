package lib

import "log"

// 全局异常捕获函数
func HandlePanic() {
	if r := recover(); r != nil {
		log.Printf("Recovered from panic: %v", r)
	}
}
