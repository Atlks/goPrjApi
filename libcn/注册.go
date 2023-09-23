package 类库包

func I注册流程(注册信息 map[string]string) {

	检查结果 := 检查注册信息(注册信息)
	如果(检查通过(检查结果),

		添加操作日志("用户 (@用户名@)登录，时间 @当前时间@"),
	)
	如果(I检查不通过(检查结果), 提示并终止(检查结果))

}

func 检查注册信息(注册信息 map[string]string) string {
	return I检查结果_通过
}
