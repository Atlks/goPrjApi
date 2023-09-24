package main

import (
	. "fmt"
	"goapiPrj/lib"
	"testing"
)

func TestHandler1(t *testing.T) {

	html := "1"
	Println(html)

	url := "http://localhost:8000/callrmt?callfun=exit"
	http_get(url)
	//------------------------

	Println("")
	Println("")
	Println("")
	url = "http://localhost:8000/callrmt?callfun=login%20111356,26916DD661300B25,1BC0036763DE22EC"

	html = http_get(url)

	url = "http://localhost:8000/callrmt?callfun=qryAgtBal%20111356"

	html = lib.Http_getV2(url)

	//...
	url = "http://localhost:8000/callrmt?callfun=searchPlayer%20777"
	html = lib.Http_getV2(url)

	url = "http://localhost:8000/callrmt?callfun=shangfen%20777%202"
	html = lib.Http_getV2(url)

	url = "http://localhost:8000/callrmt?callfun=xiafen%20777%201"
	lib.Http_getV2(url)

	//-------------------

	url = "http://localhost:8000/callrmt?callfun=kick%20777"
	lib.Http_getV2(url)
	url = "http://localhost:8000/callrmt?callfun=QryShangxiafen"
	lib.Http_getV2(url)

	url = "http://localhost:8000/callrmt?callfun=kexiafenBal%20777"
	lib.Http_getV2(url)

	url = "http://localhost:8000/callrmt?callfun=addUser%20111%201"
	lib.Http_getV2(url)

	url = "http://localhost:8000/callrmt?callfun=oplog_qry"
	//--------------------10
	lib.Http_getV2(url)
	url = "http://localhost:8000/callrmt?callfun=includeXAjax%20head.htm"
	http_get(url)

	url = "111356,26916DD661300B25,1BC0036763DE22EC"
}

func http_get(url string) string {
	return lib.Http_get(url)
}
