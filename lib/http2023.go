package lib

import (
	. "fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Http_getV2(url string) string {
	Println("\r\n" + url)
	Log_info("")
	Log_info("")
	Log_info("")
	Log_info("" + url)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Cookie", "agtid=111356; desCode=26916DD661300B25; md5Code=1BC0036763DE22EC")
	//resp, _ := http.Get(url)

	client := &http.Client{Timeout: time.Second * 10}
	resp, _ := client.Do(req)

	//http.Header.Set()
	//	http.Header.Set("Cookie", "agtid=111356; desCode=26916DD661300B25; md5Code=1BC0036763DE22EC")
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

func Http_get(url string) string {
	Println("\r\n" + url)
	Log_info("")
	Log_info("")
	Log_info("")
	Log_info("" + url)
	resp, _ := http.Get(url)

	//http.Header.Set()
	//	http.Header.Set("Cookie", "agtid=111356; desCode=26916DD661300B25; md5Code=1BC0036763DE22EC")
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
