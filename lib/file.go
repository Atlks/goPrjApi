package lib

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// wrtLgTypeDate 创建目录并将对象编码为 JSON 格式写入带有时间戳的文件中
func FwrtLgTypeDate(logdir string, o interface{}) {
	// 创建目录
	err := os.MkdirAll(logdir, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("Failed to create directory: %s", err))
	}

	// 获取当前时间并格式化为文件名
	timestamp := time.Now().Format("20060102_150405_000")
	fileName := fmt.Sprintf("%s/%s.json", logdir, timestamp)

	// 将对象编码为 JSON
	data := Json_encode(o)

	// 写入数据到文件
	Ffile_put_contents(fileName, data, false)
	fmt.Printf("Successfully wrote to file: %s\n", fileName)
}

// filePutContents 写入字符串数据到文件
// 如果文件不存在，则创建该文件；如果文件已存在，根据标志决定是覆盖还是追加写入。
// 返回写入的字节数
func Ffile_put_contents(filename string, data string, append bool) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", r)
		}
	}()

	var file *os.File
	var err error

	if append {
		// 以追加模式打开文件
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	} else {
		// 以创建模式打开文件（会覆盖文件内容）
		file, err = os.Create(filename)
	}

	if err != nil {
		panic(fmt.Sprintf("Failed to open or create file: %s", err))
	}
	defer file.Close()

	// 将字符串转换为字节数组
	bytesData := []byte(data)

	// 写入数据到文件
	bytesWritten, err := file.Write(bytesData)
	if err != nil {
		panic(fmt.Sprintf("Failed to write to file: %s", err))
	}

	return bytesWritten
}

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
