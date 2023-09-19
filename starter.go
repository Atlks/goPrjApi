package main

import (
	"fmt"
	"net/http"
)

/*
*
http://localhost:8000/api?call=Log%2011111
*/
func Handler1(w http.ResponsgodoceWriter, r *http.Request) {
	values := r.URL.Query()
	callfun := values.Get("call")
	fmt.Fprintf(w, "callfun=>"+callfun)
}

//func main() {
//
//}

func main22() {

	//	exec()

	http.HandleFunc("/", Handler1)
	http.HandleFunc("/api", Handler1)
	http.ListenAndServe(":8000", nil)
}

func exec() {
	//datapath := "C:\\electron\\dist\\electron.exe   "
	//fmt.Println(datapath)
	//cmd := exec.Command(datapath, "C:\\modyfing\\jbbot\\zmng\\dsktp.js")
	//cmd.Run()
}

// C:\Users\attil\sdk\go1.21.0\bin\go.exe run C:\Users\attil\GolandProjects\awesomeProject\starter.go
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
