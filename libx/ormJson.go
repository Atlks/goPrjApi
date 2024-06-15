package libx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// 模拟的调试函数
//func setDbgFunEnter(method string, args ...interface{}) {
//	fmt.Printf("Enter: %s, args: %v\n", method, args)
//}

func setDbgValRtval(method string, result interface{}) {
	fmt.Printf("Return: %s, result: %v\n", method, result)
}

// 模拟的查询函数，返回一个切片，其中包含多个映射
// qrySglFL is the equivalent of the provided C# function in Go.

// arraySlice is a placeholder function to slice the array. Implement as required.

// qrySglFL is the equivalent of the provided C# function in Go.
func qrySglFL(dbFileName string) []map[string]interface{} {
	method := "qrySglFL" // Placeholder for method name retrieval in Go

	setDbgFunEnter(method, dbFileName)

	if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
		err := ioutil.WriteFile(dbFileName, []byte("[]"), 0644)
		if err != nil {
			panic(fmt.Sprintf("Error creating file: %v", err))
		}
	}

	txt, err := ioutil.ReadFile(dbFileName)
	if err != nil {
		panic(fmt.Sprintf("Error reading file: %v", err))
	}
	if len(txt) == 0 {
		txt = []byte("[]")
	}

	var list []map[string]interface{}
	err = json.Unmarshal(txt, &list)
	if err != nil {
		panic(fmt.Sprintf("Error unmarshalling JSON: %v", err))
	}

	setDbgValRtval(method, arraySlice(list, 0, 1))

	return list
}

// 合并两个切片
func arrayMerge(arr1, arr2 []map[string]interface{}) []map[string]interface{} {
	return append(arr1, arr2...)
}

// 切片切割
func arraySlice(arr []map[string]interface{}, start, end int) []map[string]interface{} {
	if end > len(arr) {
		end = len(arr)
	}
	return arr[start:end]
}

//func arraySlice(list []map[string]interface{}, start, end int) []map[string]interface{} {
//	if start < 0 || end > len(list) || start > end {
//		return []map[string]interface{}{}
//	}
//	return list[start:end]
//}

func QryJson(dbfS string) []map[string]interface{} {
	method := "qry" // 手动设置方法名
	setDbgFunEnter(method, dbfS)
	dbArr := strings.Split(dbfS, ",")

	var arr []map[string]interface{}
	for _, dbf := range dbArr {
		dbf = strings.TrimSpace(dbf)
		if dbf == "" {
			continue
		}

		// 检查文件所在目录是否存在，不存在则创建目录
		directory := filepath.Dir(dbf)
		if _, err := os.Stat(directory); os.IsNotExist(err) {
			os.MkdirAll(directory, os.ModePerm)
		}

		if _, err := os.Stat(dbf); os.IsNotExist(err) {
			fmt.Printf("not exist file dbf => %s\n", dbf)
			continue
		}

		sortedLists := qrySglFL(dbf)
		arr = arrayMerge(arr, sortedLists)
	}

	setDbgValRtval(method, arraySlice(arr, 0, 2))

	return arr
}

// 模拟的查询函数
func qryDep(dbFileName string) []map[string]interface{} {
	method := "qryDep"
	setDbgFunEnter(method, dbFileName)

	if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
		ioutil.WriteFile(dbFileName, []byte("[]"), 0644)
	}

	txt, err := ioutil.ReadFile(dbFileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	if strings.TrimSpace(string(txt)) == "" {
		txt = []byte("[]")
	}

	var list []map[string]interface{}
	err = json.Unmarshal(txt, &list)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}

	setDbgValRtval(method, list[:min(3, len(list))])

	return list
}

// 添加或替换键值
func addRplsKeyV(listIot map[string]map[string]interface{}, key string, objSave map[string]interface{}) {
	listIot[key] = objSave
}

// 从映射生成列表
func lstFrmIot(listIot map[string]map[string]interface{}) []map[string]interface{} {
	var arrayList []map[string]interface{}
	for _, value := range listIot {
		arrayList = append(arrayList, value)
	}
	return arrayList
}

// 写入文件
func wriToDbf(lst interface{}, dbfl string) {
	data, err := json.MarshalIndent(lst, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	ioutil.WriteFile(dbfl, data, 0644)
}

func SaveJson(SortedList1 map[string]interface{}, dbfile string) {
	method := "save"
	setDbgFunEnter(method, dbfile)

	if !fileExists(dbfile) {
		writeFileIfNotExist(dbfile, "[]")
	}

	list := qryDep(dbfile)
	listIot := lst2IOT(list)

	key := SortedList1["id"].(string)
	addRplsKeyV(listIot, key, SortedList1)

	saveListHpmod := lstFrmIot(listIot)
	wriToDbf(saveListHpmod, dbfile)

	setDbgValRtval(method, 0)
}

// fileExists 检查文件是否存在
func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

// writeFileIfNotExist 检查文件是否存在，如果不存在，则创建目录和文件，并写入"ttt"
func writeFileIfNotExist(filePath string, txt string) error {
	// 获取文件目录
	dir := filepath.Dir(filePath)

	// 检查目录是否存在，如果不存在，则创建目录
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}

	// 检查文件是否存在，如果不存在，则创建文件并写入"ttt"
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		defer file.Close()

		_, err = file.WriteString(txt)
		if err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	} else {
		fmt.Println("File already exists, skipping creation.")
	}

	return nil
}

// 模拟的函数，根据需要自行实现
func lst2IOT(list []map[string]interface{}) map[string]map[string]interface{} {
	listIot := make(map[string]map[string]interface{})
	for _, item := range list {
		id := item["id"].(string)
		listIot[id] = item
	}
	return listIot
}

// 最小值函数
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
