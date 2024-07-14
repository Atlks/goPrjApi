package lib

import (
	"encoding/json"
	"fmt"
	"github.com/gorhill/cronexpr"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func RunTmrTasksCron() {
	// 创建一个定时器，每分钟检查一次是否需要执行任务
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			tmrTask1start()
		}
	}
}

func tmrTask1start() {
	tmrtaskDir := fmt.Sprintf("%s/cfg/tmrtask", "prjdir")
	list := getListFrmDir(tmrtaskDir)
	for _, hs := range list {
		now := time.Now()
		cronExpression := hs["cron"].(string)
		schedule := cronexpr.MustParse(cronExpression)
		RunDateTime := schedule.Next(now)
		if RunDateTime.Hour() == now.Hour() && RunDateTime.Minute() == now.Minute()+1 {
			zhuliLog := fmt.Sprintf("tmrlg/%s%d%d%d.json", hs["basename"].(string), now.Month(), now.Day(), now.Hour())
			if _, err := os.Stat(zhuliLog); os.IsNotExist(err) {
				err = ioutil.WriteFile(zhuliLog, []byte("pushlog"), 0644)
				if err != nil {
					fmt.Println("Error writing log file:", err)
					continue
				}
				callx(hs["fun"].(string))
			}
		}
	}
}

func CallTmrTasks() {
	prjdir := "/path/to/prjdir" // 替换为实际的项目目录路径
	tmrtaskDir := filepath.Join(prjdir, "cfg", "tmrtask")
	list := getListFrmDir(tmrtaskDir)
	for _, hs := range list {
		now := time.Now()

		times := strings.Split(hs["time"].(string), " ")
		zhuliLog := fmt.Sprintf("tmrlg/%s%d%d%d.json", hs["basename"], now.Month(), now.Day(), now.Hour())
		if isIn(now.Hour(), times) && now.Minute() == 1 && !fileExists(zhuliLog) {
			err := ioutil.WriteFile(zhuliLog, []byte("pushlog"), 0644)
			if err != nil {
				fmt.Println("Error writing file:", err)
			}

			// var txtkeepBtnMenu string = "" // 如果需要设置内容，可以在这里设置

			callx(hs["fun"].(string))
		}
	}
}

// 判断小时是否在列表中的函数
func isIn(hour any, times []string) bool {
	for _, o := range times {
		if strings.EqualFold((ToString(hour)), o) {
			return true
		}
	}
	return false
}

// 模拟调用外部函数的函数
func callx(fun string) {
	// 这里实现调用外部函数的逻辑
}
func getListFrmDir(directoryPath string) []map[string]interface{} {
	fileList := make([]map[string]interface{}, 0)

	// 获取目录中所有 JSON 文件的路径
	jsonFiles, err := filepath.Glob(filepath.Join(directoryPath, "*.json"))
	if err != nil {
		fmt.Println("Error globbing JSON files:", err)
		return fileList
	}

	for _, filePath := range jsonFiles {
		// 读取文件内容
		jsonContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", filePath, err)
			continue
		}

		// 解析 JSON 文件为 map[string]interface{}
		var jsonObject map[string]interface{}
		if err := json.Unmarshal(jsonContent, &jsonObject); err != nil {
			fmt.Printf("Error parsing JSON file %s: %v\n", filePath, err)
			continue
		}

		// 构造结果 map
		hashtable := make(map[string]interface{})
		for key, value := range jsonObject {
			hashtable[key] = value
		}

		// 添加文件名和路径信息
		hashtable["fname"] = filepath.Base(filePath)
		hashtable["fpath"] = filePath
		// 获取文件名（不包括扩展名）
		hashtable["basename"] = strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))

		// 添加到列表中
		fileList = append(fileList, hashtable)
	}

	return fileList
}

// 模拟检查文件是否存在的函数
//func fileExists(filename string) bool {
//	_, err := os.Stat(filename)
//	return !os.IsNotExist(err)
//}
