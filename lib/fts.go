package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

func SearchMatch(dataDir string, matchKwds string) []map[string]any {

	// 新建一个空切片（动态数组）
	//var srchRztArr []Object
	splitStr := strings.Fields(matchKwds)
	// 定义一个变量来存储解析后的JSON数组
	var result []map[string]interface{}
	needini := true
	// 使用 for range 循环遍历数组
	for index, value := range splitStr {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
		value = strings.TrimSpace(value)
		if len(value) > 0 {
			fpath := dataDir + "/" + value + ".json"
			maparr, _ := readJSONFileToMapArray(fpath)
			if needini {
				result = maparr
				needini = false
			} else {
				result = intersectArrays(result, maparr)
			}

			//	append(srchRztArr, maparr)

		}

	}

	//做交集
	return result

}

// intersectArrays 计算两个[]map[string]interface{}的交集
func intersectArrays(arr1, arr2 []map[string]interface{}) []map[string]interface{} {
	result := []map[string]interface{}{}

	// 使用map存储每个map中的元素，key为id
	map1 := make(map[string]map[string]interface{})
	for _, m := range arr1 {
		map1[m["id"].(string)] = m
	}

	// 遍历arr2，如果元素在map1中存在，则加入结果集
	for _, m := range arr2 {
		if _, ok := map1[m["id"].(string)]; ok && reflect.DeepEqual(map1[m["id"].(string)], m) {
			result = append(result, m)
		}
	}

	return result
}

func readJSONFileToMapArray(filename string) ([]map[string]interface{}, error) {
	// 打开JSON文件
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	// 读取文件内容
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}

	// 定义一个变量来存储解析后的JSON数组
	var result []map[string]interface{}

	// 解析JSON数据
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling JSON: %v", err)
	}

	return result, nil
}
