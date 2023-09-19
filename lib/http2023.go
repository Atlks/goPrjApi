package lib

import (
	. "fmt"
	"io/ioutil"
	"net/http"
)

func Http_get(url string) string {
	Println("\r\n" + url)
	Log_info("")
	Log_info("")
	Log_info("")
	Log_info("" + url)
	resp, _ := http.Get(url)

	defer resp.Body.Close()

	// 读取资源数据 body: []byte
	body, _ := ioutil.ReadAll(resp.Body)
	//io.Copy(os.Stdout, resp.Body)
	s := string(body)
	content := Left(s, 300)
	Println("[http get]ret=>" + content)
	//log.Println("[http get]ret=>" + content)
	Log_info("[http get]ret=>" + content)
	return s
}
