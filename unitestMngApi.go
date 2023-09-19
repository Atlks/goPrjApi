package main

import (
	. "fmt"
	"goapiPrj/lib"
)

func main() {

	html := "1"
	Println(html)

	url := "http://localhost:8000/api?callfun=exit"
	http_get(url)
	//------------------------

	Println("\r\n\r\n")
	url = "http://localhost:8000/api?callfun=login%20111356,26916DD661300B25,1BC0036763DE22EC"

	html = http_get(url)

	Println("\r\n\r\n")
	url = "http://localhost:8000/api?callfun=qryAgtBal%20111356"

	html = http_get(url)
	Println("\r\n\r\n")
	//...
	url = "http://localhost:8000/api?callfun=searchPlayer%20777"
	html = http_get(url)
	Println("\r\n\r\n")
	url = "http://localhost:8000/api?callfun=shangfen%20777%202"
	html = http_get(url)
	Println("\r\n\r\n")
	url = "http://localhost:8000/api?callfun=xiafen%20777%201"
	http_get(url)

	//-------------------
	Println("\r\n\r\n")
	url = "http://localhost:8000/api?callfun=kick%20777"
	http_get(url)
	url = "http://localhost:8000/api?callfun=QryShangxiafen"
	http_get(url)

	url = "http://localhost:8000/api?callfun=kexiafenBal%20777"
	http_get(url)

	url = "http://localhost:8000/api?callfun=addUser%20111%201"
	http_get(url)

	url = "http://localhost:8000/api?callfun=oplog_qry"
	//--------------------10
	http_get(url)
	url = "http://localhost:8000/api?callfun=includeXAjax%20head.htm"
	http_get(url)

	url = "111356,26916DD661300B25,1BC0036763DE22EC"
}

func http_get(url string) string {
	return lib.Http_get(url)
}
