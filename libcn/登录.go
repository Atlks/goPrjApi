package 类库包

import "fmt"

func I登录流程(用户名 string, 密码 string) {

	检查结果 := 检查用户名密码(用户名, 密码)
	如果(检查通过(检查结果),
		发放登录凭据(用户名),
		添加操作日志("用户 (@用户名@)登录，时间 @当前时间@"),
	)

	如果(检查不通过(检查结果), 提示("用户名密码不对"), 终止流程)

}

// funs(true, 终止流程, 终止流程)
func funs(b bool, fs ...func(...any)) {

}

func 添加操作日志(内容 string) func(...any) {
	return func(xx ...any) {
		fmt.Println("...操作日志内容  " + 内容 + "... ...")
	}
}

func 终止流程(x ...any) {
	I显示("终止流程....")
}

func 检查用户名密码(用户名 string, 密码 string) bool {
	I显示("检查用户名密码...true...")
	return true
}

func 检查不通过(检查结果 bool) bool {
	return !检查结果
}
func 检查通过(检查结果 bool) bool {
	return 检查结果
}
func 如果(条件 bool, 指令集 ...func(...any)) {
	if 条件 {
		for index, value := range 指令集 {

			fmt.Println(index)
			value()
			//var fun func()
			//fun : = value
			//fmt.Println("Index =", index, "Value =", value)
		}
	}
}
func 提示(内容2 any) func(...any) {
	return func(内容 ...any) {
		fmt.Println(内容2)
	}

}
func 发放登录凭据(用户名 string) func(...any) {
	return func(内容 ...any) {
		fmt.Println("...发放登录凭据  " + 用户名 + "...ing...")
	}
}
