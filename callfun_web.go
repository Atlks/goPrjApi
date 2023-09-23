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
		"login": 类库包.I登录流程,
		"reg":   类库包.I注册流程,
	}
	funx := m[fun]
	fmt.Println(funx)

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
