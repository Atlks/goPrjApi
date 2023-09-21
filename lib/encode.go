package lib

import (
	"encoding/json"
)

func Json_encode(obj any) string {

	bytes, _ := json.Marshal(obj)
	return (string(bytes))
}

func Json_decode(str string) {
	//unmarshaling
}

func Url_encode(str string) {
	//unmarshaling
}
func Url_decode(str string) {
	//unmarshaling
}

func Md5(str string) {
	//unmarshaling
}
