package 类库包

import (
	"fmt"
	"testing"
)

func TestI登录流程(t *testing.T) {

	//I登录流程("aaa", "123")
	m := map[string]string{
		"用户名": "我是谁啊",
		"密码":  "111",
	}
	fmt.Println(m)
	//I登录流程(m)

}
