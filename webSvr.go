package main

import (
	"fmt"
	"goapiPrj/lib"
	"net/http"
	"strings"
)

/*
*
http://localhost:80/api?call=reg%2011111
*/
func Handler2(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	callfun := values.Get("call")
	fmt.Fprintf(w, "callfun=>"+callfun)
	arr := strings.Split(callfun, " ")
	fun := arr[0]
	m := map[string]any{}
	funx := m[fun]
	fmt.Println(funx)
	fmt.Println(555)
	fmt.Fprintf(w, "555")
}

func main() {

	//	exec()
	lib.MainPain()
	//	类库.M2()
	//FF()

	http.HandleFunc("/", Handler2)
	http.HandleFunc("/api", Handler2)
	http.ListenAndServe(":80", nil)
}
