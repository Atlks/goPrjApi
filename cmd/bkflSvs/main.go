package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	mdfDate := "2024-12-09" // 指定日期
	modifiedAfter, err := time.Parse("2006-01-02", mdfDate)
	if err != nil {
		fmt.Println("时间解析失败:", err)
		return
	}

	sourceDirPath := "C:\\0prjNtpc\\digital-bet-service" // 主目录路径
	// sourceDirPath = "C:\\0prjNtpc\\digital-bet-service - p3"
	sourceDirPath = "C:\\0prjNtpc\\digital-bet-service -thb"

	targetDirPath := "/bkPrjPgthb241209" // 输出目录路径

	// targetDirPath = "bkPrjDBS24111"
	extnames := "java,xml" // 要处理的文件扩展名

	// 执行文件复制
	err = copyModifiedFiles(modifiedAfter, extnames, sourceDirPath, targetDirPath)
	if err != nil {
		fmt.Println("文件复制失败:", err)
	}
}

// 复制修改日期在指定日期之后的文件到输出目录，并保留原有的目录结构
func copyModifiedFiles(modifiedAfter time.Time, extnames, sourceDirPath, targetDirPath string) error {
	// 打开源目录
	sourceDir, err := os.Open(sourceDirPath)
	if err != nil {
		return fmt.Errorf("主目录不存在：%v", sourceDirPath)
	}
	defer sourceDir.Close()

	// 创建目标目录
	err = os.MkdirAll(targetDirPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("无法创建目标目录：%v", targetDirPath)
	}

	// 遍历源目录的文件
	err = filepath.Walk(sourceDirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// 获取文件的修改时间
		modifiedDate := info.ModTime()
		if modifiedDate.After(modifiedAfter) {
			// 获取扩展名
			fileExtension := filepath.Ext(path)
			if fileExtension == "" {
				return nil
			}

			// 去掉点并检查扩展名是否符合
			fileExt := strings.TrimPrefix(fileExtension, ".")
			if isInSet(fileExt, extnames) {
				// 获取相对路径
				relativePath := strings.TrimPrefix(path, sourceDirPath)
				targetFile := filepath.Join(targetDirPath, relativePath)

				// 创建目标文件夹
				err = os.MkdirAll(filepath.Dir(targetFile), os.ModePerm)
				if err != nil {
					return fmt.Errorf("无法创建目录：%v", filepath.Dir(targetFile))
				}

				// 复制文件
				err := copyFile(path, targetFile)
				if err != nil {
					return fmt.Errorf("文件复制失败：%v", err)
				}

				fmt.Printf("已复制文件: %s 到 %s\n", path, targetFile)
			}
		}

		return nil
	})

	return err
}

// 复制文件函数
func copyFile(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

// 判断扩展名是否在extnames中，extnames是逗号分隔的字符串
func isInSet(ext, extnames string) bool {
	extList := strings.Split(extnames, ",")
	for _, e := range extList {
		if strings.TrimSpace(e) == ext {
			return true
		}
	}
	return false
}
