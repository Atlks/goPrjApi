package lib

import (
	. "fmt"
	"io/ioutil"
	"net/http"
)

func http_get(url string) string {
	Println("\r\n" + url)
	resp, _ := http.Get(url)

	defer resp.Body.Close()

	// 读取资源数据 body: []byte
	body, _ := ioutil.ReadAll(resp.Body)
	//io.Copy(os.Stdout, resp.Body)
	s := string(body)
	Println("[http get]ret=>" + s)
	return s
}
