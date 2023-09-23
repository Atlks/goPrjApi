package lib

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"net/url"
)

func EncodeShellCmd(rcd string) string {

	// String ( JSON.stringify(rcd)  )
	return ""
}

func Strip_tags(t string) string {
	return ""
}

func Json_encode(obj any) string {

	bytes, _ := json.Marshal(obj)
	return (string(bytes))
}

func Json_decode(str string) map[string]interface{} {
	data := []byte(str)
	mapInterface := make(map[string]interface{})
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
