package libx

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

func Func_get_args() []interface{} {
	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	name := fn.Name()

	fmt.Printf("Calling function: %s\n", name)

	var args []interface{}

	// Using reflect to get the arguments of the caller function
	v := reflect.ValueOf(fn)
	for i := 0; i < v.Type().NumIn(); i++ {
		args = append(args, v.Type().In(i))
	}

	return args
}

// 全局变量
var dbgpad int

// 设置调试函数入口
func SetDbgFunEnter(method string, funcGetArgs interface{}) {
	// 增加缩进
	dbgpad += 4

	// 序列化参数
	funcGetArgsJSON, err := json.Marshal(funcGetArgs)
	if err != nil {
		fmt.Println("Error serializing arguments:", err)
		return
	}

	// 构建日志消息
	msglog := fmt.Sprintf("\n\n\n%s FUN %s((%s))", strings.Repeat(" ", dbgpad), method, string(funcGetArgsJSON))

	// 打印日志消息
	fmt.Println(msglog)
}

func SetDbgVal(method string, vname string, val string) {
	msglog := fmt.Sprintf("%s%s():: %s=>%s", strings.Repeat(" ", dbgpad+3), method, vname, val)
	fmt.Println(msglog)
}

func SetDbgValRtval(method string, results interface{}) {
	resultsJSON, err := json.Marshal(results)
	if err != nil {
		fmt.Println("Error serializing results:", err)
		return
	}

	msglog := fmt.Sprintf("%sENDFUN %s():: ret=>%s", strings.Repeat(" ", dbgpad), method, string(resultsJSON))
	fmt.Println(msglog)
	dbgpad -= 4
}
