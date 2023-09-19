package lib

import (
	"strings"
)

var tmp string

func Right(s string, lenx int) string {

	if Len(s) == 0 {
		return ""
	}

	rzt := leftPart2(s, lenx)
	if Len(rzt) == 0 {
		return tmp
	} else {
		return rzt
	}

}
func Left(s string, lenx int) string {

	if Len(s) == 0 {
		return ""
	}

	rzt := leftPart2(s, lenx)
	if Len(rzt) == 0 {
		return tmp
	} else {
		return rzt
	}

}
func CONCAT(s1 string, s2 string) string {

	return s1 + s2
}

//func REPLACE(s string, find string, toRplsStr string) string {
//
//}

//	func F长度(字符串 string) int {
//		return Len(字符串)
//	}
//
//	func 长度(字符串 string) int {
//		return Len(字符串)
//	}
func Len(s string) int {
	return strings.Count(s, "") - 1
}

func Substr(s string, start int, endIdx int) string {
	str := s // "a中文cd"
	str = string([]rune(str)[start:endIdx])
	return str
	//fmt.Println(str)
}

func leftPart2(s string, lenx int) string {
	defer func() string {
		if err := recover(); err != nil {
			//	fmt.Println(err)

			content := s[1 : len(s)-1]
			tmp = content
			return content
			//  str [ 1 : len (str)- 1]
			//...handle  打日志等
		}
		return s
	}()

	content := s[1:lenx]

	return content
}
