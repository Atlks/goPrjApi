package lib

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/go-ego/gse"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	seg gse.Segmenter
)

// dbgCls 模拟调试功能
type dbgCls struct{}

func (dbg *dbgCls) setDbgFunEnter(method string, args ...interface{}) {
	fmt.Printf("Entering method: %s with args: %v\n", method, args)
}

var dbg = &dbgCls{}

// 模拟其他方法
func setuNameFrmTgmsgJson(messageElement map[string]interface{}, o map[string]interface{}) {
	if from, ok := messageElement["from"].(map[string]interface{}); ok {
		if firstName, ok := from["first_name"].(string); ok {
			o["first_name"] = firstName
		}
	}
}

func setGrpFromTgjson(messageElement map[string]interface{}, o map[string]interface{}) {
	if chat, ok := messageElement["chat"].(map[string]interface{}); ok {
		if title, ok := chat["title"].(string); ok {
			o["grp"] = title
		}
	}
}

func ConvertUnixTimeStampToDateTime(stmp int64) string {
	return time.Unix(stmp, 0).Format("2006-01-02 15:04:05")
}

func File_exists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func ReadLinesToHashSet(filepath string) map[string]struct{} {
	set := make(map[string]struct{})
	if !File_exists(filepath) {
		return set
	}

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return set
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		set[line] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
	return set
}

func IsNumeric(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func IsAllPunctuation(s string) bool {
	for _, r := range s {
		if !isPunctuation(r) {
			return false
		}
	}
	return true
}

func isPunctuation(r rune) bool {
	// 简单判断是否为标点符号
	return r >= '!' && r <= '/'
}

func ormJSonFLSave(doc map[string]interface{}, filename string) {
	data, err := json.Marshal(doc)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

func ReadAndCreateIndex4tgmsg(directoryPath string) {
	const __METHOD__ = "ReadAndCreateIndex4tgmsg"
	dbg.setDbgFunEnter(__METHOD__, directoryPath)

	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		return
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(directoryPath, file.Name())
			jsonContent, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Printf("An error occurred: %s\n", err)
				continue
			}

			var doc map[string]interface{}
			err = json.Unmarshal(jsonContent, &doc)
			if err != nil {
				fmt.Printf("An error occurred: %s\n", err)
				continue
			}

			if messageElement, ok := doc["message"].(map[string]interface{}); ok {
				if textElement, ok := messageElement["text"].(string); ok {
					fmt.Println(textElement)
					crtIdx(messageElement, textElement)
				} else {
					fmt.Printf("The 'message' property in the file %s is not an object.\n", file.Name())
				}
			}
		}
	}
}

func crtIdx(messageElement map[string]interface{}, textElement string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in crtIdx", r)
		}
	}()

	othInf := make(map[string]interface{})
	setuNameFrmTgmsgJson(messageElement, othInf)
	setGrpFromTgjson(messageElement, othInf)

	if date, ok := messageElement["date"].(float64); ok {
		stmp := int64(date)
		othInf["timeStamp"] = stmp
		othInf["time"] = ConvertUnixTimeStampToDateTime(stmp)
	}

	CreateIndexPart2(textElement, othInf)
}

var TrdSmplfconfigMap, _ = iniToMap("D:\\0prj\\mdsj\\mdsjprj\\libBiz\\trd2smpLibV2.ini")

func TraditionalToSimplified(traditional string) string {
	// 将繁体中文文本进行分词
	// 使用 mahonia 进行繁简体转换
	//traditional = "可轉換得到相對應的繁體字或簡體字"
	simplified := TradtToSmplf(traditional, TrdSmplfconfigMap)
	//decoder := mahonia.NewDecoder("gbk")
	//	simplified := decoder.ConvertString(traditional)
	return simplified
}

func TradtToSmplf(input string, replacements map[string]string) string {
	var result strings.Builder

	for _, char := range input {
		if replacement, found := replacements[string(char)]; found {
			result.WriteString(replacement)
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}
func CreateIndexPart2(msgxv1 string, grpinfo map[string]interface{}) {
	msgx := TraditionalToSimplified(msgxv1)
	//	msgx := msgxv1
	words := fenci(msgx)
	for _, word := range words {

		word = strings.TrimSpace(word)
		if len2024(word) <= 1 || IsNumeric(word) || IsAllPunctuation(word) {
			continue
		}

		doc := make(map[string]interface{})
		timestamp := time.Now().Format("20060102_150405_000")
		doc["id"] = timestamp
		doc["kwd"] = word
		doc["txt"] = msgxv1
		doc["grpinfo"] = grpinfo

		filename := fmt.Sprintf("fullTxtSrchIdxdataDir/%s.json", word)
		SaveJson(doc, filename)
	}
}

func len2024(word string) int {
	// 计算字符串长度
	charCount := utf8.RuneCountInString(word)
	return charCount
}

func fenci(msgx string) []string {
	// 创建 gse 分词器实例
	segmenter, _ := gse.New()

	// 加载用户自定义词典
	userDict := ReadLinesToHashSet("user_dict.txt")
	for word := range userDict {
		segmenter.AddToken(word, 1)
	}

	// 加载位置词典
	postnKywd位置词set := ReadLinesToHashSet("位置词.txt")
	for word := range postnKywd位置词set {
		segmenter.AddToken(word, 1)
	}

	// 执行分词
	segments := segmenter.Cut(msgx, true)
	return segments
}

// jieba fenci err..bcs  NewJieba err ..
//func fenciDep(msgx string) []string {
//	// 创建结巴分词器实例
//	segmenter := gojieba.NewJieba()
//	defer segmenter.Free()
//
//	//segmenter.LoadUserDict("user_dict.txt")
//	segmenter.AddWord("会所")
//	segmenter.AddWord("妙瓦底")
//	segmenter.AddWord("御龙湾")
//
//	userDict := ReadLinesToHashSet("user_dict.txt")
//	for word := range userDict {
//		segmenter.AddWord(word)
//	}
//
//	postnKywd位置词set := ReadLinesToHashSet("位置词.txt")
//	for word := range postnKywd位置词set {
//		segmenter.AddWord(word)
//	}
//
//	words := segmenter.CutForSearch(msgx, true)
//	return words
//}
