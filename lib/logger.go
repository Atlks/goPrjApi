package lib

import (
	"log"
	"os"
)

// str or err
func Log_info(s any) {

	file, _ := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	defer file.Close() // 关闭文件

	log.SetOutput(file)
	log.Print(s)

}
