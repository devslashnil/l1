package task

import "strings"

var justString string

func createHugeString(count int) string {
	return strings.Repeat("语", count)
}

func someFunc() {
	v := createHugeString(1 << 10)
	// байта не достаточно для уникод символов
	s := []rune(v)
	justString = string(s[:100])
}

func RunSomeFunc() string {
	someFunc()
	return justString
}
