package lib

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

// ParseQueryString converts a URL query string to a map
func parse_str_qrystr(queryStr string) (map[string]string, error) {
	// Parse the query string
	values, err := url.ParseQuery(queryStr)
	if err != nil {
		return nil, err
	}

	// Create a map to hold the parsed values
	result := make(map[string]string)
	for key, value := range values {
		// Join multiple values with a comma if there are any
		result[key] = value[0]
	}

	return result, nil
}

func EncodeShellCmd(rcd string) string {

	// String ( JSON.stringify(rcd)  )
	return ""
}

func HexEncode(src []byte) string {
	// 编码
	//src := []byte("hello")
	//maxEnLen := hex.EncodedLen(len(src)) // 最大编码长度
	//dst1 := make([]byte, maxEnLen)
	//n := hex.Encode(dst1, src)
	dst2 := hex.EncodeToString(src)
	//fmt.Println("编码后的结果:", string(dst1[:n]))
	//fmt.Println("编码后的结果:", dst2)
	return dst2
}

func UnicodeEncode(ch string) string {
	return EscapeUnicode(ch)
}

func HtmlEncode(ch string) string {
	return EscapeUnicode(ch)
}

// EscapeUnicode 字符转码成unicode编码
func EscapeUnicode(text string) string {
	unicodeText := strconv.QuoteToASCII(text)
	// 去掉返回内容两端多余的双引号
	return unicodeText[1 : len(unicodeText)-1]
}

// UnescapeUnicode 将unicode编码转换成中文
func UnescapeUnicode(uContent string) (string, error) {
	// 转码前需要先增加上双引号，Quote增加双引号会将\u转义成\\u，同时会处理一些非打印字符
	content := strings.Replace(strconv.Quote(uContent), `\\u`, `\u`, -1)
	text, err := strconv.Unquote(content)
	if err != nil {
		return "", err
	}
	return text, nil
}

func Strip_tags(t string) string {
	return ""
}

func Json_encode(obj any) string {

	//prefix ==per line frt char
	bytes, _ := json.MarshalIndent(obj, "", "  ")
	return (string(bytes))
}
func EncodeJson(obj any) string {

	//prefix ==per line frt char
	bytes, _ := json.MarshalIndent(obj, "", "  ")
	return (string(bytes))
}

func Json_decode(str string) map[string]any {
	data := []byte(str)
	mapInterface := make(map[string]any)
	json.Unmarshal(data, &mapInterface)
	return mapInterface
}

//for k, v := range m {
//mapInterface[key] = value
//}

func Url_encode(urlStr string) string {
	escapeUrl := url.QueryEscape(urlStr)
	return escapeUrl
}
func Url_decode(escapeUrl string) string {
	urldecode, _ := url.QueryUnescape(escapeUrl)
	return urldecode
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Base64_encode(btarr []byte) string {
	encoded := base64.StdEncoding.EncodeToString(btarr)
	return encoded
}

func Base64_decode(encoded string) []byte {

	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	return (decoded)
}
