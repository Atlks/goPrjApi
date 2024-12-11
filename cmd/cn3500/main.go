package main

import (
	"awesomeProject/lib"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	var fl = "C:\\0prj\\paymentJava2024\\lib\\1000cnschar.txt"
	arr1 := rdFileSpltTxtToArrDedulp(fl)
	//fmt.Printf(lib.EncodeJson(arr1))
	arr2 := rdFileSpltTxtToArrDedulp("C:\\0prj\\paymentJava2024\\lib\\2500cnChar.txt")
	arr3 := MergeArr(arr1, arr2)
	arr3 = fltArrEmptyNnotCnchar(arr3)
	foreachArr(arr3)
	json3500 := lib.EncodeJson(arr3)
	fmt.Printf(json3500)
	lib.Write("3500cn chars.json", json3500)
	fmt.Printf(strconv.Itoa(lib.LenArr(arr3)))
}

func foreachArr(arr3 interface{}) {
	// 将输入的切片转换为 []string
	arr1 := toStrArr(arr3)
	// 遍历并打印每个元素
	for _, r := range arr1 {
		fmt.Println(r) // 使用 fmt.Println 进行输出
	}
}

// toStrArr 函数：将 interface{} 转换为 []string
func toStrArr(arr3 interface{}) []string {
	// 检查 arr3 是否是切片类型
	v := reflect.ValueOf(arr3)
	if v.Kind() != reflect.Slice {
		fmt.Println("Input is not a slice")
		return []string{}
	}

	// 用于存储字符串切片
	var result []string

	// 遍历切片
	for i := 0; i < v.Len(); i++ {
		item := v.Index(i).Interface()
		str, ok := item.(string)
		if ok {
			result = append(result, str)
		}
	}

	return result
}

/*
*
过滤掉非汉字字符，以及空字符
*/
func fltArrEmptyNnotCnchar(arr3 interface{}) []string {
	// 检查 arr3 是否是切片类型
	v := reflect.ValueOf(arr3)
	if v.Kind() != reflect.Slice {
		fmt.Println("Input is not a slice")
		return []string{}
	}

	// 用于存储过滤后的结果
	result := []string{}

	// 遍历切片
	for i := 0; i < v.Len(); i++ {
		item := v.Index(i).Interface()
		str, ok := item.(string)
		if !ok {
			continue // 如果不是字符串，跳过
		}

		if str == "," {
			fmt.Println("dbg")
		}

		// 如果是汉字字符串且非空，添加到结果中
		if isCnChar(str) {
			result = append(result, str)
		}
	}

	return result
}

// 如果是汉字字符串
func isCnChar(str string) bool {
	if str == "" {
		return false
	}
	// 检查字符串是否仅包含汉字
	for _, r := range str {
		if !unicode.Is(unicode.Han, r) {
			return false
		}
	}
	return true
}

// MergeArr 函数：合并两个字符串切片，并去除重复元素
func MergeArr(arr1 []string, arr2 []string) interface{} {
	// 创建一个 map 用于存储唯一值
	uniqueMap := make(map[string]bool)
	result := []string{}

	// 将第一个切片的元素加入 map 和结果中
	for _, item := range arr1 {
		if !uniqueMap[item] {
			uniqueMap[item] = true
			result = append(result, item)
		}
	}

	// 将第二个切片的元素加入 map 和结果中
	for _, item := range arr2 {
		if !uniqueMap[item] {
			uniqueMap[item] = true
			result = append(result, item)
		}
	}

	return result
}

func rdFileSpltTxtToArrDedulp(fl string) []string {
	var txt = ReadFile(fl)
	var arr1 []string = split(txt)

	// 定义一个map，键是string类型，值也是string类型
	//var myMap map[string]string

	// 使用make函数初始化map
	//myMap = make(map[string]string)
	arr1 = deDulip(arr1)
	return arr1
}

/*
*
去除重复的item
*/
func deDulip(list []string) []string {
	// 创建一个 map 用于存储已出现的元素
	seen := make(map[string]bool)
	// 用于存储去重后的结果
	result := []string{}

	// 遍历输入的列表
	for _, item := range list {
		// 如果元素未出现过，则添加到结果中，并标记为已出现
		if !seen[item] {
			result = append(result, item)
			seen[item] = true
		}
	}
	return result
}

/**
切分字符串，返回一个字符数组
*/
// split 函数将字符串切分为字符数组并返回
func split(txt string) []string {
	// 将字符串按字符切分
	return strings.Split(txt, "")
}

func ReadFile(fl string) string {
	// 读取文件内容
	content, err := ioutil.ReadFile(fl)
	if err != nil {
		// 如果出现错误，记录日志并返回空字符串
		log.Printf("Error reading file %s: %v", fl, err)
		return ""
	}
	// 返回文件内容作为字符串
	return string(content)
}
