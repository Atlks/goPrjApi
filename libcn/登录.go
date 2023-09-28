package 类库包

import (
	"fmt"
)

// 登录.go
var 检查结果_凭据不存在 = "检查结果_没有凭据"
var 检查结果_凭据无效_用户名无效 = "凭据无效_用户名格式不对"

var 检查结果_凭据过期 = "凭据过期"
var 检查结果_凭据校验签名不通过 = "凭据校验签名不通过"
var 检查结果_凭据无效 = "凭据无效_缺少必要栏目"
var 检查结果_用户不存在 = "用户不存在"
var 检查结果_密码错误 = "密码错误"
var 检查结果_通过 = "pass"
var 检查结果 = ""

func I登录流程密码模式(凭据 map[string]string) {

	执行指令序列(
		//---------检查登录凭据(凭据) 按照一下规则;
		如果(I不存在(凭据), 提示并终止(检查结果_凭据不存在)),
		如果(I长度(凭据["用户名"]) < 3, 提示并终止(检查结果_凭据无效_用户名无效)),

		如果(用户不存在(凭据["用户名"]), 提示并终止(检查结果_用户不存在)),
		如果(核验密码错误(凭据["密码"]), 提示并终止(检查结果_密码错误)),

		// 检查通过，开始登陆
		// 如果是登录凭据是密码类型则发放登录凭据(凭据)
		发放登录凭据(凭据["用户名"]),
		//添加操作日志("用户:(@用户名@)  事件:登录，时间:@当前时间@ ip:@ip@"),
	)()

}

func 核验密码错误(密码 string) bool {
	//  return true;
	return false
}

//func I登录流程(凭据 map[string]string) {
//
//	检查结果 := 检查登录凭据(凭据)
//	如果(检查通过(检查结果),
//		如果是登录凭据是密码类型则发放登录凭据(凭据),
//		添加操作日志("用户 (@用户名@)登录，时间 @当前时间@"),
//	)
//	如果(检查不通过(检查结果), 提示并终止(检查结果))
//
//}

func 如果是登录凭据是密码类型则发放登录凭据(凭据 map[string]string) func(...any) any {
	return func(xx ...any) any {
		fmt.Println("...则发放登录凭据... ...")
		登录凭据 := map[string]string{
			"用户名": 凭据["用户名"],
			"过期日": "2025-1-1",
		}
		return 登录凭据

	}
}

func 检查登录凭据(凭据 map[string]string) string {
	如果(I不存在(凭据), 提示并终止(检查结果_凭据不存在))
	如果(I长度((凭据["用户名"])) < 3, 提示并终止(检查结果_凭据无效_用户名无效))

	如果(用户不存在(凭据["用户名"]), 提示并终止(检查结果_用户不存在))
	如果(凭据["类型"] == "密码", 检查密码(凭据))
	如果(凭据["类型"] == "token", 检查token(凭据))

	return 检查结果_通过
}

func 检查token(凭据 map[string]string) func(...any) any {
	return func(xx ...any) any {
		fmt.Println("...检查token 检查token... ...")
		return ""
	}
}

func 检查密码(凭据 map[string]string) func(...any) any {
	return func(xx ...any) any {
		fmt.Println("...检查密码... ...")
		return ""
	}
}

func 用户不存在(凭 interface{}) bool {
	return false
}

func 提示并终止(提示内容 string) func(...any) any {
	return func(xx ...any) any {
		fmt.Println("...提示并终止" + 提示内容 + "  ... ...")
		panic("终止异常@" + 提示内容)
		return ""
	}
}

//func I登录流程_用户名模式(用户名 string, 密码 string) {
//
//	检查结果 := 检查用户名密码(用户名, 密码)
//	如果(检查通过(检查结果),
//		发放登录凭据(用户名),
//		添加操作日志("用户 (@用户名@)登录，时间 @当前时间@"),
//	)
//
//	如果(检查不通过(检查结果), 提示并终止("用户名密码不对"))
//
//}

// funs(true, 终止流程, 终止流程)
func funs(b bool, fs ...func(...any)) {

}

func 添加操作日志(内容 string) func(...any) any {
	return func(x ...any) any {
		fmt.Println("...操作日志内容  " + 内容 + "... ...")
		return nil
	}
}

func 终止流程(x ...any) {
	I显示("终止流程....")
}

func 检查用户名密码(用户名 string, 密码 string) string {
	I显示("检查用户名密码...true...")
	return 检查结果_通过
}

func 检查不通过(检查结果 string) bool {
	return !检查通过(检查结果)
}
func 检查通过(检查结果 string) bool {
	if 检查结果 == 检查结果_通过 {
		return true
	}

	return false
}
func 如果(条件 bool, 指令集 ...func(...any) any) func() {

	return func() {
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

}

func 如果2(条件 bool, 指令集 ...func(...any) any) func() {

	return func() {
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

}

func 发放登录凭据(用户名 string) func() {
	return func() {
		fmt.Println("...发放登录凭据  " + 用户名 + "...ing...")

	}
}
