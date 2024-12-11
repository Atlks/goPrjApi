package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/unidoc/unioffice/docx"
)

// 将 Word 文档转换为 TXT 文件
func convertWordToTxt(srcDir, destDir string) error {
	// 确保目标目录存在
	err := os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("无法创建目标目录: %v", err)
	}

	// 遍历源目录，查找所有 Word 文档
	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 只处理 .docx 文件
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".docx") {
			// 转换 Word 文件为 TXT
			err := convertFile(path, destDir)
			if err != nil {
				return fmt.Errorf("转换文件失败 %s: %v", path, err)
			}
			fmt.Printf("已转换文件: %s\n", path)
		}
		return nil
	})

	return err
}

// 将单个 Word 文档转换为 TXT 文件
func convertFile(wordFilePath, destDir string) error {
	// 打开 Word 文档
	doc, err := docx.Open(wordFilePath)
	if err != nil {
		return fmt.Errorf("无法打开 Word 文件 %s: %v", wordFilePath, err)
	}
	defer doc.Close()

	// 提取文档中的文本
	text := doc.Text()

	// 创建对应的 TXT 文件路径
	txtFilePath := filepath.Join(destDir, strings.TrimSuffix(filepath.Base(wordFilePath), ".docx")+".txt")

	// 将文本写入 TXT 文件
	err = ioutil.WriteFile(txtFilePath, []byte(text), 0644)
	if err != nil {
		return fmt.Errorf("无法写入 TXT 文件 %s: %v", txtFilePath, err)
	}

	return nil
}

func main() {
	srcDir := "./source"  // 源目录，包含 Word 文件
	destDir := "./output" // 输出目录，保存 TXT 文件

	// 转换文件
	err := convertWordToTxt(srcDir, destDir)
	if err != nil {
		fmt.Println("转换失败:", err)
	} else {
		fmt.Println("所有文件转换完成！")
	}
}
