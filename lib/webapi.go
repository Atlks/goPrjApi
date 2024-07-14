package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func StartWebapi(httpHdlrSpel func(HttpContext), api_prefix string) {
	// 配置Kestrel监听特定端口
	port := 5200

	// 定义请求处理函数
	requestHandler := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// 处理异常
				log.Printf("Recovered in startWebapi: %v", err)
				// 如果需要，可以添加额外的错误处理逻辑，例如打印错误或记录日志
			}
		}()
		context := HttpContext{Response: w, Request: r}
		httpHdlr(context, api_prefix, httpHdlrSpel)
	}

	// 启动HTTP服务器
	http.HandleFunc("/", requestHandler)
	log.Printf("Starting server on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}

func httpHdlr(context HttpContext, api_prefix string, httpHdlrApiSpecl func(HttpContext)) {
	//context2 := context
	//	response := context2.Response
	request := context.Request
	//	url := fmt.Sprintf("%s://%s%s%s", request.URL.Scheme, request.Host, request.URL.Path, request.URL.RawQuery)

	// 获取查询字符串
	queryString := request.URL.RawQuery
	path := request.URL.Path

	// 静态资源处理器映射表
	extNhdlrChoosrMaplist := map[string]func(HttpContext){
		"txt":  Html_httpHdlrfilTxtHtml,
		"css":  Html_httpHdlrfilTxtHtml,
		"js":   Html_httpHdlrfilTxtHtml,
		"html": Html_httpHdlrfilTxtHtml,
		"htm":  Html_httpHdlrfilTxtHtml,
		"json": jsonfl_httpHdlrFilJson,
		"jpg":  img_httpHdlrFilImg,
		"png":  img_httpHdlrFilImg,
	}

	httpHdlrFil(context, extNhdlrChoosrMaplist)

	// 处理特定API
	httpHdlrApiSpecl(context)

	// 设置响应内容类型和编码
	SetRespContentTypeNencode(context, "application/json; charset=utf-8")
	fn := path[1:]
	result := callxTryx(api_prefix+fn, queryString[1:])

	// 发送响应
	SendResp(ToString(result), context)
}
func Html_httpHdlrfilTxtHtml(context HttpContext) {
	// 获取当前请求的 URL
	request := context.Request
	path := request.URL.Path

	// 设置响应内容类型和编码
	context.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
	f := filepath.Join(webrootDir, decodeUrl(path))
	rzt2, err := ReadAllText(f)
	if err != nil {
		http.Error(context.Response, "File not found", http.StatusNotFound)
		return
	}
	context.Response.Write([]byte(rzt2))
	jmp2end()
}

var webrootDir = "/path/to/webroot" // 需要设置为实际的webroot目录
func decodeUrl(path string) string {
	// 示例解码URL逻辑，根据需要实现
	return path
}

func ReadAllText(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

//	func jmp2end() {
//		// 示例跳转结束逻辑，根据需要实现
//		fmt.Println("End of request handling")
//	}
func httpHdlrFil(ctx HttpContext, handlers map[string]func(HttpContext)) {
	path := ctx.Request.URL.Path
	for ext, handler := range handlers {
		if strings.HasSuffix(path, ext) {
			if handler != nil {
				handler(ctx)
			}
			return
		}
	}
}

func SetRespContentTypeNencode(ctx HttpContext, contentType string) {
	ctx.Response.Header().Set("Content-Type", contentType)
}

//func callxTryx(funcName string, params string) string {
//	// 示例调用逻辑
//	return fmt.Sprintf("Called %s with params %s", funcName, params)
//}

func SendResp(result string, ctx HttpContext) {
	// 示例发送响应逻辑
	ctx.Response.Write([]byte(result))
}

// 定义一个函数类型，代表被调用的方法签名
//type methodFunc func(args ...interface{}) interface{}

// 函数映射，用于存储方法名与具体方法实现的映射关系
var methodMap = make(map[string]methodFunc)

// 初始化方法映射
func initMethodMap() {
	// 使用反射获取当前包中所有的函数名，并放入方法映射中
	//funcs := ""
	////reflect.ValueOf(getFuncNames)
	//for i := 0; i < funcs.NumField(); i++ {
	//	funcName := funcs.Type().Field(i).Name
	//	methodMap[funcName] = nil // 暂时设为nil，可以根据需要设置具体的函数指针
	//}
}

// 模拟获取函数名的结构体
type getFuncNames struct {
	Func1 func()
	Func2 func(int)
	Func3 func(string) int
	// 添加更多函数名
}

func http请求处理器(特定api func(HttpContext), http上下文 HttpContext, api前缀 string) {
	//context2 := http上下文
	//HTTP响应对象 := context2.Response
	HTTP请求对象 := http上下文.Request
	//	url := fmt.Sprintf("%s://%s%s%s", HTTP请求对象.URL.Scheme, HTTP请求对象.Host, HTTP请求对象.URL.Path, HTTP请求对象.URL.RawQuery)
	路径 := HTTP请求对象.URL.Path

	// 静态资源处理器映射表
	扩展名与处理器对应表 := map[string]func(HttpContext){
		"txt":  html文件处理器,
		"css":  html文件处理器,
		"js":   html文件处理器,
		"html": html文件处理器,
		"htm":  html文件处理器,
		"json": jsonfl_httpHdlrFilJson,
		"jpg":  img_httpHdlrFilImg,
		"png":  img_httpHdlrFilImg,
	}

	文件响应处理(http上下文, 扩展名与处理器对应表)

	// 处理特定API
	特定api(http上下文)

	// 设置响应内容类型和编码
	设置响应内容类型和编码(http上下文, "application/json; charset=utf-8")
	函数名称 := 路径[1:]
	查询字符串 := HTTP请求对象.URL.RawQuery
	输出结果 := 调用(api前缀+函数名称, 查询字符串)

	发送响应(输出结果, http上下文)
}

func html文件处理器(ctx HttpContext) {
	// 示例HTML文件处理逻辑
	fmt.Fprintf(ctx.Response, "HTML File Handler")
}

func jsonfl_httpHdlrFilJson(ctx HttpContext) {
	// 示例JSON文件处理逻辑
	fmt.Fprintf(ctx.Response, "JSON File Handler")
}

func img_httpHdlrFilImg(ctx HttpContext) {
	// 示例图片文件处理逻辑
	fmt.Fprintf(ctx.Response, "Image File Handler")
}

func 文件响应处理(ctx HttpContext, 扩展名与处理器对应表 map[string]func(HttpContext)) {
	// 示例文件响应处理逻辑
	路径 := ctx.Request.URL.Path
	for 扩展名, 处理器 := range 扩展名与处理器对应表 {
		if strings.HasSuffix(路径, 扩展名) {
			处理器(ctx)
			return
		}
	}
}

type HttpContext struct {
	Response http.ResponseWriter
	Request  *http.Request
}

func 设置响应内容类型和编码(ctx HttpContext, contentType string) {
	ctx.Response.Header().Set("Content-Type", contentType)
}

func 调用(funcName string, params string) interface{} {
	// 示例调用逻辑
	return fmt.Sprintf("Called %s with params %s", funcName, params)
}

func 发送响应(result interface{}, ctx HttpContext) {
	// 示例发送响应逻辑
	fmt.Fprintf(ctx.Response, "%v", result)
}

// dep
func wbapiStart() {
	// 设置路由和处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 获取当前请求的 URL
		//url := fmt.Sprintf("%s://%s%s%s", r.URL.Scheme, r.URL.Host, r.URL.Path, r.URL.RawQuery)

		// 获取查询字符串
		queryString := r.URL.Query().Encode()
		method := r.URL.Path[1:]

		// 调用自定义函数
		result := callxTryx("Wbapi_"+method, queryString)

		// 设置响应头
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		// 返回结果
		fmt.Fprintf(w, "%v", result)
	})

	// 监听特定端口
	port := ":5000"
	fmt.Printf("Server is running on port %s...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}

// 模拟 callxTryx 函数
func callxTryx(methodName string, args ...interface{}) interface{} {
	fmt.Println("methodName:", methodName)

	// 打印方法调用信息和参数
	printCallFunArgs(methodName, args)

	// 这里模拟获取所有程序集和类型列表
	// 在 Go 中通常不需要显式获取所有类型和方法，因为 Go 是静态编译的语言
	fmt.Println("Simulating getting assemblies and types...")

	// 模拟调用方法
	// 这里假设 methodName 是一个函数或方法的名称
	// 实际情况下，需要根据 methodName 找到对应的函数或方法并调用
	// 在 Go 中，可以通过反射来动态调用方法
	// 这里只是简单的演示，实际应用中需要根据具体情况实现
	result := invokeMethod(methodName, args...)

	// 打印方法返回值
	printRetAdv(methodName, result)

	return result
}

// 模拟调用方法
func invokeMethod(methodName string, args ...interface{}) interface{} {
	// 这里模拟调用方法并返回结果
	// 在 Go 中可以使用反射来实现类似的功能
	fmt.Println("Simulating invoking method:", methodName)

	// 模拟方法返回一个结果
	return "Simulated result"
}

// 模拟打印方法调用信息和参数
func printCallFunArgs(methodName string, args []interface{}) {
	fmt.Println("Printing call function arguments for method:", methodName)
	fmt.Println("Arguments:", args)
}

// 模拟打印方法返回值
func printRetAdv(methodName string, result interface{}) {
	fmt.Println("Printing return value for method:", methodName)
	fmt.Println("Result:", result)
}
