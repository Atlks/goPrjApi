package main

import (
	"fmt"
	类库包 "goapiPrj/libcn"
	"net/http"
	"strings"
)

/*
*
http://localhost:80/api?call=reg%2011111
*/
func Handler1(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	callfun := values.Get("call")
	fmt.Fprintf(w, "callfun=>"+callfun)
	arr := strings.Split(callfun, " ")
	fun := arr[0]
	m := map[string]any{
		"注册流程":  类库包.I注册流程,
		"login": I登录流程,
		"reg":   类库包.I注册流程,
	}
	funx := m[fun]
	fmt.Println(funx)

}

func I登录流程(凭据 map[string]string) {

	//检查结果 := 检查登录凭据(凭据)
	//如果(检查通过(检查结果),
	//	如果是登录凭据是密码类型则发放登录凭据(凭据),
	//	添加操作日志("用户 (@用户名@)登录，时间 @当前时间@"),
	//)
	//如果(检查不通过(检查结果), 提示并终止(检查结果))

}

func main() {

	//	exec()

	http.HandleFunc("/", Handler1)
	http.HandleFunc("/api", Handler1)
	http.ListenAndServe(":80", nil)
}

func exec() {
	//datapath := "C:\\electron\\dist\\electron.exe   "
	//fmt.Println(datapath)
	//cmd := exec.Command(datapath, "C:\\modyfing\\jbbot\\zmng\\dsktp.js")
	//cmd.Run()
}

// C:\Users\attil\sdk\go1.21.0\bin\go.exe run C:\Users\attil\GolandProjects\awesomeProject\callfun_web.go
// go run xx.go

func 进入() {
	fmt.Println("进入")
	提示信息 := 123
	fmt.Println(提示信息)
	fmt.Println("进入")
}
func ff() {
	fmt.Println("hello22233")
}
