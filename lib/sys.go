package lib

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Echo(s any) {
	fmt.Println(s)
}

func IntToStr(xx int) string {
	return strconv.Itoa(xx)
}

func RdmMax(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func Rdm() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(999999999)
}

func Rdm2() int {
	n := rand.Int()

	return n
}

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomString returns a random string with a fixed length
func RandomString(n int, allowedChars ...[]rune) string {
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
