package main

import (
	. "fmt"
	ss "goapiPrj/lib"
	"io/ioutil"
	"net/http"
)

func main() {

	html := "1"
	Println(html)

	Println("\r\n\r\n")
	url := "http://localhost:8000/api?callfun=login%20111356,26916DD661300B25,1BC0036763DE22EC"

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
	Println("\r\n\r\n")
	url = "http://localhost:8000/api?callfun=kick%20777"
	http_get(url)
	url = "http://localhost:8000/api?callfun=QryShangxiafen"
	http_get(url)
}

func http_get(url string) string {
	Println("\r\n" + url)
	resp, _ := http.Get(url)

	defer resp.Body.Close()

	// 读取资源数据 body: []byte
	body, _ := ioutil.ReadAll(resp.Body)
	//io.Copy(os.Stdout, resp.Body)
	s := string(body)
	content := ss.Left(s, 300)
	Println("[http get]ret=>" + content)
	return s
}
