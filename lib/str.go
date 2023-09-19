package lib

import (
	"fmt"
	"strings"
)

var tmp string

func Left(s string, lenx int) string {
	rzt := leftPart2(s, lenx)
	if Len(rzt) == 0 {
		return tmp
	} else {
		return rzt
	}

}

func Len(s string) int {
	return strings.Count(s, "") - 1
}

func leftPart2(s string, lenx int) string {
	defer func() string {
		if err := recover(); err != nil {
			fmt.Println(err)

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
