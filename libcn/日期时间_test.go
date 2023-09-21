package 类库包

import (
	"fmt"
	"os"
	"testing"
)

func TestF当前时间(t *testing.T) {

	I显示(F当前时间())
	//if ans := F当前时间(); ans != 3 {
	//	t.Errorf("1 + 2 expected be 3, but %d got", ans)
	//}
}

func TestF时间戳(t *testing.T) {

	I显示(F时间戳())
	//if ans := F当前时间(); ans != 3 {
	//	t.Errorf("1 + 2 expected be 3, but %d got", ans)
	//}
}

func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
