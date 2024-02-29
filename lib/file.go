package lib

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ReadToStr(f string) string {

	content, _ := ioutil.ReadFile(f)
	conStr := string(content)
	return conStr
}
func ReadToJsonArr(f string) []map[string]any {
	content, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal("Error when opening file==>", err)
	}

	// Now let's unmarshall the data into `payload`
	//var payload map[string]any
	//err = json.Unmarshal(content, &payload)
	//if err != nil {
	//	log.Fatal("Error during Unmarshal()==>>", err)
	//}

	//定义一个slice
	var slice1 []map[string]any
	//注意：反序列化map,不需要make，因为make操作被封装到Unmarshal函数
	err2 := json.Unmarshal([]byte(content), &slice1)
	if err2 != nil {
		fmt.Printf("unmarshal err==>%v\n", err2)
		log.Fatal("Error during Unmarshal()==>>", err2)
	}

	return slice1

}

func ReadLine(f string, lineHdlr func(line string)) {
	readFile, _ := os.Open(f)

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lineHdlr(fileScanner.Text())
	}
}

func Write(f string, data string) {

	file, err := os.OpenFile(f, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.WriteString(data); err != nil {
		panic(err)
	}
}

func WriteJsonArr(f string, data3 []map[string]any) {

	// 有个坑，Python、Java的写文件默认函数操作默认是覆盖的，而是Golang的OpenFile函数写入默认是追加的
	// os.O_TRUNC 覆盖写入，不加则追加写入
	file, err := os.OpenFile(f, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := Json_encode(data3)

	if _, err := file.WriteString(data); err != nil {
		panic(err)
	}
}
