package main

import (
	"awesomeProject/libx"
	"fmt"
	"path/filepath"
)

func main() {
	// Example usage of the save function

	// 初始化变量
	saveDataDir := "测试数据表目录"
	dbg := make(map[string]interface{})
	sortedListNew := map[string]interface{}{
		"国家": "China",
		"name": "tomm", "id": "007",
	}

	// 调用函数

	str := libx.Qe_save(sortedListNew, saveDataDir, func(row map[string]interface{}) int {

		prtnKey := "国家"
		wrtFile := filepath.Join(saveDataDir, fmt.Sprintf("%v.json", row[prtnKey]))
		//使用增量还是全量模式，都是由存储引擎决定的，orm框架与查询引擎是不管的，主管业务层面
		libx.SaveJson(row, wrtFile)
		return 0
	}, dbg)
	fmt.Print(str)

	fmt.Print(111)
	libx.MthdFrmPkg1()

	// 创建 session 数据
	chtsSesss := map[string]interface{}{
		"id": "1",
		"nm": "tommm",
	}

	dbFileName := "example2024.db"

	// 调用 save 方法

	libx.Save(chtsSesss, dbFileName)

	dbFileName = "example2024.json"
	//libx.SaveJson(chtsSesss, dbFileName)
	list := libx.QryJson(dbFileName)
	fmt.Print(libx.JsonEncode(list))
	//m()
}
