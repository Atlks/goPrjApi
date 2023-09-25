package 类库包

import "goapiPrj/lib"

func 去除Html标记(html string) string {
	return lib.Strip_tags(html)
}

func 去除Html标记(html string) string {
	return lib.Strip_tags(html)
}

func shell参数编码(shell参数 string) string {
	return lib.EncodeShellCmd(shell参数)
}

func json编码(内容 any) string {
	return lib.Json_encode(内容)
}

func json解码(内容 string) map[string]interface{} {
	return lib.Json_decode(内容)
}

func url编码(内容 string) string {
	return lib.Url_encode(内容)

}
func url解码(内容 string) string {
	return lib.Url_decode(内容)
}

func md5签名编码(内容 string) string {
	return lib.Md5(内容)
}

func base64编码(字节数组 []byte) string {
	return lib.Base64_encode(字节数组)
}

func base64解码(字符串 string) []byte {
	return lib.Base64_decode(字符串)
}
