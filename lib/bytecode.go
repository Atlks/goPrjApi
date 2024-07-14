package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

// 模拟 ldfld 函数，根据字段名获取字段值或默认值
func ldfld(obj map[string]interface{}, fld string, defVal interface{}) interface{} {
	if val, ok := obj[fld]; ok {
		return val
	}
	return defVal
}

// 模拟 ldfldAsStr 函数，获取字段值并转换为字符串
func ldfldAsStr(obj map[string]interface{}, fld string, defVal interface{}) string {
	val := ldfld(obj, fld, "").(string)
	return val
}

// 模拟 getFld 函数，根据对象类型调用相应的方法获取字段值或返回默认值
func getFld(obj interface{}, fld string, defVal interface{}) interface{} {
	if m, ok := obj.(map[string]interface{}); ok {
		return ldfld(m, fld, defVal)
	} else {
		return ldfld(obj.(map[string]interface{}), fld, defVal)
	}
}

// 模拟 foreach_map 函数，遍历 map 并应用处理函数
func foreach_map(m map[string]interface{}, fun func(string, interface{}) interface{}) {
	for key, value := range m {
		// 仿照 C# 的 Convert.ToInt64(de.Key) == Program.groupId 进行判断
		// if Convert.ToInt64(de.Key) == Program.groupId {
		//     continue
		// }
		// var chatid = Convert.ToInt64(de.Key)

		// 尝试执行传入的函数
		tryExecute(key, value, fun)
	}
}

// 模拟 tryExecute 函数，用于尝试执行传入的函数并处理异常
func tryExecute(key string, value interface{}, fun func(string, interface{}) interface{}) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in tryExecute:", r)
			// print_catchEx("foreach_map", r)
		}
	}()

	// 执行传入的函数
	fun(key, value)
}

// 定义一个函数类型，代表被调用的方法签名
type FuncType func(string) string

// 模拟 foreach_HashSet 函数
func foreach_HashSet(originalSet map[string]bool, fun FuncType) map[string]bool {
	updatedSet := make(map[string]bool)

	// 遍历原始集合
	for str := range originalSet {
		updatedSet[fun(str)] = true
	}

	return updatedSet
}

// 定义一个函数类型，代表被调用的方法签名
type methodFunc func(args ...interface{}) interface{}

// 模拟委托类型
type Delegate struct {
	Method methodFunc
}

// 模拟函数：print_call_FunArgs
func print_call_FunArgs(methodName string, args []interface{}) {
	fmt.Printf("Calling method %s with arguments: %v\n", methodName, args)
}

// 模拟函数：json_encode_noFmt
func json_encode_noFmt(args []interface{}) string {
	// 这里只是简单地打印参数，实际中需要根据具体需求实现
	return fmt.Sprintf("%v", args)
}

// 模拟函数：jmp2end
func jmp2end() {
	fmt.Println("jmp2end called")
}

// 模拟函数：logErr2024
func logErr2024(e error, methodName, errdir string, dbgobj map[string]interface{}) {
	fmt.Printf("Error in method %s: %v\n", methodName, e)
	// 这里只是简单地打印错误日志，实际中需要根据具体需求实现
}

// SaveToFile 将字符串内容保存到指定的文本文件中
func SaveToFile(filename, content string) error {
	// 打开文件，如果不存在则创建，如果存在则清空文件
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将内容写入文件
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

// SaveToFileUsingIoutil 使用 ioutil.WriteFile 将字符串内容保存到指定的文本文件中
func SaveToFileUsingIoutil(filename, content string) error {
	// 将内容写入文件
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}

	return nil
}

// 模拟函数：call_user_func
func call_user_func(callback Delegate, args ...interface{}) interface{} {
	methodName := ""
	//reflect.ValueOf(callback.Method).Pointer().Name()

	print_call_FunArgs(methodName, args)

	var o interface{}
	var e error

	// 模拟异常处理
	defer func() {
		if r := recover(); r != nil {
			if e1, ok := r.(jmp2endEx); ok {
				panic(e1)
			}
			if e2, ok := r.(error); ok {
				if reflect.TypeOf(r).String() == "System.Reflection.TargetInvocationException" && fmt.Sprintf("%v", e2) == "jmp2endEx" {
					fmt.Printf("---catch ex----call mtdh:%s  prm:%s\n", methodName, json_encode_noFmt(args))
					fmt.Println("Caught jmp2endEx")
					print_ret(methodName, 0)
					jmp2end()
				}
			}
			// 输出异常信息
			fmt.Printf("---catch ex----call mtdh:%s  prm:%s\n", methodName, json_encode_noFmt(args))
			fmt.Println(r)
			dbgobj := make(map[string]interface{})
			dbgobj["mtth"] = fmt.Sprintf("%s(((%s)))", methodName, json_encode_noFmt(args))
			logErr2024(e, methodName, "errdir", dbgobj)
		}
	}()

	// 模拟调用方法
	o = callback.Method(args...)

	// 输出方法调用结果
	if o != nil {
		print_ret(methodName, o)
	} else {
		print_ret(methodName, 0)
	}

	return o
}

func (e jmp2endEx) Error() string {
	return "jmp2endEx triggered"
}

///
/*
func main() {
	// 示例调用Jmp2end函数
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(jmp2endEx); ok {
				fmt.Println("Recovered from jmp2endEx")
			} else {
				panic(r) // 重新抛出非jmp2endEx类型的异常
			}
		}
	}()

	Jmp2end()
}

*/
func Jmp2end() {
	// jmp2exitFlag = true; // 这个flag在Go中可以根据需要定义和使用
	panic(jmp2endEx{})
}

// 自定义异常类型 jmp2endEx
type jmp2endEx struct{}

func ToString(value interface{}) string {
	if value == nil {
		return ""
	}

	// 如果值本身就是字符串类型
	if str, ok := value.(string); ok {
		return str
	}

	// 其他类型转换为字符串
	return fmt.Sprintf("%v", value)
}
func ConvertToString(value interface{}) string {
	if value == nil {
		return ""
	}

	// 如果值本身就是字符串类型
	if str, ok := value.(string); ok {
		return str
	}

	// 其他类型转换为字符串
	return fmt.Sprintf("%v", value)
}
func mai77n() {
	// 测试调用 call_user_func
	callback := Delegate{Method: YourMethodName} // 假设 YourMethodName 是一个已定义的函数
	args := []interface{}{}                      // 可以根据需要添加参数

	result := call_user_func(callback, args...)
	fmt.Println("Result:", result)
}

// 模拟具体方法实现
func YourMethodName(args ...interface{}) interface{} {
	fmt.Println("YourMethodName called with args:", args)
	return "YourMethodName result"
}

// 模拟函数：print_ret
func print_ret(methodName string, result interface{}) {
	fmt.Printf("Method %s returned: %v\n", methodName, result)
}
