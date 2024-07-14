package lib

import (
	"encoding/json"
	"github.com/tealeg/xlsx"
)

func ReadXlsxToJson(filePath string) (string, error) {
	// 打开xlsx文件
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		return "", err
	}

	// 创建一个map存储数据
	data := make(map[string]string)

	// 遍历文件中的所有表格和行
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			// 检查行是否有至少3列
			if len(row.Cells) >= 3 {
				key := row.Cells[0].String()
				value := row.Cells[2].String()
				data[key] = value
			}
		}
	}

	// 将map序列化为json
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

//func getLemma(word string) string {
//	// 创建一个新的 lemmatizer
//	lemmatizer, err := golem.NewLemmatizer()
//	if err != nil {
//		fmt.Printf("Error creating lemmatizer: %v\n", err)
//		return ""
//	}
//
//	// 获取单词的原始形式lemma root
//	lemma, err := lemmatizer.Lemma(word, "noun")
//	if err != nil {
//		fmt.Printf("Error getting lemma: %v\n", err)
//		return ""
//	}
//
//	return lemma
//}
